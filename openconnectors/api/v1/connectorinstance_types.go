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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:scope:Cluster

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConnectorInstanceSpec defines the parameters needed to connect to an
// Open Connectors Instance
type ConnectorInstanceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//DisplayName that prefixes derrived applications
	DisplayName string `json:"displayName"`

	//FilterTags is used to determine relevant instances in SAP Cloud Platform Open Connectors
	// +optional
	FilterTags []string `json:"filterTags,omitempty"`

	// RefreshIntervalSeconds is used to determine how often the state form SAP Cloud Platform Open
	// Connectors shall be refresed (< 1 means no refresh)
	RefreshIntervalSeconds int `json:"refreshIntervalSeconds"`

	//Hostname of the Open Connectors Instance (without scheme, i.e. https)
	Host string `json:"hostName"`

	//Organization secret used to connect to the Open Connectors Instance
	OrganizationSecret string `json:"organizationSecret"`

	//User secret used to connect to the Open Connectors Instance
	UserSecret string `json:"userSecret"`
}

// ConnectorInstanceStatus defines the observed state of ConnectorInstance
type ConnectorInstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// State defines the current state (Error or success)
	State string `json:"state"`

	//ErrorReason provides the reason for an error if in error state
	ErrorReason string `json:"errorReason"`

	//OwnedApplicationList, lists all applications owned by this Connector Instance
	OwnedApplicationList []string `json:"ownedApplicationList"`
}

// ConnectorInstance is the Schema for the connectorinstances API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ConnectorInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConnectorInstanceSpec   `json:"spec,omitempty"`
	Status ConnectorInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectorInstanceList contains a list of ConnectorInstance
type ConnectorInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectorInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConnectorInstance{}, &ConnectorInstanceList{})
}
