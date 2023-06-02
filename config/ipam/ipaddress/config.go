package ipaddress

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_ip_address", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.References["interface_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.VirtInterface",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["vlan_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Vlan",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["vrf_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Vrf",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
