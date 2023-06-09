package iprange

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_ip_range", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.Kind = "IPRange"
		r.ShortGroup = "ipam"
		r.References["role_id"] = config.Reference{
			Type: "IpamRole",
			// Type:      "github.com/fire-ant/provider-netbox/apis/ipam/v1alpha1.Role",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1.Tenant",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["vrf_id"] = config.Reference{
			Type: "Vrf",
			// Type:      "github.com/fire-ant/provider-netbox/apis/ipam/v1alpha1.Vrf",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
