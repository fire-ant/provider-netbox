package service

import (
	"github.com/fire-ant/provider-netbox/config/common"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_service", func(r *config.Resource) {
		r.ExternalName = config.NameAsIdentifier
		r.References["virtual_machine_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/virtual/v1alpha1.Machine",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})
}
