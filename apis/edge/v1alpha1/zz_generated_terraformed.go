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
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this App
func (mg *App) GetTerraformResourceType() string {
	return "hsdp_edge_app"
}

// GetConnectionDetailsMapping for this App
func (tr *App) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"principal[*].oauth2_password": "spec.forProvider.principal[*].oauth2PasswordSecretRef", "principal[*].password": "spec.forProvider.principal[*].passwordSecretRef", "principal[*].service_private_key": "spec.forProvider.principal[*].servicePrivateKeySecretRef", "principal[*].uaa_password": "spec.forProvider.principal[*].uaaPasswordSecretRef"}
}

// GetObservation of this App
func (tr *App) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this App
func (tr *App) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this App
func (tr *App) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this App
func (tr *App) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this App
func (tr *App) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this App using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *App) LateInitialize(attrs []byte) (bool, error) {
	params := &AppParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *App) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this Config
func (mg *Config) GetTerraformResourceType() string {
	return "hsdp_edge_config"
}

// GetConnectionDetailsMapping for this Config
func (tr *Config) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"principal[*].oauth2_password": "spec.forProvider.principal[*].oauth2PasswordSecretRef", "principal[*].password": "spec.forProvider.principal[*].passwordSecretRef", "principal[*].service_private_key": "spec.forProvider.principal[*].servicePrivateKeySecretRef", "principal[*].uaa_password": "spec.forProvider.principal[*].uaaPasswordSecretRef"}
}

// GetObservation of this Config
func (tr *Config) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Config
func (tr *Config) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Config
func (tr *Config) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Config
func (tr *Config) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Config
func (tr *Config) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Config using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Config) LateInitialize(attrs []byte) (bool, error) {
	params := &ConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Config) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this CustomCert
func (mg *CustomCert) GetTerraformResourceType() string {
	return "hsdp_edge_custom_cert"
}

// GetConnectionDetailsMapping for this CustomCert
func (tr *CustomCert) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"principal[*].oauth2_password": "spec.forProvider.principal[*].oauth2PasswordSecretRef", "principal[*].password": "spec.forProvider.principal[*].passwordSecretRef", "principal[*].service_private_key": "spec.forProvider.principal[*].servicePrivateKeySecretRef", "principal[*].uaa_password": "spec.forProvider.principal[*].uaaPasswordSecretRef"}
}

// GetObservation of this CustomCert
func (tr *CustomCert) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this CustomCert
func (tr *CustomCert) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this CustomCert
func (tr *CustomCert) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this CustomCert
func (tr *CustomCert) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this CustomCert
func (tr *CustomCert) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this CustomCert using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *CustomCert) LateInitialize(attrs []byte) (bool, error) {
	params := &CustomCertParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *CustomCert) GetTerraformSchemaVersion() int {
	return 1
}
