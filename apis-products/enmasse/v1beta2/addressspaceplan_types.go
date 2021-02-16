// +kubebuilder:object:generate=false
// +kubebuilder:skip
// +kubebuilder:skipversion
package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AddressSpacePlanSpec defines the desired state of AddressSpacePlan
type AddressSpacePlanSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	DisplayName      string                         `json:"displayName"`
	DisplayOrder     int                            `json:"displayOrder"`
	InfraConfigRef   string                         `json:"infraConfigRef"`
	ShortDescription string                         `json:"shortDescription"`
	LongDescription  string                         `json:"longDescription"`
	AddressSpaceType string                         `json:"addressSpaceType"`
	ResourceLimits   AddressSpacePlanResourceLimits `json:"resourceLimits"`
	AddressPlans     []string                       `json:"addressPlans"`
}

// AddressSpacePlanStatus defines the observed state of AddressSpacePlan
type AddressSpacePlanStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

type AddressSpacePlanResourceLimits struct {
	Router    float32 `json:"router"`
	Broker    float32 `json:"broker"`
	Aggregate float32 `json:"aggregate"`
}

// AddressSpacePlan is the Schema for the addressspaceplans API
type AddressSpacePlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AddressSpacePlanSpec   `json:"spec,omitempty"`
	Status AddressSpacePlanStatus `json:"status,omitempty"`
}

// AddressSpacePlanList contains a list of AddressSpacePlan
type AddressSpacePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AddressSpacePlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AddressSpacePlan{}, &AddressSpacePlanList{})
}
