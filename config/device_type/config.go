package device_type

import "github.com/upbound/upjet/pkg/config"

// https://github.com/upbound/upjet/issues/139#issuecomment-1315083167
// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device_type", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "netbox"
		// r.ExternalName = config.NameAsIdentifier
		r.Kind = "DeviceType"
		// r.ShortGroup = "DeviceType"
	})
}
