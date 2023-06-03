package deviceinterface

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators. `https://github.com/upbound/upjet/issues/139#issuecomment-1315083167`
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device_interface", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "netbox"
		// r.ExternalName = config.NameAsIdentifier
		r.Kind = "DeviceInterface"
		r.ShortGroup = "dcim"
		r.References["device_id"] = config.Reference{
			Type: "Device",
			// Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Device",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
