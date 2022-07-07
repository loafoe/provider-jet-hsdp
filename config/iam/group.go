package iam

import (
	"github.com/crossplane-contrib/provider-jet-hsdp/apis/rconfig"
	"github.com/crossplane/terrajet/pkg/config"
)

// GroupConfigure configures individual resources by adding custom ResourceConfigurators.
func GroupConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.References["managing_organization"] = config.Reference{
			Type:         "Org",
			Extractor:    rconfig.ExtractResourceIDFuncPath,
			RefFieldName: "OrganizationRef",
		}
		r.References["services"] = config.Reference{
			Type:         "Service",
			Extractor:    rconfig.ExtractResourceIDFuncPath,
			RefFieldName: "ServiceRefs",
		}
		r.References["roles"] = config.Reference{
			Type:         "Role",
			Extractor:    rconfig.ExtractResourceIDFuncPath,
			RefFieldName: "RoleRefs",
		}
	})
}
