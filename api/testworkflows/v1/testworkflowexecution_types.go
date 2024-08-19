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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// TestWorkflowExecutionSpec defines the desired state of TestWorkflowExecution
type TestWorkflowExecutionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	TestWorkflow     *corev1.LocalObjectReference  `json:"testWorkflow" expr:"include"`
	ExecutionRequest *TestWorkflowExecutionRequest `json:"executionRequest,omitempty" expr:"include"`
}

// TestWorkflowExecutionRequest contains TestWorkflow execution parameters
type TestWorkflowExecutionRequest struct {
	// custom execution name
	Name   string                        `json:"name,omitempty" expr:"template"`
	Config map[string]intstr.IntOrString `json:"config,omitempty" expr:"template"`
	// test workflow execution name started the test workflow execution
	TestWorkflowExecutionName string `json:"testWorkflowExecutionName,omitempty" expr:"template"`
	// whether webhooks should be disabled for this execution
	DisableWebhooks bool `json:"disableWebhooks,omitempty"`
	// test workflow execution tags
	Tags map[string]string `json:"tags,omitempty" expr:"template"`
}

// TestWorkflowExecutionStatus defines the observed state of TestWorkflowExecution
type TestWorkflowExecutionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	LatestExecution *TestWorkflowExecutionDetails `json:"latestExecution,omitempty"`
	// test workflow execution generation
	Generation int64 `json:"generation,omitempty"`
}

// TestWorkflowExecutionDetails contains the details of TestWorkflowExecution
type TestWorkflowExecutionDetails struct {
	// unique execution identifier
	Id string `json:"id"`
	// execution name
	Name string `json:"name"`
	// execution namespace
	Namespace string `json:"namespace,omitempty"`
	// sequence number for the execution
	Number int32 `json:"number,omitempty"`
	// when the execution has been scheduled to run
	ScheduledAt metav1.Time `json:"scheduledAt,omitempty"`
	// when the execution result's status has changed last time (queued, passed, failed)
	StatusAt metav1.Time `json:"statusAt,omitempty"`
	// structured tree of steps
	Signature []TestWorkflowSignature `json:"signature,omitempty"`
	Result    *TestWorkflowResult     `json:"result,omitempty"`
	// additional information from the steps, like referenced executed tests or artifacts
	Output []TestWorkflowOutput `json:"output,omitempty"`
	// generated reports from the steps, like junit
	Reports []TestWorkflowReport `json:"reports,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Workflow *TestWorkflow `json:"workflow"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	ResolvedWorkflow *TestWorkflow `json:"resolvedWorkflow,omitempty"`
	// test workflow execution name started the test workflow execution
	TestWorkflowExecutionName string `json:"testWorkflowExecutionName,omitempty"`
	// whether webhooks should be disabled for this execution
	DisableWebhooks bool `json:"disableWebhooks,omitempty"`
	// test workflow execution tags
	Tags map[string]string `json:"tags,omitempty"`
}

// TestWorkflowSignature has signature of TestWorkflow
type TestWorkflowSignature struct {
	// step reference
	Ref string `json:"ref,omitempty"`
	// step name
	Name string `json:"name,omitempty"`
	// step category, that may be used as name fallback
	Category string `json:"category,omitempty"`
	// is the step/group meant to be optional
	Optional bool `json:"optional,omitempty"`
	// is the step/group meant to be negative
	Negative bool `json:"negative,omitempty"`
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	Children []TestWorkflowSignature `json:"children,omitempty"`
}

// TestWorkflowResult contains result of TestWorkflow
type TestWorkflowResult struct {
	Status          *TestWorkflowStatus `json:"status"`
	PredictedStatus *TestWorkflowStatus `json:"predictedStatus"`
	// when the pod was created
	QueuedAt metav1.Time `json:"queuedAt,omitempty"`
	// when the pod has been successfully assigned
	StartedAt metav1.Time `json:"startedAt,omitempty"`
	// when the pod has been completed
	FinishedAt metav1.Time `json:"finishedAt,omitempty"`
	// Go-formatted (human-readable) duration
	Duration string `json:"duration,omitempty"`
	// Go-formatted (human-readable) total duration (incl. pause)
	TotalDuration string `json:"totalDuration,omitempty"`
	// Duration in milliseconds
	DurationMs int32 `json:"durationMs,omitempty"`
	// Pause duration in milliseconds
	PausedMs int32 `json:"pausedMs,omitempty"`
	// Total duration in milliseconds (incl. pause)
	TotalDurationMs int32                             `json:"totalDurationMs,omitempty"`
	Pauses          []TestWorkflowPause               `json:"pauses,omitempty"`
	Initialization  *TestWorkflowStepResult           `json:"initialization,omitempty"`
	Steps           map[string]TestWorkflowStepResult `json:"steps,omitempty"`
}

// TestWorkflowStatus has status of TestWorkflow
// +kubebuilder:validation:Enum=queued;running;paused;passed;failed;aborted
type TestWorkflowStatus string

// List of TestWorkflowStatus
const (
	QUEUED_TestWorkflowStatus  TestWorkflowStatus = "queued"
	RUNNING_TestWorkflowStatus TestWorkflowStatus = "running"
	PAUSED_TestWorkflowStatus  TestWorkflowStatus = "paused"
	PASSED_TestWorkflowStatus  TestWorkflowStatus = "passed"
	FAILED_TestWorkflowStatus  TestWorkflowStatus = "failed"
	ABORTED_TestWorkflowStatus TestWorkflowStatus = "aborted"
)

// TestWorkflowPause defines pause of TestWorkflow
type TestWorkflowPause struct {
	// step at which it was paused
	Ref string `json:"ref"`
	// when the pause has started
	PausedAt metav1.Time `json:"pausedAt"`
	// when the pause has ended
	ResumedAt metav1.Time `json:"resumedAt,omitempty"`
}

// TestWorkflowStepResult contains step result of TestWorkflow
type TestWorkflowStepResult struct {
	ErrorMessage string                  `json:"errorMessage,omitempty"`
	Status       *TestWorkflowStepStatus `json:"status,omitempty"`
	ExitCode     int64                   `json:"exitCode,omitempty"`
	// when the container was created
	QueuedAt metav1.Time `json:"queuedAt,omitempty"`
	// when the container was started
	StartedAt metav1.Time `json:"startedAt,omitempty"`
	// when the container was finished
	FinishedAt metav1.Time `json:"finishedAt,omitempty"`
}

// TestWorkfloStepwStatus has step status of TestWorkflow
type TestWorkflowStepStatus string

// List of TestWorkflowStepStatus
// +kubebuilder:validation:Enum=queued;running;paused;passed;failed;timeout;skipped;aborted
const (
	QUEUED_TestWorkflowStepStatus  TestWorkflowStepStatus = "queued"
	RUNNING_TestWorkflowStepStatus TestWorkflowStepStatus = "running"
	PAUSED_TestWorkflowStepStatus  TestWorkflowStepStatus = "paused"
	PASSED_TestWorkflowStepStatus  TestWorkflowStepStatus = "passed"
	FAILED_TestWorkflowStepStatus  TestWorkflowStepStatus = "failed"
	TIMEOUT_TestWorkflowStepStatus TestWorkflowStepStatus = "timeout"
	SKIPPED_TestWorkflowStepStatus TestWorkflowStepStatus = "skipped"
	ABORTED_TestWorkflowStepStatus TestWorkflowStepStatus = "aborted"
)

// TestWorkflowOutput defines output of TestWorkflow
type TestWorkflowOutput struct {
	// step reference
	Ref string `json:"ref,omitempty"`
	// output kind name
	Name string `json:"name,omitempty"`
	// value returned
	Value map[string]DynamicList `json:"value,omitempty"`
}

// TestWorkflowStepReport contains report of TestWorkflow
type TestWorkflowReport struct {
	// step reference
	Ref string `json:"ref,omitempty"`
	// report kind/type
	Kind string `json:"kind,omitempty"`
	// file path to full report in artifact storage
	File    string                     `json:"file,omitempty"`
	Summary *TestWorkflowReportSummary `json:"summary,omitempty"`
}

// TestWorkflowStepReportSummary contains report summary of TestWorkflow
type TestWorkflowReportSummary struct {
	// total number of test cases
	Tests int32 `json:"tests,omitempty"`
	// number of passed test cases
	Passed int32 `json:"passed,omitempty"`
	// number of failed test cases
	Failed int32 `json:"failed,omitempty"`
	// number of skipped test cases
	Skipped int32 `json:"skipped,omitempty"`
	// number of error test cases
	Errored int32 `json:"errored,omitempty"`
	// total duration of all test cases in milliseconds
	Duration int64 `json:"duration,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TestWorkflowExecution is the Schema for the workflows API
type TestWorkflowExecution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// TestWorkflowExecution specification
	Spec TestWorkflowExecutionSpec `json:"spec" expr:"include"`
	// TestWorkflowExecutionStatus specification
	Status TestWorkflowExecutionStatus `json:"status,omitempty" expr:"include"`
}

//+kubebuilder:object:root=true

// TestWorkflowExecutionList contains a list of TestWorkflowExecutiom
type TestWorkflowExecutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestWorkflowExecution `json:"items" expr:"include"`
}

func init() {
	SchemeBuilder.Register(&TestWorkflowExecution{}, &TestWorkflowExecutionList{})
}
