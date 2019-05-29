/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/akrausesap/goexperiments/openconnectors/util/connectorclient"
	"github.com/akrausesap/goexperiments/openconnectors/util/reconciletrigger"

	"k8s.io/apimachinery/pkg/types"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	openconnectors "github.com/akrausesap/goexperiments/openconnectors/api/v1"

	applicationoperator "github.com/kyma-project/kyma/components/application-operator/pkg/apis/applicationconnector/v1alpha1"
)

// ConnectorInstanceReconciler reconciles a ConnectorInstance object
type ConnectorInstanceReconciler struct {
	client.Client
	Log    logr.Logger
	Events chan event.GenericEvent
}

//cache of all reconcilers
var recociletriggers map[types.NamespacedName]reconciletrigger.ReconcileTrigger = make(map[types.NamespacedName]reconciletrigger.ReconcileTrigger)
var mutex = &sync.Mutex{}

// +kubebuilder:rbac:groups=openconnectors.incubator.kyma-project.io,resources=connectorinstances,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=openconnectors.incubator.kyma-project.io,resources=connectorinstances/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=applicationconnector.kyma-project.io,resources=applications,verbs=get;update;patch;create;list;watch

func createOpenConnectorsAuthorizationHeader(userSecret string, orgSecret string, APIKey string) (result *map[string][]string) {
	header := make(map[string][]string)

	header["Authorization"] = []string{
		fmt.Sprintf("User %s, Organization %s, Element %s",
			userSecret, orgSecret, APIKey),
	}
	return &header
}

func createApplicationName(openConnectorsInstanceName string, openConnectorsInstanceID string) string {
	return fmt.Sprintf("%s-%s", openConnectorsInstanceName, openConnectorsInstanceID)
}

func (r *ConnectorInstanceReconciler) setConnectorStatus(ctx context.Context,
	connectorInstance *openconnectors.ConnectorInstance) error {

	return r.Status().Update(ctx, connectorInstance)
}

func (r *ConnectorInstanceReconciler) setConnectorStatusError(ctx context.Context,
	connectorInstance *openconnectors.ConnectorInstance,
	errorReason string) error {

	connectorInstance.Status.State = "Error"
	connectorInstance.Status.ErrorReason = errorReason

	return r.setConnectorStatus(ctx, connectorInstance)
}

func (r *ConnectorInstanceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("connectorinstance", req.NamespacedName)

	var connector openconnectors.ConnectorInstance
	if err := r.Get(ctx, req.NamespacedName, &connector); err != nil {
		log.Error(err, "unable to fetch ConnectorInstance")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		if apierrs.IsNotFound(err) {
			//clean cache and leave, even if nothing is in there
			mutex.Lock()
			delete(recociletriggers, req.NamespacedName)
			mutex.Unlock()
			return ctrl.Result{Requeue: false}, nil
		} else {
			return ctrl.Result{Requeue: true}, err
		}
	}

	//Check for and stop reconciletriggers (can be the case if reconciler was triggered through change to resource)
	mutex.Lock()
	trigger := recociletriggers[req.NamespacedName]
	if trigger.IsInitialized() {
		trigger.Stop()
	} else {
		trigger = reconciletrigger.NewReconcileTrigger(r.Events, log, connector)
		recociletriggers[req.NamespacedName] = trigger
	}
	mutex.Unlock()

	//Read Connector Instances from Open Connectors

	connectorInstances, err := connectorclient.GetConnectorInstances(connector.Spec.Host,
		connector.Spec.UserSecret,
		connector.Spec.OrganizationSecret,
		connector.Spec.FilterTags,
	)
	if err != nil {
		log.Error(err, "unable to read list of Connector Instances from Open Connectors")
		r.setConnectorStatusError(ctx, &connector, err.Error())
		return ctrl.Result{Requeue: true}, err
	}

	//Label to identify/mark applications owned by this connectorinstance
	ownerLabels := map[string]string{
		"owned-by-connectorinstance-name": connector.Name,
		"owned-by-connectorinstance-uid":  fmt.Sprint(connector.UID),
	}

	//Get Applications that already exist
	var ownedApplicationList applicationoperator.ApplicationList
	if err := r.List(ctx, &ownedApplicationList, client.MatchingLabels(ownerLabels)); err != nil && !apierrs.IsNotFound(err) {
		log.Error(err, "unable to fetch List of Applications")
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		r.setConnectorStatusError(ctx, &connector, err.Error())
		return ctrl.Result{Requeue: true}, err
	}

	//filter those that need to be created newly (ignore deletions, changes, etc.)
	//this is part of another day's discussion
	connectorInstancesToCreate := &[]connectorclient.ConnectorInstance{}

	for i := range connectorInstances {

		targetApplicationName := createApplicationName(connector.Name, connectorInstances[i].ID)
		found := false
		for _, application := range ownedApplicationList.Items {
			//If exists stof and notify
			if application.Name == targetApplicationName {
				found = true
				break
			}
		}
		// if not found add to connectorInstancesToCreate
		if !found {
			*connectorInstancesToCreate = append(*connectorInstancesToCreate, connectorInstances[i])
		}
	}

	for _, connectorInstance := range *connectorInstancesToCreate {
		applicationName := createApplicationName(connector.Name, connectorInstance.ID)
		application := applicationoperator.Application{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Application",
				APIVersion: "applicationconnector.kyma-project.io/v1alpha1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:   applicationName,
				Labels: ownerLabels,
				OwnerReferences: []metav1.OwnerReference{
					metav1.OwnerReference{
						APIVersion: connector.APIVersion,
						Kind:       connector.Kind,
						Name:       connector.Name,
						UID:        connector.UID,
					},
				},
			},

			Spec: applicationoperator.ApplicationSpec{
				Services: []applicationoperator.Service{
					applicationoperator.Service{
						Name: connectorInstance.ID,
						DisplayName: fmt.Sprintf("%s - %s - %s",
							connector.Spec.DisplayName,
							connectorInstance.ConnectorName,
							connectorInstance.Name),
						Entries: []applicationoperator.Entry{
							applicationoperator.Entry{
								TargetUrl: fmt.Sprintf("https://%s/elements/api-v2/", connector.Spec.Host),
								SpecificationUrl: fmt.Sprintf("https://%s/elements/api-v2/instances/%s/docs",
									connector.Spec.Host, connectorInstance.ID),
								Headers: createOpenConnectorsAuthorizationHeader(
									connector.Spec.OrganizationSecret,
									connector.Spec.UserSecret,
									connectorInstance.APIKey),
								Type: "API",
							},
						},
					},
				},
			}}

		if err := r.Create(ctx, &application); err != nil {
			log.Error(err, fmt.Sprintf("unable to create Application for %s / %s",
				connectorInstance.ID, connectorInstance.Name))
			r.setConnectorStatusError(ctx, &connector, err.Error())
			return ctrl.Result{Requeue: true}, err
		}

		connector.Status.OwnedApplicationList = append(connector.Status.OwnedApplicationList, applicationName)
	}

	//Set Status
	connector.Status.State = "Success"

	if err := r.Status().Update(ctx, &connector); err != nil {

		//No status update possible, hence no set to error status
		log.Error(err, "unable to update ConnectorInstance status")
		return ctrl.Result{}, err
	}

	// If there should be a periodic refresh, schedule next refresh
	if refreshInterval := connector.Spec.RefreshIntervalSeconds; refreshInterval > 0 {

		trigger.Start(time.Duration(refreshInterval) * time.Second)

	}
	return ctrl.Result{Requeue: false}, nil
}

func (r *ConnectorInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {

	return ctrl.NewControllerManagedBy(mgr).
		Watches(&source.Channel{Source: r.Events},
			&handler.EnqueueRequestForObject{}).
		For(&openconnectors.ConnectorInstance{}).
		Complete(r)
}
