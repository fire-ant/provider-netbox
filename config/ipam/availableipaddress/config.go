package availableipaddress

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_available_ip_address", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "ipam"
		r.Kind = "AvailableIPAdrress"
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1.Tenant",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["prefix_id"] = config.Reference{
			Type: "Prefix",
			// Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Prefix",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["ip_range_id"] = config.Reference{
			Type: "IPRange",
			// Type:      "github.com/fire-ant/provider-netbox/apis/ip/v1alpha1.IPRange",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["interface_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.DeviceInterface",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["vrf_id"] = config.Reference{
			Type: "Vrf",
			// Type:      "github.com/fire-ant/provider-netbox/apis/ipam/v1alpha1.Vrf",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
