/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/fire-ant/provider-infoblox-nios/apis/dns/v1alpha1"
	v1alpha1ip "github.com/fire-ant/provider-infoblox-nios/apis/ip/v1alpha1"
	v1alpha1ipv4 "github.com/fire-ant/provider-infoblox-nios/apis/ipv4/v1alpha1"
	v1alpha1ipv6 "github.com/fire-ant/provider-infoblox-nios/apis/ipv6/v1alpha1"
	v1alpha1network "github.com/fire-ant/provider-infoblox-nios/apis/network/v1alpha1"
	v1alpha1apis "github.com/fire-ant/provider-infoblox-nios/apis/v1alpha1"
	v1beta1 "github.com/fire-ant/provider-infoblox-nios/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1ip.SchemeBuilder.AddToScheme,
		v1alpha1ipv4.SchemeBuilder.AddToScheme,
		v1alpha1ipv6.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
