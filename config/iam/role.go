package iam

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// RoleConfigure configures individual resources by adding custom ResourceConfigurators.
func RoleConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_role", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
	})
}
