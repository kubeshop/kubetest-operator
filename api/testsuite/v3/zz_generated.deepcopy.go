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

package v3

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownloadArtifactOptions) DeepCopyInto(out *DownloadArtifactOptions) {
	*out = *in
	if in.PreviousStepNumbers != nil {
		in, out := &in.PreviousStepNumbers, &out.PreviousStepNumbers
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownloadArtifactOptions.
func (in *DownloadArtifactOptions) DeepCopy() *DownloadArtifactOptions {
	if in == nil {
		return nil
	}
	out := new(DownloadArtifactOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunningContext) DeepCopyInto(out *RunningContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunningContext.
func (in *RunningContext) DeepCopy() *RunningContext {
	if in == nil {
		return nil
	}
	out := new(RunningContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuite) DeepCopyInto(out *TestSuite) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuite.
func (in *TestSuite) DeepCopy() *TestSuite {
	if in == nil {
		return nil
	}
	out := new(TestSuite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestSuite) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteBatchStep) DeepCopyInto(out *TestSuiteBatchStep) {
	*out = *in
	if in.DownloadArtifacts != nil {
		in, out := &in.DownloadArtifacts, &out.DownloadArtifacts
		*out = new(DownloadArtifactOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Execute != nil {
		in, out := &in.Execute, &out.Execute
		*out = make([]TestSuiteStepSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteBatchStep.
func (in *TestSuiteBatchStep) DeepCopy() *TestSuiteBatchStep {
	if in == nil {
		return nil
	}
	out := new(TestSuiteBatchStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteExecutionCore) DeepCopyInto(out *TestSuiteExecutionCore) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(TestSuiteExecutionStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteExecutionCore.
func (in *TestSuiteExecutionCore) DeepCopy() *TestSuiteExecutionCore {
	if in == nil {
		return nil
	}
	out := new(TestSuiteExecutionCore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteExecutionRequest) DeepCopyInto(out *TestSuiteExecutionRequest) {
	*out = *in
	if in.Variables != nil {
		in, out := &in.Variables, &out.Variables
		*out = make(map[string]Variable, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExecutionLabels != nil {
		in, out := &in.ExecutionLabels, &out.ExecutionLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RunningContext != nil {
		in, out := &in.RunningContext, &out.RunningContext
		*out = new(RunningContext)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteExecutionRequest.
func (in *TestSuiteExecutionRequest) DeepCopy() *TestSuiteExecutionRequest {
	if in == nil {
		return nil
	}
	out := new(TestSuiteExecutionRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteList) DeepCopyInto(out *TestSuiteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TestSuite, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteList.
func (in *TestSuiteList) DeepCopy() *TestSuiteList {
	if in == nil {
		return nil
	}
	out := new(TestSuiteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestSuiteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteSpec) DeepCopyInto(out *TestSuiteSpec) {
	*out = *in
	if in.Before != nil {
		in, out := &in.Before, &out.Before
		*out = make([]TestSuiteBatchStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]TestSuiteBatchStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.After != nil {
		in, out := &in.After, &out.After
		*out = make([]TestSuiteBatchStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExecutionRequest != nil {
		in, out := &in.ExecutionRequest, &out.ExecutionRequest
		*out = new(TestSuiteExecutionRequest)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteSpec.
func (in *TestSuiteSpec) DeepCopy() *TestSuiteSpec {
	if in == nil {
		return nil
	}
	out := new(TestSuiteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteStatus) DeepCopyInto(out *TestSuiteStatus) {
	*out = *in
	if in.LatestExecution != nil {
		in, out := &in.LatestExecution, &out.LatestExecution
		*out = new(TestSuiteExecutionCore)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteStatus.
func (in *TestSuiteStatus) DeepCopy() *TestSuiteStatus {
	if in == nil {
		return nil
	}
	out := new(TestSuiteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestSuiteStepSpec) DeepCopyInto(out *TestSuiteStepSpec) {
	*out = *in
	out.Delay = in.Delay
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestSuiteStepSpec.
func (in *TestSuiteStepSpec) DeepCopy() *TestSuiteStepSpec {
	if in == nil {
		return nil
	}
	out := new(TestSuiteStepSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Variable) DeepCopyInto(out *Variable) {
	*out = *in
	in.ValueFrom.DeepCopyInto(&out.ValueFrom)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Variable.
func (in *Variable) DeepCopy() *Variable {
	if in == nil {
		return nil
	}
	out := new(Variable)
	in.DeepCopyInto(out)
	return out
}
