/*


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

// JenkinsPluginSpec defines the desired state of JenkinsPlugin
type JenkinsPluginSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of JenkinsPlugin. Edit JenkinsPlugin_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// JenkinsPluginStatus defines the observed state of JenkinsPlugin
type JenkinsPluginStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// JenkinsPlugin is the Schema for the jenkinsplugins API
type JenkinsPlugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JenkinsPluginSpec   `json:"spec,omitempty"`
	Status JenkinsPluginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JenkinsPluginList contains a list of JenkinsPlugin
type JenkinsPluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JenkinsPlugin `json:"items"`
}

func init() {
	SchemeBuilder.Register(&JenkinsPlugin{}, &JenkinsPluginList{})
}
