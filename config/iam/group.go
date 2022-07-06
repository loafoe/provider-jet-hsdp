package iam

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// GroupConfigure configures individual resources by adding custom ResourceConfigurators.
func GroupConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_group", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.ExternalName = config.IdentifierFromProvider
	})
}
