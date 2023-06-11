package deviceinterface

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators. `https://github.com/upbound/upjet/issues/139#issuecomment-1315083167`
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device_interface", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.Kind = "DeviceInterface"
		r.ShortGroup = "dcim"
		r.References["device_id"] = config.Reference{
			Type:      "Device",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
