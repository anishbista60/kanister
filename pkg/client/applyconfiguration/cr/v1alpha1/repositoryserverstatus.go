/*
Copyright 2024 The Kanister Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	crv1alpha1 "github.com/kanisterio/kanister/pkg/apis/cr/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RepositoryServerStatusApplyConfiguration represents an declarative configuration of the RepositoryServerStatus type for use
// with apply.
type RepositoryServerStatusApplyConfiguration struct {
	Conditions []v1.Condition                       `json:"conditions,omitempty"`
	ServerInfo *ServerInfoApplyConfiguration        `json:"serverInfo,omitempty"`
	Progress   *crv1alpha1.RepositoryServerProgress `json:"progress,omitempty"`
}

// RepositoryServerStatusApplyConfiguration constructs an declarative configuration of the RepositoryServerStatus type for use with
// apply.
func RepositoryServerStatus() *RepositoryServerStatusApplyConfiguration {
	return &RepositoryServerStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *RepositoryServerStatusApplyConfiguration) WithConditions(values ...v1.Condition) *RepositoryServerStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}

// WithServerInfo sets the ServerInfo field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServerInfo field is set to the value of the last call.
func (b *RepositoryServerStatusApplyConfiguration) WithServerInfo(value *ServerInfoApplyConfiguration) *RepositoryServerStatusApplyConfiguration {
	b.ServerInfo = value
	return b
}

// WithProgress sets the Progress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Progress field is set to the value of the last call.
func (b *RepositoryServerStatusApplyConfiguration) WithProgress(value crv1alpha1.RepositoryServerProgress) *RepositoryServerStatusApplyConfiguration {
	b.Progress = &value
	return b
}