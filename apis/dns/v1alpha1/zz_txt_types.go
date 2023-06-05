/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TXTObservation struct {

	// Description of the TXT-record.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// DNS view in which the record's zone exists.
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// Extensible attributes of the TXT-record to be added/updated, as a map in JSON format
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	// FQDN for the TXT-Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// TTL value of the TXT-Record
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Data to be associated with TXT_Record.
	Text *string `json:"text,omitempty" tf:"text,omitempty"`
}

type TXTParameters struct {

	// Description of the TXT-record.
	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// DNS view in which the record's zone exists.
	// +kubebuilder:validation:Optional
	DNSView *string `json:"dnsView,omitempty" tf:"dns_view,omitempty"`

	// Extensible attributes of the TXT-record to be added/updated, as a map in JSON format
	// +kubebuilder:validation:Optional
	ExtAttrs *string `json:"extAttrs,omitempty" tf:"ext_attrs,omitempty"`

	// FQDN for the TXT-Record.
	// +kubebuilder:validation:Optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// TTL value of the TXT-Record
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// Data to be associated with TXT_Record.
	// +kubebuilder:validation:Optional
	Text *string `json:"text,omitempty" tf:"text,omitempty"`
}

// TXTSpec defines the desired state of TXT
type TXTSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TXTParameters `json:"forProvider"`
}

// TXTStatus defines the observed state of TXT.
type TXTStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TXTObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TXT is the Schema for the TXTs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,infoblox-nios}
type TXT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.fqdn)",message="fqdn is a required parameter"
	Spec   TXTSpec   `json:"spec"`
	Status TXTStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TXTList contains a list of TXTs
type TXTList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TXT `json:"items"`
}

// Repository type metadata.
var (
	TXT_Kind             = "TXT"
	TXT_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TXT_Kind}.String()
	TXT_KindAPIVersion   = TXT_Kind + "." + CRDGroupVersion.String()
	TXT_GroupVersionKind = CRDGroupVersion.WithKind(TXT_Kind)
)

func init() {
	SchemeBuilder.Register(&TXT{}, &TXTList{})
}
