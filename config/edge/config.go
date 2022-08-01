package edge

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// ConfigConfigure configures individual resources by adding custom ResourceConfigurators.
func ConfigConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_edge_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
	})
}
