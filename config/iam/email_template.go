package iam

import (
	"github.com/crossplane-contrib/provider-jet-hsdp/apis/rconfig"
	"github.com/crossplane/terrajet/pkg/config"
)

// EmailTemplateConfigure configures individual resources by adding custom ResourceConfigurators.
func EmailTemplateConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_email_template", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.References["managing_organization"] = config.Reference{
			Type:         "Org",
			Extractor:    rconfig.ExtractResourceIDFuncPath,
			RefFieldName: "OrganizationRef",
		}
	})
}
