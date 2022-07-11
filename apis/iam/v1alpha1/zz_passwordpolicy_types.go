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

type ChallengePolicyObservation struct {
}

type ChallengePolicyParameters struct {

	// +kubebuilder:validation:Optional
	DefaultQuestions []*string `json:"defaultQuestions,omitempty" tf:"default_questions,omitempty"`

	// +kubebuilder:validation:Optional
	MaxIncorrectAttempts *float64 `json:"maxIncorrectAttempts,omitempty" tf:"max_incorrect_attempts,omitempty"`

	// +kubebuilder:validation:Optional
	MinAnswerCount *float64 `json:"minAnswerCount,omitempty" tf:"min_answer_count,omitempty"`

	// +kubebuilder:validation:Optional
	MinQuestionCount *float64 `json:"minQuestionCount,omitempty" tf:"min_question_count,omitempty"`
}

type ComplexityObservation struct {
}

type ComplexityParameters struct {

	// +kubebuilder:validation:Optional
	MaxLength *float64 `json:"maxLength,omitempty" tf:"max_length,omitempty"`

	// +kubebuilder:validation:Optional
	MinLength *float64 `json:"minLength,omitempty" tf:"min_length,omitempty"`

	// +kubebuilder:validation:Optional
	MinLowercase *float64 `json:"minLowercase,omitempty" tf:"min_lowercase,omitempty"`

	// +kubebuilder:validation:Optional
	MinNumerics *float64 `json:"minNumerics,omitempty" tf:"min_numerics,omitempty"`

	// +kubebuilder:validation:Optional
	MinSpecialChars *float64 `json:"minSpecialChars,omitempty" tf:"min_special_chars,omitempty"`

	// +kubebuilder:validation:Optional
	MinUppercase *float64 `json:"minUppercase,omitempty" tf:"min_uppercase,omitempty"`
}

type PasswordPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Policy *string `json:"Policy,omitempty" tf:"_policy,omitempty"`
}

type PasswordPolicyParameters struct {

	// +kubebuilder:validation:Optional
	ChallengePolicy []ChallengePolicyParameters `json:"challengePolicy,omitempty" tf:"challenge_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ChallengesEnabled *bool `json:"challengesEnabled,omitempty" tf:"challenges_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Complexity []ComplexityParameters `json:"complexity" tf:"complexity,omitempty"`

	// +kubebuilder:validation:Optional
	ExpiryPeriodInDays *float64 `json:"expiryPeriodInDays,omitempty" tf:"expiry_period_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	HistoryCount *float64 `json:"historyCount,omitempty" tf:"history_count,omitempty"`

	// +kubebuilder:validation:Required
	ManagingOrganization *string `json:"managingOrganization" tf:"managing_organization,omitempty"`
}

// PasswordPolicySpec defines the desired state of PasswordPolicy
type PasswordPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PasswordPolicyParameters `json:"forProvider"`
}

// PasswordPolicyStatus defines the observed state of PasswordPolicy.
type PasswordPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PasswordPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PasswordPolicy is the Schema for the PasswordPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdpjet}
type PasswordPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PasswordPolicySpec   `json:"spec"`
	Status            PasswordPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PasswordPolicyList contains a list of PasswordPolicys
type PasswordPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PasswordPolicy `json:"items"`
}

// Repository type metadata.
var (
	PasswordPolicy_Kind             = "PasswordPolicy"
	PasswordPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PasswordPolicy_Kind}.String()
	PasswordPolicy_KindAPIVersion   = PasswordPolicy_Kind + "." + CRDGroupVersion.String()
	PasswordPolicy_GroupVersionKind = CRDGroupVersion.WithKind(PasswordPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&PasswordPolicy{}, &PasswordPolicyList{})
}
