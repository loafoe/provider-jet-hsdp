package iam

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// PropositionConfigure configures individual resources by adding custom ResourceConfigurators.
func PropositionConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_proposition", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
	})
}
