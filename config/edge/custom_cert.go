package edge

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// CustomCertConfigure configures individual resources by adding custom ResourceConfigurators.
func CustomCertConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_custom_cert", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
	})
}
