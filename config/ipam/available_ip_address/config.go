package available_ip_address

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_available_ip_address", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Tenant",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["prefix_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Prefix",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["ip_range_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/ip/v1alpha1.IPRange",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["interface_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.VirtInterface",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["interface_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Vrf",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
