package cluster

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_cluster", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "virtualization"
		r.References["cluster_group_id"] = config.Reference{
			Type:      "Group",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["cluster_type_id"] = config.Reference{
			Type:      "ClusterType",
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
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{
				"site_id",
				"tenant_id",
				"cluster_group_id",
				"cluster_type_id",
			},
		}
		if s, ok := r.TerraformResource.Schema["cluster_group_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["cluster_type_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["site_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["tenant_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
	})
}
