// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PriorityLevelConfigurationReferenceApplyConfiguration represents an declarative configuration of the PriorityLevelConfigurationReference type for use
// with apply.
type PriorityLevelConfigurationReferenceApplyConfiguration struct {
	Name *string `json:"name,omitempty"`
}

// PriorityLevelConfigurationReferenceApplyConfiguration constructs an declarative configuration of the PriorityLevelConfigurationReference type for use with
// apply.
func PriorityLevelConfigurationReference() *PriorityLevelConfigurationReferenceApplyConfiguration {
	return &PriorityLevelConfigurationReferenceApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PriorityLevelConfigurationReferenceApplyConfiguration) WithName(value string) *PriorityLevelConfigurationReferenceApplyConfiguration {
	b.Name = &value
	return b
}
