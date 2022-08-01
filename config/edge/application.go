package edge

import (
	"github.com/crossplane/terrajet/pkg/config"
)

// ApplicationConfigure configures individual resources by adding custom ResourceConfigurators.
func ApplicationConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_edge_app", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
	})
}
