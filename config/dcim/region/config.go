package region

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_region", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.OmittedFields = []string{
			"id",
		}
	})
}
