package tenant_group

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_tenant_group", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier

	})
}
