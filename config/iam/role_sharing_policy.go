package iam

import (
	"github.com/crossplane-contrib/provider-jet-hsdp/apis/rconfig"
	"github.com/crossplane/terrajet/pkg/config"
)

// RoleSharingPolicyConfigure configures individual resources by adding custom ResourceConfigurators.
func RoleSharingPolicyConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_role_sharing_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.References["role_id"] = config.Reference{
			Type:         "Role",
			Extractor:    rconfig.ExtractResourceIDFuncPath,
			RefFieldName: "RoleRef",
		}
		r.References["target_organization_id"] = config.Reference{
			Type:         "Org",
			Extractor:    rconfig.ExtractResourceIDFuncPath,
			RefFieldName: "OrganizationRef",
		}
	})
}
