package iprange

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_ip_range", func(r *config.Resource) {
		// r.ExternalName = config.NameAsIdentifier
		r.Kind = "IPRange"
		r.References["role_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/ipam/v1alpha1.Role",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Tenant",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["vrf_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Vrf",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
