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

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane-contrib/provider-jet-hsdp/config/iam"
	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resourcePrefix = "hsdp"
	modulePath     = "github.com/crossplane-contrib/provider-jet-hsdp"
)

//go:embed schema.json
var providerSchema string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList([]string{
			"hsdp_iam_group$",
			"hsdp_iam_org$",
			"hsdp_iam_role$",
			"hsdp_iam_service$",
			"hsdp_iam_proposition$",
			"hsdp_iam_application$",
			"hsdp_iam_client$",
			"hsdp_iam_password_policy$",
			"hsdp_iam_role_sharing_policy$",
			"hsdp_iam_email_template$",
		}))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		iam.GroupConfigure,
		iam.RoleConfigure,
		iam.OrgConfigure,
		iam.ServiceConfigure,
		iam.PropositionConfigure,
		iam.ApplicationConfigure,
		iam.ClientConfigure,
		iam.PasswordPolicyConfigure,
		iam.RoleSharingPolicyConfigure,
		iam.EmailTemplateConfigure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
