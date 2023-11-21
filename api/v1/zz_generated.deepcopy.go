//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreDNSConfig) DeepCopyInto(out *CoreDNSConfig) {
	*out = *in
	if in.Plugin != nil {
		in, out := &in.Plugin, &out.Plugin
		*out = new(CoreDNSPluginConfig)
		**out = **in
	}
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreDNSConfig.
func (in *CoreDNSConfig) DeepCopy() *CoreDNSConfig {
	if in == nil {
		return nil
	}
	out := new(CoreDNSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreDNSPluginConfig) DeepCopyInto(out *CoreDNSPluginConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreDNSPluginConfig.
func (in *CoreDNSPluginConfig) DeepCopy() *CoreDNSPluginConfig {
	if in == nil {
		return nil
	}
	out := new(CoreDNSPluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressConfig) DeepCopyInto(out *IngressConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressConfig.
func (in *IngressConfig) DeepCopy() *IngressConfig {
	if in == nil {
		return nil
	}
	out := new(IngressConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IsolationConfig) DeepCopyInto(out *IsolationConfig) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.PodSecurityStandard != nil {
		in, out := &in.PodSecurityStandard, &out.PodSecurityStandard
		*out = new(string)
		**out = **in
	}
	if in.NodeProxyPermission != nil {
		in, out := &in.NodeProxyPermission, &out.NodeProxyPermission
		*out = new(NodeProxyPermission)
		**out = **in
	}
	if in.ResourceQuota != nil {
		in, out := &in.ResourceQuota, &out.ResourceQuota
		*out = new(ResourceQuota)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IsolationConfig.
func (in *IsolationConfig) DeepCopy() *IsolationConfig {
	if in == nil {
		return nil
	}
	out := new(IsolationConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeProxyPermission) DeepCopyInto(out *NodeProxyPermission) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeProxyPermission.
func (in *NodeProxyPermission) DeepCopy() *NodeProxyPermission {
	if in == nil {
		return nil
	}
	out := new(NodeProxyPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuota) DeepCopyInto(out *ResourceQuota) {
	*out = *in
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.ScopeSelector != nil {
		in, out := &in.ScopeSelector, &out.ScopeSelector
		*out = new(corev1.ScopeSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]corev1.ResourceQuotaScope, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuota.
func (in *ResourceQuota) DeepCopy() *ResourceQuota {
	if in == nil {
		return nil
	}
	out := new(ResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncConfig) DeepCopyInto(out *SyncConfig) {
	*out = *in
	if in.Ingresses != nil {
		in, out := &in.Ingresses, &out.Ingresses
		*out = new(IngressConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncConfig.
func (in *SyncConfig) DeepCopy() *SyncConfig {
	if in == nil {
		return nil
	}
	out := new(SyncConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualCluster) DeepCopyInto(out *VirtualCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualCluster.
func (in *VirtualCluster) DeepCopy() *VirtualCluster {
	if in == nil {
		return nil
	}
	out := new(VirtualCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterList) DeepCopyInto(out *VirtualClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterList.
func (in *VirtualClusterList) DeepCopy() *VirtualClusterList {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterSpec) DeepCopyInto(out *VirtualClusterSpec) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CoreDNS != nil {
		in, out := &in.CoreDNS, &out.CoreDNS
		*out = new(CoreDNSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Isolation != nil {
		in, out := &in.Isolation, &out.Isolation
		*out = new(IsolationConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Sync != nil {
		in, out := &in.Sync, &out.Sync
		*out = new(SyncConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterSpec.
func (in *VirtualClusterSpec) DeepCopy() *VirtualClusterSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualClusterStatus) DeepCopyInto(out *VirtualClusterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualClusterStatus.
func (in *VirtualClusterStatus) DeepCopy() *VirtualClusterStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualClusterStatus)
	in.DeepCopyInto(out)
	return out
}
