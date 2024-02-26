package device

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "dcim"
		r.References["device_type_id"] = config.Reference{
			Type:      "DeviceType",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["role_id"] = config.Reference{
			Type:      "DeviceRole",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["site_id"] = config.Reference{
			Type:      "Site",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		if s, ok := r.TerraformResource.Schema["device_type_id"]; ok {
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
	})
}
