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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VirtualClusterSpec defines the desired state of VirtualCluster
type VirtualClusterSpec struct {
	Affinity    *corev1.Affinity  `json:"affinity,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty"`
	CoreDNS     *CoreDNS          `json:"coredns,omitempty"`
	Isolation   *Isolation        `json:"isolation,omitempty"`
	Sync        *Sync             `json:"sync,omitempty"`
	Syncer      *Syncer           `json:"syncer,omitempty"`
	Service     *Service          `json:"service,omitempty"`
	Ingress     *Ingress          `json:"ingress,omitempty"`
	VCluster    *VCluster         `json:"vcluster,omitempty"`
}

// VirtualClusterStatus defines the observed state of VirtualCluster
type VirtualClusterStatus struct {
	// Conditions holds several conditions the vcluster might be in
	// +optional
	Conditions Conditions `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VirtualCluster is the Schema for the virtualclusters API
type VirtualCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualClusterSpec   `json:"spec,omitempty"`
	Status VirtualClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VirtualClusterList contains a list of VirtualCluster
type VirtualClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualCluster{}, &VirtualClusterList{})
}
