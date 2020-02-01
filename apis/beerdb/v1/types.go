package v1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Color string

const (
	GoldColor   = Color("gold")
	YellowColor = Color("yellow")
	BlondColor  = Color("blond")
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Beer is a beer sold by a brewery.
type Beer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Enum=gold;yellow;blond;unknown
	// +kubebuilder:default=unknown
	Color Color `json:"color,omitempty"`

	Alcohol resource.Quantity `json:"alcohol,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BeerList is a list of Sensor resources
type BeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Beer `json:"items"`
}
