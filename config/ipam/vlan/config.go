package vlan

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_vlan", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "ipam"
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1.Tenant",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["site_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.Site",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["role_id"] = config.Reference{
			Type: "IpamRole",
			// Type:      "github.com/fire-ant/provider-netbox/apis/ipam/v1alpha1.Role",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["group_id"] = config.Reference{
			Type: "Group",
			// Type:      "github.com/fire-ant/provider-netbox/apis/vlan/v1alpha1.Group",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		// if s, ok := r.TerraformResource.Schema["tenant_id"]; ok {
		// 	s.Type = schema.TypeString
		// 	s.Optional = false
		// 	s.Computed = true
		// }
		// if s, ok := r.TerraformResource.Schema["site_id"]; ok {
		// 	s.Type = schema.TypeString
		// 	s.Optional = false
		// 	s.Computed = true
		// }
		// if s, ok := r.TerraformResource.Schema["role_id"]; ok {
		// 	s.Type = schema.TypeString
		// 	s.Optional = false
		// 	s.Computed = true
		// }
		// if s, ok := r.TerraformResource.Schema["group_id"]; ok {
		// 	s.Type = schema.TypeString
		// 	s.Optional = false
		// 	s.Computed = true
		// }
	})
}
