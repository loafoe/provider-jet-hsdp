package iam

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// OrgConfigure configures individual resources by adding custom ResourceConfigurators.
func OrgConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_org", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.ExternalName = config.IdentifierFromProvider
	})
}
