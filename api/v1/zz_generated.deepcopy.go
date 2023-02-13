//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlippyCondition) DeepCopyInto(out *FlippyCondition) {
	*out = *in
	out.K8S = in.K8S
	out.StatusCheckConfig = in.StatusCheckConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlippyCondition.
func (in *FlippyCondition) DeepCopy() *FlippyCondition {
	if in == nil {
		return nil
	}
	out := new(FlippyCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlippyConfig) DeepCopyInto(out *FlippyConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlippyConfig.
func (in *FlippyConfig) DeepCopy() *FlippyConfig {
	if in == nil {
		return nil
	}
	out := new(FlippyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlippyConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlippyConfigList) DeepCopyInto(out *FlippyConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlippyConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlippyConfigList.
func (in *FlippyConfigList) DeepCopy() *FlippyConfigList {
	if in == nil {
		return nil
	}
	out := new(FlippyConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlippyConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlippyConfigSpec) DeepCopyInto(out *FlippyConfigSpec) {
	*out = *in
	if in.ImageList != nil {
		in, out := &in.ImageList, &out.ImageList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Preconditions != nil {
		in, out := &in.Preconditions, &out.Preconditions
		*out = make([]FlippyCondition, len(*in))
		copy(*out, *in)
	}
	in.ProcessFilter.DeepCopyInto(&out.ProcessFilter)
	if in.RestartObjects != nil {
		in, out := &in.RestartObjects, &out.RestartObjects
		*out = make([]RestartObject, len(*in))
		copy(*out, *in)
	}
	if in.PostFilterRestarts != nil {
		in, out := &in.PostFilterRestarts, &out.PostFilterRestarts
		*out = make([]FlippyCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlippyConfigSpec.
func (in *FlippyConfigSpec) DeepCopy() *FlippyConfigSpec {
	if in == nil {
		return nil
	}
	out := new(FlippyConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlippyConfigStatus) DeepCopyInto(out *FlippyConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlippyConfigStatus.
func (in *FlippyConfigStatus) DeepCopy() *FlippyConfigStatus {
	if in == nil {
		return nil
	}
	out := new(FlippyConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8S) DeepCopyInto(out *K8S) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8S.
func (in *K8S) DeepCopy() *K8S {
	if in == nil {
		return nil
	}
	out := new(K8S)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessFilter) DeepCopyInto(out *ProcessFilter) {
	*out = *in
	if in.PodLabels != nil {
		in, out := &in.PodLabels, &out.PodLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NamespaceLabels != nil {
		in, out := &in.NamespaceLabels, &out.NamespaceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.PreProcessRestart = in.PreProcessRestart
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessFilter.
func (in *ProcessFilter) DeepCopy() *ProcessFilter {
	if in == nil {
		return nil
	}
	out := new(ProcessFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RestartObject) DeepCopyInto(out *RestartObject) {
	*out = *in
	out.StatusCheckConfig = in.StatusCheckConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RestartObject.
func (in *RestartObject) DeepCopy() *RestartObject {
	if in == nil {
		return nil
	}
	out := new(RestartObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCheckConfig) DeepCopyInto(out *StatusCheckConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCheckConfig.
func (in *StatusCheckConfig) DeepCopy() *StatusCheckConfig {
	if in == nil {
		return nil
	}
	out := new(StatusCheckConfig)
	in.DeepCopyInto(out)
	return out
}
