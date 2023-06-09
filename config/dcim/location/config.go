package location

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_location", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.References["site_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Site",
			Extractor: common.ExtractResourceIDFuncPath,
		}
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/netbox/v1alpha1.Tenant",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
