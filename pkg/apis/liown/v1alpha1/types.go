package v1


import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NotebookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Template NotebookTemplateSpec `json:"template,omitempty"`
}

type NotebookTemplateSpec struct {
	Spec corev1.PodSpec `json:"spec,omitempty"`
}

// NotebookStatus defines the observed state of Notebook
type NotebookStatus struct {
	// Conditions is an array of current conditions
	Conditions []NotebookCondition `json:"conditions"`
	// ReadyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	ReadyReplicas int32 `json:"readyReplicas"`
	// ContainerState is the state of underlying container.
	ContainerState corev1.ContainerState `json:"containerState"`
}

type NotebookCondition struct {
	// Type is the type of the condition. Possible values are Running|Waiting|Terminated
	Type string `json:"type"`
	// Last time we probed the condition.
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// (brief) reason the container is in the current state
	Reason string `json:"reason,omitempty"`
	// Message regarding why the container is in the current state.
	Message string `json:"message,omitempty"`
}

// Notebook is the Schema for the notebooks API
type Notebook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotebookSpec   `json:"spec,omitempty"`
	Status NotebookStatus `json:"status,omitempty"`
}


// NotebookList contains a list of Notebook
type NotebookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notebook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Notebook{}, &NotebookList{})
}

