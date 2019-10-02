package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Credential holds cluster-wide information about credentials.  The canonical name is `cluster`.
type Credential struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
	// +kubebuilder:validation:Required
	// +required
	Spec CredentialSpec `json:"spec"`
	// status holds observed values from the cluster. They may not be overridden.
	// +optional
	Status CredentialStatus `json:"status"`
}

// CredentialSpec holds the cloud credentials operator configuration.
type CredentialSpec struct {

	// disabled can be set to true to fully disable the cloud credentials operator.
	// Administrators will need to manually configure credentials for all
	// cluster components.
	Disabled bool `json:"disabled"`
}

type CredentialStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CredentialList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ListMeta `json:"metadata"`
	Items           []Credential `json:"items"`
}
