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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta2

import (
	v1beta2 "k8s.io/api/apps/v1beta2"
)

// DaemonSetUpdateStrategyApplyConfiguration represents an declarative configuration of the DaemonSetUpdateStrategy type for use
// with apply.
type DaemonSetUpdateStrategyApplyConfiguration struct {
	Type          *v1beta2.DaemonSetUpdateStrategyType      `json:"type,omitempty"`
	RollingUpdate *RollingUpdateDaemonSetApplyConfiguration `json:"rollingUpdate,omitempty"`
}

// DaemonSetUpdateStrategyApplyConfiguration constructs an declarative configuration of the DaemonSetUpdateStrategy type for use with
// apply.
func DaemonSetUpdateStrategy() *DaemonSetUpdateStrategyApplyConfiguration {
	return &DaemonSetUpdateStrategyApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *DaemonSetUpdateStrategyApplyConfiguration) WithType(value v1beta2.DaemonSetUpdateStrategyType) *DaemonSetUpdateStrategyApplyConfiguration {
	b.Type = &value
	return b
}

// WithRollingUpdate sets the RollingUpdate field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RollingUpdate field is set to the value of the last call.
func (b *DaemonSetUpdateStrategyApplyConfiguration) WithRollingUpdate(value *RollingUpdateDaemonSetApplyConfiguration) *DaemonSetUpdateStrategyApplyConfiguration {
	b.RollingUpdate = value
	return b
}
