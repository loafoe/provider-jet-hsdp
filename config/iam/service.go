package iam

import (
	"github.com/crossplane-contrib/provider-jet-hsdp/apis/rconfig"
	"github.com/crossplane/terrajet/pkg/config"
)

// ServiceConfigure configures individual resources by adding custom ResourceConfigurators.
func ServiceConfigure(p *config.Provider) {
	p.AddResourceConfigurator("hsdp_iam_service", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.ExternalName = config.IdentifierFromProvider
		r.References["application_id"] = config.Reference{
			Type:         "Application",
			Extractor:    rconfig.ExtractResourceIDFuncPath,
			RefFieldName: "ApplicationRef",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["service_id"].(string); ok {
				conn["service_id"] = []byte(a)
			}
			if a, ok := attr["service_private_key"].(string); ok {
				conn["service_private_key"] = []byte(a)
			}
			return conn, nil
		}
	})
}
