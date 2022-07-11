package iam

import (
	"github.com/crossplane-contrib/provider-jet-hsdp/apis/rconfig"
	"github.com/crossplane/terrajet/pkg/config"
)

// PasswordPolicyConfigure configures individual resources by adding custom ResourceConfigurators.
func PasswordPolicyConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.References["managing_organization"] = config.Reference{
			Type:         "Org",
			Extractor:    rconfig.ExtractResourceIDFuncPath,
			RefFieldName: "OrganizationRef",
		}
	})
}
