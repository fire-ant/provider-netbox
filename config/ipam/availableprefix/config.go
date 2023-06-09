package availableprefix

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_available_prefix", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "ipam"
		r.Kind = "AvailablePrefix"
		r.References["role_id"] = config.Reference{
			Type:      "IpamRole",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["site_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.Site",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1.Tenant",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["vlan_id"] = config.Reference{
			Type:      "Vlan",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["vrf_id"] = config.Reference{
			Type:      "Vrf",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["parent_prefix_id"] = config.Reference{
			Type:      "Prefix",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		if s, ok := r.TerraformResource.Schema["parent_prefix_id"]; ok {
			s.Type = schema.TypeFloat
			s.Optional = true
			s.Computed = true
		}
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"parent_prefix_id"},
		}
	})
}
