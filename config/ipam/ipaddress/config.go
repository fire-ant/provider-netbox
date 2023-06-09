package ipaddress

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_ip_address", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "ipam"
		r.Kind = "IPAddress"
		r.References["interface_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.DeviceInterface",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["vlan_id"] = config.Reference{
			Type: "Vlan",
			// Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Vlan",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["vrf_id"] = config.Reference{
			Type: "Vrf",
			// Type:      "github.com/fire-ant/provider-netbox/apis/ipam/v1alpha1.Vrf",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
