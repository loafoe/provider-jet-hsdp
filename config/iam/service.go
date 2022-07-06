package iam

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// ServiceConfigure configures individual resources by adding custom ResourceConfigurators.
func ServiceConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
	})
}
