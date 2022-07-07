package iam

import (
	"github.com/crossplane-contrib/provider-jet-hsdp/apis/rconfig"
	"github.com/crossplane/terrajet/pkg/config"
)

// PropositionConfigure configures individual resources by adding custom ResourceConfigurators.
func PropositionConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_proposition", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.References["organization_id"] = config.Reference{
			Type:         "Org",
			Extractor:    rconfig.ExtractResourceIDFuncPath,
			RefFieldName: "OrganizationRef",
		}
	})
}
