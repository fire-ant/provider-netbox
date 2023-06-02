package devicetype

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device_type", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "netbox"
		// r.ExternalName = config.NameAsIdentifier
		r.Kind = "DeviceType"
		r.References["device_type_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Manufacturer",
			Extractor: common.ExtractResourceIDFuncPath,
		}

	})
}
