package virtualmachine

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_virtual_machine", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "virtualization"
		r.References["cluster_id"] = config.Reference{
			Type:      "Cluster",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["platform_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.Platform",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["role_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.Role",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["site_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.Site",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["device_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.Device",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1.Tenant",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		if s, ok := r.TerraformResource.Schema["cluster_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["role_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["site_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["device_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["platform_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["tenant_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
	})
}
