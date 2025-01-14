//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/onex.

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	config "github.com/superproj/onex/internal/controller/minerset/apis/config"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha1 "k8s.io/component-base/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*MinerSetControllerConfiguration)(nil), (*config.MinerSetControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MinerSetControllerConfiguration_To_config_MinerSetControllerConfiguration(a.(*MinerSetControllerConfiguration), b.(*config.MinerSetControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MinerSetControllerConfiguration)(nil), (*MinerSetControllerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MinerSetControllerConfiguration_To_v1beta1_MinerSetControllerConfiguration(a.(*config.MinerSetControllerConfiguration), b.(*MinerSetControllerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_MinerSetControllerConfiguration_To_config_MinerSetControllerConfiguration(in *MinerSetControllerConfiguration, out *config.MinerSetControllerConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.Parallelism = in.Parallelism
	out.SyncPeriod = in.SyncPeriod
	out.WatchFilterValue = in.WatchFilterValue
	if err := v1alpha1.Convert_v1alpha1_LeaderElectionConfiguration_To_config_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	out.MetricsBindAddress = in.MetricsBindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	return nil
}

// Convert_v1beta1_MinerSetControllerConfiguration_To_config_MinerSetControllerConfiguration is an autogenerated conversion function.
func Convert_v1beta1_MinerSetControllerConfiguration_To_config_MinerSetControllerConfiguration(in *MinerSetControllerConfiguration, out *config.MinerSetControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_MinerSetControllerConfiguration_To_config_MinerSetControllerConfiguration(in, out, s)
}

func autoConvert_config_MinerSetControllerConfiguration_To_v1beta1_MinerSetControllerConfiguration(in *config.MinerSetControllerConfiguration, out *MinerSetControllerConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	out.Parallelism = in.Parallelism
	out.SyncPeriod = in.SyncPeriod
	out.WatchFilterValue = in.WatchFilterValue
	if err := v1alpha1.Convert_config_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	out.Namespace = in.Namespace
	out.MetricsBindAddress = in.MetricsBindAddress
	out.HealthzBindAddress = in.HealthzBindAddress
	return nil
}

// Convert_config_MinerSetControllerConfiguration_To_v1beta1_MinerSetControllerConfiguration is an autogenerated conversion function.
func Convert_config_MinerSetControllerConfiguration_To_v1beta1_MinerSetControllerConfiguration(in *config.MinerSetControllerConfiguration, out *MinerSetControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_MinerSetControllerConfiguration_To_v1beta1_MinerSetControllerConfiguration(in, out, s)
}
