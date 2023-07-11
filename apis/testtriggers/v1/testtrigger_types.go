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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TestTrigger is the Schema for the testtriggers API
// +kubebuilder:printcolumn:name="Resource",type=string,JSONPath=`.spec.resource`
// +kubebuilder:printcolumn:name="Event",type=string,JSONPath=`.spec.event`
// +kubebuilder:printcolumn:name="Execution",type=string,JSONPath=`.spec.execution`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type TestTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestTriggerSpec   `json:"spec,omitempty"`
	Status TestTriggerStatus `json:"status,omitempty"`
}

// TestTriggerSpec defines the desired state of TestTrigger
type TestTriggerSpec struct {
	// For which Resource do we monitor Event which triggers an Action on certain conditions
	Resource TestTriggerResource `json:"resource"`
	// ResourceSelector identifies which Kubernetes Objects should be watched
	ResourceSelector TestTriggerSelector               `json:"resourceSelector"`
	CustomResource   TestTriggerCRGroupVersionResource `json:"customResource,omitempty"`
	// On which Event for a Resource should an Action be triggered
	Event TestTriggerEvent `json:"event"`
	// What resource conditions should be matched
	ConditionSpec *TestTriggerConditionSpec `json:"conditionSpec,omitempty"`
	// What resource probes should be matched
	ProbeSpec *TestTriggerProbeSpec `json:"probeSpec,omitempty"`
	// Action represents what needs to be executed for selected Execution
	Action TestTriggerAction `json:"action"`
	// Execution identifies for which test execution should an Action be executed
	Execution TestTriggerExecution `json:"execution"`
	// TestSelector identifies on which Testkube Kubernetes Objects an Action should be taken
	TestSelector TestTriggerSelector `json:"testSelector"`
	// Delay is a duration string which specifies how long should the test be delayed after a trigger is matched
	// +kubebuilder:validation:Type:=string
	// +kubebuilder:validation:Format:=duration
	Delay *metav1.Duration `json:"delay,omitempty"`
}

type TestTriggerCRGroupVersionResource struct {
	Group    string `json:"group"`
	Version  string `json:"version"`
	Resource string `json:"resource"`
}

// TestTriggerResource defines resource for test triggers
// +kubebuilder:validation:Enum=pod;deployment;statefulset;daemonset;service;ingress;event;configmap;custom-resource
type TestTriggerResource string

// List of TestTriggerResources
const (
	TestTriggerResourcePod            TestTriggerResource = "pod"
	TestTriggerResourceDeployment     TestTriggerResource = "deployment"
	TestTriggerResourceStatefulSet    TestTriggerResource = "statefulset"
	TestTriggerResourceDaemonSet      TestTriggerResource = "daemonset"
	TestTriggerResourceService        TestTriggerResource = "service"
	TestTriggerResourceCustomResource TestTriggerResource = "custom-resource"
	TestTriggerResourceIngress        TestTriggerResource = "ingress"
	TestTriggerResourceEvent          TestTriggerResource = "event"
	TestTriggerResourceConfigMap      TestTriggerResource = "configmap"
)

// TestTriggerEvent defines event for test triggers
// +kubebuilder:validation:Enum=created;modified;deleted;deployment-scale-update;deployment-image-update;deployment-env-update;deployment-containers-modified
type TestTriggerEvent string

// List of TestTriggerEvents
const (
	TestTriggerEventCreated                      TestTriggerEvent = "created"
	TestTriggerEventModified                     TestTriggerEvent = "modified"
	TestTriggerEventDeleted                      TestTriggerEvent = "deleted"
	TestTriggerCauseDeploymentScaleUpdate        TestTriggerEvent = "deployment-scale-update"
	TestTriggerCauseDeploymentImageUpdate        TestTriggerEvent = "deployment-image-update"
	TestTriggerCauseDeploymentEnvUpdate          TestTriggerEvent = "deployment-env-update"
	TestTriggerCauseDeploymentContainersModified TestTriggerEvent = "deployment-containers-modified"
)

// TestTriggerAction defines action for test triggers
// +kubebuilder:validation:Enum=run
type TestTriggerAction string

// List of TestTriggerAction
const (
	TestTriggerActionRun TestTriggerAction = "run"
)

// TestTriggerExecution defines execution for test triggers
// +kubebuilder:validation:Enum=test;testsuite
type TestTriggerExecution string

// List of TestTriggerExecution
const (
	TestTriggerExecutionTest      TestTriggerExecution = "test"
	TestTriggerExecutionTestsuite TestTriggerExecution = "testsuite"
)

// TestTriggerSelector is used for selecting Kubernetes Objects
type TestTriggerSelector struct {
	// Name selector is used to identify a Kubernetes Object based on the metadata name
	Name string `json:"name,omitempty"`
	// Namespace of the Kubernetes object
	Namespace string `json:"namespace,omitempty"`
	// LabelSelector is used to identify a group of Kubernetes Objects based on their metadata labels
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty"`
}

// TestTriggerStatus defines the observed state of TestTrigger
type TestTriggerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// TestTriggerConditionStatuses defines condition statuses for test triggers
// +kubebuilder:validation:Enum=True;False;Unknown
type TestTriggerConditionStatuses string

// List of TestTriggerConditionStatuses
const (
	TRUE_TestTriggerConditionStatuses    TestTriggerConditionStatuses = "True"
	FALSE_TestTriggerConditionStatuses   TestTriggerConditionStatuses = "False"
	UNKNOWN_TestTriggerConditionStatuses TestTriggerConditionStatuses = "Unknown"
)

// TestTriggerCondition is used for definition of the condition for test triggers
type TestTriggerCondition struct {
	Status *TestTriggerConditionStatuses `json:"status"`
	// test trigger condition
	Type_ string `json:"type"`
	// test trigger condition reason
	Reason string `json:"reason,omitempty"`
	// duration in seconds in the past from current time when the condition is still valid
	Ttl int32 `json:"ttl,omitempty"`
}

// TestTriggerConditionSpec defines the condition specification for TestTrigger
type TestTriggerConditionSpec struct {
	// list of test trigger conditions
	Conditions []TestTriggerCondition `json:"conditions,omitempty"`
	// duration in seconds the test trigger waits for conditions, until its stopped
	Timeout int32 `json:"timeout,omitempty"`
	// duration in seconds the test trigger waits between condition check
	Delay int32 `json:"delay,omitempty"`
}

// TestTriggerProbe is used for definition of the probe for test triggers
type TestTriggerProbe struct {
	// test trigger condition probe scheme to connect to host, default is http
	Scheme string `json:"scheme,omitempty"`
	// test trigger condition probe host, default is pod ip or service name
	Host string `json:"host,omitempty"`
	// test trigger condition probe path to check, default is /
	Path string `json:"path,omitempty"`
	// test trigger condition probe port to connect
	Port int32 `json:"port,omitempty"`
	// test trigger condition probe headers to submit
	Headers map[string]string `json:"headers,omitempty"`
}

// TestTriggerProbeSpec defines the probe specification for TestTrigger
type TestTriggerProbeSpec struct {
	// list of test trigger probes
	Probes []TestTriggerProbe `json:"probes,omitempty"`
	// duration in seconds the test trigger waits for probes, until its stopped
	Timeout int32 `json:"timeout,omitempty"`
	// duration in seconds the test trigger waits between probes
	Delay int32 `json:"delay,omitempty"`
}

//+kubebuilder:object:root=true

// TestTriggerList contains a list of TestTrigger
type TestTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestTrigger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TestTrigger{}, &TestTriggerList{})
}
