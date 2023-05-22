package manufacturer

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_manufacturer", func(r *config.Resource) {
		r.ShortGroup = "manufacturer"
	})
}
