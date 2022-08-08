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

package v2

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	commonv1 "github.com/kubeshop/testkube-operator/apis/common/v1"
	testkubev1 "github.com/kubeshop/testkube-operator/apis/testsuite/v1"
)

// ConvertTo converts this Script to the Hub version (v1).
func (src *TestSuite) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*testkubev1.TestSuite)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Repeats = src.Spec.Repeats
	dst.Spec.Description = src.Spec.Description
	dst.Spec.Schedule = src.Spec.Schedule

	dst.Spec.Before = make([]testkubev1.TestSuiteStepSpec, len(src.Spec.Before))
	dst.Spec.Steps = make([]testkubev1.TestSuiteStepSpec, len(src.Spec.Steps))
	dst.Spec.After = make([]testkubev1.TestSuiteStepSpec, len(src.Spec.After))

	var stepTypes = []struct {
		Source     []TestSuiteStepSpec
		Destinaton []testkubev1.TestSuiteStepSpec
	}{
		{
			Source:     src.Spec.Before,
			Destinaton: dst.Spec.Before,
		},
		{
			Source:     src.Spec.Steps,
			Destinaton: dst.Spec.Steps,
		},
		{
			Source:     src.Spec.After,
			Destinaton: dst.Spec.After,
		},
	}

	for _, stepType := range stepTypes {
		for i := range stepType.Source {
			value := stepType.Source[i]
			step := testkubev1.TestSuiteStepSpec{
				Type: value.Type,
			}

			if value.Delay != nil {
				step.Delay = &testkubev1.TestSuiteStepDelay{
					Duration: value.Delay.Duration,
				}
			}

			if value.Execute != nil {
				step.Execute = &testkubev1.TestSuiteStepExecute{
					Namespace:     value.Execute.Namespace,
					Name:          value.Execute.Name,
					StopOnFailure: value.Execute.StopOnFailure,
				}
			}

			stepType.Destinaton[i] = step
		}
	}

	if src.Spec.ExecutionRequest != nil {
		dst.Spec.Variables = make(map[string]testkubev1.Variable, len(src.Spec.ExecutionRequest.Variables))
		for key, value := range src.Spec.ExecutionRequest.Variables {
			dst.Spec.Variables[key] = testkubev1.Variable{
				Type_:     value.Type_,
				Name:      value.Name,
				Value:     value.Value,
				ValueFrom: value.ValueFrom,
			}
		}
	}

	return nil
}

// ConvertFrom converts Script from the Hub version (v1) to this version.
func (dst *TestSuite) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*testkubev1.TestSuite)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Repeats = src.Spec.Repeats
	dst.Spec.Description = src.Spec.Description
	dst.Spec.Schedule = src.Spec.Schedule

	dst.Spec.Before = make([]TestSuiteStepSpec, len(src.Spec.Before))
	dst.Spec.Steps = make([]TestSuiteStepSpec, len(src.Spec.Steps))
	dst.Spec.After = make([]TestSuiteStepSpec, len(src.Spec.After))

	var stepTypes = []struct {
		Source     []testkubev1.TestSuiteStepSpec
		Destinaton []TestSuiteStepSpec
	}{
		{
			Source:     src.Spec.Before,
			Destinaton: dst.Spec.Before,
		},
		{
			Source:     src.Spec.Steps,
			Destinaton: dst.Spec.Steps,
		},
		{
			Source:     src.Spec.After,
			Destinaton: dst.Spec.After,
		},
	}

	for _, stepType := range stepTypes {
		for i := range stepType.Source {
			value := stepType.Source[i]
			step := TestSuiteStepSpec{
				Type: value.Type,
			}

			if value.Delay != nil {
				step.Delay = &TestSuiteStepDelay{
					Duration: value.Delay.Duration,
				}
			}

			if value.Execute != nil {
				step.Execute = &TestSuiteStepExecute{
					Namespace:     value.Execute.Namespace,
					Name:          value.Execute.Name,
					StopOnFailure: value.Execute.StopOnFailure,
				}
			}

			stepType.Destinaton[i] = step
		}
	}

	if len(src.Spec.Variables) != 0 || len(src.Spec.Params) != 0 {
		dst.Spec.ExecutionRequest = &TestSuiteExecutionRequest{}
		dst.Spec.ExecutionRequest.Variables = make(map[string]Variable, len(src.Spec.Variables)+len(src.Spec.Params))
		for key, value := range src.Spec.Params {
			dst.Spec.ExecutionRequest.Variables[key] = Variable{
				Type_: commonv1.VariableTypeBasic,
				Name:  key,
				Value: value,
			}
		}

		for key, value := range src.Spec.Variables {
			dst.Spec.ExecutionRequest.Variables[key] = Variable{
				Type_:     value.Type_,
				Name:      value.Name,
				Value:     value.Value,
				ValueFrom: value.ValueFrom,
			}
		}
	}

	// Status
	return nil
}
