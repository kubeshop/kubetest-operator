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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestTrigger) DeepCopyInto(out *TestTrigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestTrigger.
func (in *TestTrigger) DeepCopy() *TestTrigger {
	if in == nil {
		return nil
	}
	out := new(TestTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestTrigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestTriggerCondition) DeepCopyInto(out *TestTriggerCondition) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(TestTriggerConditionStatuses)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestTriggerCondition.
func (in *TestTriggerCondition) DeepCopy() *TestTriggerCondition {
	if in == nil {
		return nil
	}
	out := new(TestTriggerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestTriggerConditionSpec) DeepCopyInto(out *TestTriggerConditionSpec) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TestTriggerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestTriggerConditionSpec.
func (in *TestTriggerConditionSpec) DeepCopy() *TestTriggerConditionSpec {
	if in == nil {
		return nil
	}
	out := new(TestTriggerConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestTriggerList) DeepCopyInto(out *TestTriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TestTrigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestTriggerList.
func (in *TestTriggerList) DeepCopy() *TestTriggerList {
	if in == nil {
		return nil
	}
	out := new(TestTriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestTriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestTriggerProbe) DeepCopyInto(out *TestTriggerProbe) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestTriggerProbe.
func (in *TestTriggerProbe) DeepCopy() *TestTriggerProbe {
	if in == nil {
		return nil
	}
	out := new(TestTriggerProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestTriggerProbeSpec) DeepCopyInto(out *TestTriggerProbeSpec) {
	*out = *in
	if in.Probes != nil {
		in, out := &in.Probes, &out.Probes
		*out = make([]TestTriggerProbe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestTriggerProbeSpec.
func (in *TestTriggerProbeSpec) DeepCopy() *TestTriggerProbeSpec {
	if in == nil {
		return nil
	}
	out := new(TestTriggerProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestTriggerSelector) DeepCopyInto(out *TestTriggerSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestTriggerSelector.
func (in *TestTriggerSelector) DeepCopy() *TestTriggerSelector {
	if in == nil {
		return nil
	}
	out := new(TestTriggerSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestTriggerSpec) DeepCopyInto(out *TestTriggerSpec) {
	*out = *in
	in.ResourceSelector.DeepCopyInto(&out.ResourceSelector)
	if in.ConditionSpec != nil {
		in, out := &in.ConditionSpec, &out.ConditionSpec
		*out = new(TestTriggerConditionSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ProbeSpec != nil {
		in, out := &in.ProbeSpec, &out.ProbeSpec
		*out = new(TestTriggerProbeSpec)
		(*in).DeepCopyInto(*out)
	}
	in.TestSelector.DeepCopyInto(&out.TestSelector)
	if in.Delay != nil {
		in, out := &in.Delay, &out.Delay
		*out = new(metav1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestTriggerSpec.
func (in *TestTriggerSpec) DeepCopy() *TestTriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TestTriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestTriggerStatus) DeepCopyInto(out *TestTriggerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestTriggerStatus.
func (in *TestTriggerStatus) DeepCopy() *TestTriggerStatus {
	if in == nil {
		return nil
	}
	out := new(TestTriggerStatus)
	in.DeepCopyInto(out)
	return out
}
