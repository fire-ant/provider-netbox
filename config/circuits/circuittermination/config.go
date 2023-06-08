package circuittermination

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_circuit_termination", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "circuits"
		r.References["site_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.Site",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["circuit_id"] = config.Reference{
			Type:      "Circuit",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
			// Extractor:         "github.com/upbound/upjet/pkg/resource.ExtractParamPath(\"cid\", false)",
		}
	})
}
