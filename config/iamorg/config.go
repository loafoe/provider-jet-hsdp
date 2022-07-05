package iamorg

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_org", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.ExternalName = config.IdentifierFromProvider
	})
}
