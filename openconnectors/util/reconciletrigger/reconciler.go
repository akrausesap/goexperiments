package reconciletrigger

import (
	"errors"
	"fmt"
	"sync"
	"time"

	openconnectors "github.com/akrausesap/goexperiments/openconnectors/api/v1"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

//ReconcileTrigger contains all data needed to periodically trigger reconciliation for
// an Open Connectors Instance, only create via NewReconcileTrigger()
type ReconcileTrigger struct {
	eventChannel  chan<- event.GenericEvent
	log           logr.Logger
	meta          metav1.Object
	object        runtime.Object
	communication chan bool
	mutex         *sync.Mutex
	initialized   bool
}

//NewReconcileTrigger creactes a ReconcileTrigger Object
func NewReconcileTrigger(eventChannel chan<- event.GenericEvent,
	log logr.Logger, connector openconnectors.ConnectorInstance) (result ReconcileTrigger) {

	connectorCopy := connector.DeepCopy()

	result = ReconcileTrigger{
		eventChannel: eventChannel,
		log:          log,
		meta:         connectorCopy.GetObjectMeta(),
		object:       connectorCopy,
		mutex:        &sync.Mutex{},
		initialized:  true,
	}
	return
}

func (rt *ReconcileTrigger) IsInitialized() bool {
	return rt.initialized
}

//Start an event triggger
func (rt *ReconcileTrigger) Start(duration time.Duration) error {
	if !rt.initialized {
		return errors.New("ReconcileTrigger is not initialized")
	}
	rt.mutex.Lock()
	rt.communication = make(chan bool)
	rt.mutex.Unlock()
	go func() {
		select {
		//If stopped from outside
		case <-rt.communication:
			rt.log.Info(fmt.Sprintf("Event for refresh/reconcile of %s not triggered", rt.meta.GetName()))
			return
		//If time expires
		case <-time.After(duration):
			rt.log.Info(fmt.Sprintf("Event for refresh/reconcile of %s triggered", rt.meta.GetName()))
			rt.eventChannel <- event.GenericEvent{
				Meta:   rt.meta,
				Object: rt.object,
			}
		}
		rt.mutex.Lock()
		if rt.communication != nil {
			close(rt.communication)
			rt.communication = nil
		}
		rt.mutex.Unlock()
	}()
	return nil
}

//Stop an event triggger (without actually pulling the trigger)
func (rt *ReconcileTrigger) Stop() error {

	if !rt.initialized {
		return errors.New("ReconcileTrigger is not initialized")
	}
	rt.mutex.Lock()
	if rt.communication != nil {
		close(rt.communication)

	}
	rt.mutex.Unlock()
	return nil
}
