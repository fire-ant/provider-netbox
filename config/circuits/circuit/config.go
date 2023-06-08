package circuit

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_circuit", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "circuits"
		r.References["provider_id"] = config.Reference{
			Type:      "CircuitProvider",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["type_id"] = config.Reference{
			Type:      "CircuitType",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1.Tenant",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		// if s, ok := r.TerraformResource.Schema["tenant_id"]; ok {
		// 	s.Type = schema.TypeInt
		// 	s.Optional = true
		// 	s.Computed = false
		// }
		// 	TypeInvalid ValueType = iota
		// 	TypeBool
		// 	TypeInt
		// 	TypeFloat
		// 	TypeString
		// 	TypeList
		// 	TypeMap
		// 	TypeSet
	})
}
