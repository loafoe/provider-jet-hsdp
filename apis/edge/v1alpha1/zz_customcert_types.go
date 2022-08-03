/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CustomCertObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CustomCertParameters struct {

	// +kubebuilder:validation:Required
	CertPem *string `json:"certPem" tf:"cert_pem,omitempty"`

	// +kubebuilder:validation:Optional
	Principal []CustomCertPrincipalParameters `json:"principal,omitempty" tf:"principal,omitempty"`

	// +kubebuilder:validation:Required
	PrivateKeyPem *string `json:"privateKeyPem" tf:"private_key_pem,omitempty"`

	// +kubebuilder:validation:Required
	SerialNumber *string `json:"serialNumber" tf:"serial_number,omitempty"`

	// +kubebuilder:validation:Optional
	Sync *bool `json:"sync,omitempty" tf:"sync,omitempty"`
}

type CustomCertPrincipalObservation struct {
}

type CustomCertPrincipalParameters struct {

	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// +kubebuilder:validation:Optional
	Oauth2ClientID *string `json:"oauth2ClientId,omitempty" tf:"oauth2_client_id,omitempty"`

	// +kubebuilder:validation:Optional
	Oauth2PasswordSecretRef *v1.SecretKeySelector `json:"oauth2PasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServicePrivateKeySecretRef *v1.SecretKeySelector `json:"servicePrivateKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UaaPasswordSecretRef *v1.SecretKeySelector `json:"uaaPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	UaaUsername *string `json:"uaaUsername,omitempty" tf:"uaa_username,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// CustomCertSpec defines the desired state of CustomCert
type CustomCertSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomCertParameters `json:"forProvider"`
}

// CustomCertStatus defines the observed state of CustomCert.
type CustomCertStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomCertObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomCert is the Schema for the CustomCerts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdpjet}
type CustomCert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomCertSpec   `json:"spec"`
	Status            CustomCertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomCertList contains a list of CustomCerts
type CustomCertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomCert `json:"items"`
}

// Repository type metadata.
var (
	CustomCert_Kind             = "CustomCert"
	CustomCert_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomCert_Kind}.String()
	CustomCert_KindAPIVersion   = CustomCert_Kind + "." + CRDGroupVersion.String()
	CustomCert_GroupVersionKind = CRDGroupVersion.WithKind(CustomCert_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomCert{}, &CustomCertList{})
}
