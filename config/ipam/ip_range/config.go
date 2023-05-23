package ip_range

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_ip_range", func(r *config.Resource) {
		// r.ExternalName = config.NameAsIdentifier
		r.Kind = "IPRange"
		// r.ShortGroup = "IPRange"
		r.ExternalName.OmittedFields = []string{
			"id",
		}
	})
}
