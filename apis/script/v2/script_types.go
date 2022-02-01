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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ScriptSpec defines the desired state of Script
type ScriptSpec struct {
	// script type
	Type_ string `json:"type,omitempty"`
	// script execution custom name
	Name string `json:"name,omitempty"`
	// execution params passed to executor
	Params map[string]string `json:"params,omitempty"`
	// script content object
	Content *ScriptContent `json:"content,omitempty"`
	// script tags
	Tags []string `json:"tags,omitempty"`
}

type ScriptContent struct {
	// script type
	Type_ string `json:"type,omitempty"`
	// repository of script content
	Repository *Repository `json:"repository,omitempty"`
	// script content body
	Data string `json:"data,omitempty"`
	// uri of script content
	Uri string `json:"uri,omitempty"`
}

// Repository represents VCS repo, currently we're habdling Git only
type Repository struct {
	// VCS repository type
	Type_ string `json:"type"`
	// uri of content file or git directory
	Uri string `json:"uri"`
	// branch/tag name for checkout
	Branch string `json:"branch"`
	// if needed we can checkout particular path (dir or file) in case of BIG/mono repositories
	Path string `json:"path,omitempty"`
	// git auth username for private repositories
	Username string `json:"username,omitempty"`
	// git auth token for private repositories
	Token string `json:"token,omitempty"`
}

// ScriptStatus defines the observed state of Script
type ScriptStatus struct {
	LastExecution   metav1.Time `json:"last_execution,omitempty"`
	ExecutionsCount int         `json:"executions_count,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// Script is the Schema for the scripts API
type Script struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScriptSpec   `json:"spec,omitempty"`
	Status ScriptStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ScriptList contains a list of Script
type ScriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Script `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Script{}, &ScriptList{})
}
