package device

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_device", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.References["device_type_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/device/v1alpha1.DeviceType",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["role_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/device/v1alpha1.Role",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["site_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Site",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		// r.LateInitializer = config.LateInitializer{
		// 	IgnoredFields: []string{"site_id"},
		// }
		if s, ok := r.TerraformResource.Schema["device_type_id"]; ok {
			s.Type = schema.TypeString
			s.Optional = false
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["role_id"]; ok {
			s.Type = schema.TypeString
			s.Optional = false
			s.Computed = true
		}
		if s, ok := r.TerraformResource.Schema["site_id"]; ok {
			s.Type = schema.TypeString
			s.Optional = false
			s.Computed = true
		}
	})
}
