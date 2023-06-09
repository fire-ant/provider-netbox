package tenant

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_tenant", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.References["group_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1.Group",
			Extractor: common.ExtractResourceIDFuncPath,
		}

	})
}
