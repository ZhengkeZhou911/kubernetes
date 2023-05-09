//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ConfigurationMap) DeepCopyInto(out *ConfigurationMap) {
	{
		in := &in
		*out = make(ConfigurationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationMap.
func (in ConfigurationMap) DeepCopy() ConfigurationMap {
	if in == nil {
		return nil
	}
	out := new(ConfigurationMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DetectLocalConfiguration) DeepCopyInto(out *DetectLocalConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DetectLocalConfiguration.
func (in *DetectLocalConfiguration) DeepCopy() *DetectLocalConfiguration {
	if in == nil {
		return nil
	}
	out := new(DetectLocalConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyConfiguration) DeepCopyInto(out *KubeProxyConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.ClientConnection = in.ClientConnection
	in.Logging.DeepCopyInto(&out.Logging)
	in.IPTables.DeepCopyInto(&out.IPTables)
	in.IPVS.DeepCopyInto(&out.IPVS)
	out.Winkernel = in.Winkernel
	in.NFTables.DeepCopyInto(&out.NFTables)
	out.DetectLocal = in.DetectLocal
	if in.NodePortAddresses != nil {
		in, out := &in.NodePortAddresses, &out.NodePortAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OOMScoreAdj != nil {
		in, out := &in.OOMScoreAdj, &out.OOMScoreAdj
		*out = new(int32)
		**out = **in
	}
	in.Conntrack.DeepCopyInto(&out.Conntrack)
	out.ConfigSyncPeriod = in.ConfigSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyConfiguration.
func (in *KubeProxyConfiguration) DeepCopy() *KubeProxyConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeProxyConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyConntrackConfiguration) DeepCopyInto(out *KubeProxyConntrackConfiguration) {
	*out = *in
	if in.MaxPerCore != nil {
		in, out := &in.MaxPerCore, &out.MaxPerCore
		*out = new(int32)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(int32)
		**out = **in
	}
	if in.TCPEstablishedTimeout != nil {
		in, out := &in.TCPEstablishedTimeout, &out.TCPEstablishedTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.TCPCloseWaitTimeout != nil {
		in, out := &in.TCPCloseWaitTimeout, &out.TCPCloseWaitTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	out.UDPTimeout = in.UDPTimeout
	out.UDPStreamTimeout = in.UDPStreamTimeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyConntrackConfiguration.
func (in *KubeProxyConntrackConfiguration) DeepCopy() *KubeProxyConntrackConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyConntrackConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyIPTablesConfiguration) DeepCopyInto(out *KubeProxyIPTablesConfiguration) {
	*out = *in
	if in.MasqueradeBit != nil {
		in, out := &in.MasqueradeBit, &out.MasqueradeBit
		*out = new(int32)
		**out = **in
	}
	if in.LocalhostNodePorts != nil {
		in, out := &in.LocalhostNodePorts, &out.LocalhostNodePorts
		*out = new(bool)
		**out = **in
	}
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyIPTablesConfiguration.
func (in *KubeProxyIPTablesConfiguration) DeepCopy() *KubeProxyIPTablesConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyIPTablesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyIPVSConfiguration) DeepCopyInto(out *KubeProxyIPVSConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	if in.ExcludeCIDRs != nil {
		in, out := &in.ExcludeCIDRs, &out.ExcludeCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.TCPTimeout = in.TCPTimeout
	out.TCPFinTimeout = in.TCPFinTimeout
	out.UDPTimeout = in.UDPTimeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyIPVSConfiguration.
func (in *KubeProxyIPVSConfiguration) DeepCopy() *KubeProxyIPVSConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyIPVSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyNFTablesConfiguration) DeepCopyInto(out *KubeProxyNFTablesConfiguration) {
	*out = *in
	if in.MasqueradeBit != nil {
		in, out := &in.MasqueradeBit, &out.MasqueradeBit
		*out = new(int32)
		**out = **in
	}
	if in.LocalhostNodePorts != nil {
		in, out := &in.LocalhostNodePorts, &out.LocalhostNodePorts
		*out = new(bool)
		**out = **in
	}
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyNFTablesConfiguration.
func (in *KubeProxyNFTablesConfiguration) DeepCopy() *KubeProxyNFTablesConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyNFTablesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyWinkernelConfiguration) DeepCopyInto(out *KubeProxyWinkernelConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyWinkernelConfiguration.
func (in *KubeProxyWinkernelConfiguration) DeepCopy() *KubeProxyWinkernelConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyWinkernelConfiguration)
	in.DeepCopyInto(out)
	return out
}
