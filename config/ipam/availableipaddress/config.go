package availableipaddress

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_available_ip_address", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "ipam"
		r.Kind = "AvailableIPAddress"
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1.Tenant",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["prefix_id"] = config.Reference{
			Type:      "Prefix",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["ip_range_id"] = config.Reference{
			Type:      "IPRange",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["interface_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.DeviceInterface",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["vrf_id"] = config.Reference{
			Type:      "Vrf",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		if s, ok := r.TerraformResource.Schema["prefix_id"]; ok {
			s.Type = schema.TypeString
			s.Optional = true
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["ip_range_id"]; ok {
			s.Type = schema.TypeString
			s.Optional = true
			s.Computed = true
		}
	})
}
