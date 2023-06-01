package asn

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_asn", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.References["group_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Rir",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
