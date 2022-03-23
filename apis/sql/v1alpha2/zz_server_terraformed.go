/*
Copyright 2022 The Crossplane Authors.

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

package v1alpha2

import (
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this MSSQLServer
func (mg *MSSQLServer) GetTerraformResourceType() string {
	return "azurerm_mssql_server"
}

// GetConnectionDetailsMapping for this MSSQLServer
func (tr *MSSQLServer) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"administrator_login_password": "spec.forProvider.administratorLoginPasswordSecretRef", "extended_auditing_policy[*].storage_account_access_key": "spec.forProvider.extendedAuditingPolicy[*].storageAccountAccessKeySecretRef"}
}

// GetObservation of this MSSQLServer
func (tr *MSSQLServer) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MSSQLServer
func (tr *MSSQLServer) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MSSQLServer
func (tr *MSSQLServer) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MSSQLServer
func (tr *MSSQLServer) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MSSQLServer
func (tr *MSSQLServer) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MSSQLServer using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MSSQLServer) LateInitialize(attrs []byte) (bool, error) {
	params := &MSSQLServerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MSSQLServer) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MSSQLServerTransparentDataEncryption
func (mg *MSSQLServerTransparentDataEncryption) GetTerraformResourceType() string {
	return "azurerm_mssql_server_transparent_data_encryption"
}

// GetConnectionDetailsMapping for this MSSQLServerTransparentDataEncryption
func (tr *MSSQLServerTransparentDataEncryption) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MSSQLServerTransparentDataEncryption
func (tr *MSSQLServerTransparentDataEncryption) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MSSQLServerTransparentDataEncryption
func (tr *MSSQLServerTransparentDataEncryption) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MSSQLServerTransparentDataEncryption
func (tr *MSSQLServerTransparentDataEncryption) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MSSQLServerTransparentDataEncryption
func (tr *MSSQLServerTransparentDataEncryption) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MSSQLServerTransparentDataEncryption
func (tr *MSSQLServerTransparentDataEncryption) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MSSQLServerTransparentDataEncryption using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MSSQLServerTransparentDataEncryption) LateInitialize(attrs []byte) (bool, error) {
	params := &MSSQLServerTransparentDataEncryptionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MSSQLServerTransparentDataEncryption) GetTerraformSchemaVersion() int {
	return 0
}
