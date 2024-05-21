/*
Copyright (c) 2024 Intel Corporation.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	MegaEmbedding MegaServiceType = "embedding"
	MegaRetrieval MegaServiceType = "retrieval"
	MegaGuardrail MegaServiceType = "guardrail"
	MegaLLM       MegaServiceType = "llm"
	MegaReranking MegaServiceType = "reranking"
	MegaTTS       MegaServiceType = "tts"
	MegaASR       MegaServiceType = "asr"
)

// Placeholder for service configurations
type ServiceConfiguration struct {
}

// MegaServiceType defines type of MegaService
type MegaServiceType string

// ServicePort defines port information of MegaService
type ServicePort struct {
	Name    string `json:"name,omitempty"`
	PortNum int    `json:"portnum"`
}

// MegaServiceSpec defines the desired state of MegaService
type MegaServiceSpec struct {
	ServiceIP   string          `json:"serviceIp"`
	ServiceType MegaServiceType `json:"megaServiceType"`
	Ports       []ServicePort   `json:"ip"`
	//optional
	Configurations ServiceConfiguration `json:"configuration,omitempty"`
}

// MegaServiceStatus defines the observed state of MegaService
type MegaServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MegaService is the Schema for the megaservices API
type MegaService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MegaServiceSpec   `json:"spec,omitempty"`
	Status MegaServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MegaServiceList contains a list of MegaService
type MegaServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MegaService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MegaService{}, &MegaServiceList{})
}
