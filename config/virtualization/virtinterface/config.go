package virtinterface

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_interface", func(r *config.Resource) {
		r.Kind = "VirtInterface"
		r.ExternalName = config.NameAsIdentifier
		r.ShortGroup = "virtualization"
		r.References["virtual_machine_id"] = config.Reference{
			Type:      "Machine",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		if s, ok := r.TerraformResource.Schema["virtual_machine_id"]; ok {
			s.Optional = true
			s.Computed = true
		}
	})
}
