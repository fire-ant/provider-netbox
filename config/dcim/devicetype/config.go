package devicetype

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device_type", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "DeviceType"
		r.ShortGroup = "dcim"
		r.References["manufacturer_id"] = config.Reference{
			Type:      "Manufacturer",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}

	})
}
