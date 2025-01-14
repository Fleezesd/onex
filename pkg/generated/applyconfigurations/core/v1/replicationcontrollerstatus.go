// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ReplicationControllerStatusApplyConfiguration represents an declarative configuration of the ReplicationControllerStatus type for use
// with apply.
type ReplicationControllerStatusApplyConfiguration struct {
	Replicas             *int32                                             `json:"replicas,omitempty"`
	FullyLabeledReplicas *int32                                             `json:"fullyLabeledReplicas,omitempty"`
	ReadyReplicas        *int32                                             `json:"readyReplicas,omitempty"`
	AvailableReplicas    *int32                                             `json:"availableReplicas,omitempty"`
	ObservedGeneration   *int64                                             `json:"observedGeneration,omitempty"`
	Conditions           []ReplicationControllerConditionApplyConfiguration `json:"conditions,omitempty"`
}

// ReplicationControllerStatusApplyConfiguration constructs an declarative configuration of the ReplicationControllerStatus type for use with
// apply.
func ReplicationControllerStatus() *ReplicationControllerStatusApplyConfiguration {
	return &ReplicationControllerStatusApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *ReplicationControllerStatusApplyConfiguration) WithReplicas(value int32) *ReplicationControllerStatusApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithFullyLabeledReplicas sets the FullyLabeledReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FullyLabeledReplicas field is set to the value of the last call.
func (b *ReplicationControllerStatusApplyConfiguration) WithFullyLabeledReplicas(value int32) *ReplicationControllerStatusApplyConfiguration {
	b.FullyLabeledReplicas = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *ReplicationControllerStatusApplyConfiguration) WithReadyReplicas(value int32) *ReplicationControllerStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// WithAvailableReplicas sets the AvailableReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AvailableReplicas field is set to the value of the last call.
func (b *ReplicationControllerStatusApplyConfiguration) WithAvailableReplicas(value int32) *ReplicationControllerStatusApplyConfiguration {
	b.AvailableReplicas = &value
	return b
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *ReplicationControllerStatusApplyConfiguration) WithObservedGeneration(value int64) *ReplicationControllerStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ReplicationControllerStatusApplyConfiguration) WithConditions(values ...*ReplicationControllerConditionApplyConfiguration) *ReplicationControllerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
