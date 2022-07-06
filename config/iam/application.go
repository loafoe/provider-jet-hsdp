package iam

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// ApplicationConfigure configures individual resources by adding custom ResourceConfigurators.
func ApplicationConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_application", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
	})
}
