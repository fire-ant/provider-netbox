package primaryip

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_primary_ip", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "primaryip"
		r.Kind = "PrimaryIp"
		r.References["virtual_machine_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/virtualization/v1alpha1.Machine",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["ip_address_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/ipam/v1alpha1.IPAddress",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
	})
}
