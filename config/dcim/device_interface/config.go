package device_interface

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators. `https://github.com/upbound/upjet/issues/139#issuecomment-1315083167`
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device_interface", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "netbox"
		// r.ExternalName = config.NameAsIdentifier
		r.Kind = "DeviceInterface"
		r.ExternalName.OmittedFields = []string{
			"id",
		}
	})
}