/*
Copyright 2023.

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

// ClusterDomainSpec defines the desired state of ClusterDomain
type ClusterDomainSpec struct {
	Domain string `json:"domain,omitempty"`
}

// ClusterDomainStatus defines the observed state of ClusterDomain
type ClusterDomainStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:printcolumn:name="Domain",type=string,JSONPath=`.spec.domain`

// ClusterDomain is the Schema for the clusterdomains API
type ClusterDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterDomainSpec   `json:"spec,omitempty"`
	Status ClusterDomainStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterDomainList contains a list of ClusterDomain
type ClusterDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterDomain `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterDomain{}, &ClusterDomainList{})
}
