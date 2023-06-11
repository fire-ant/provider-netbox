package ipaddress

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/upbound/upjet/pkg/config"
)

const (
	// SelfPackagePath is the golang path for this package.
	SelfPackagePath = "github.com/ipam/ipaddress"
)

var (
	// PathInstanceGroupExtractor is the golang path to InstanceGroupExtractor function
	// in this package.
	PathObjectTypeSelector = SelfPackagePath + ".ObjectTypeSelector()"
)

func ObjectTypeSelector() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		paved, err := fieldpath.PaveObject(mg)
		if err != nil {
			return ""
		}
		var ot string
		r, err := paved.GetString("spec.forProvider.objectType")
		if err != nil {
			switch r {
			case "dcim.interface":
				ot = "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.DeviceInterface"
			case "virtualization.vminterface":
				ot = "github.com/fire-ant/provider-netbox/apis/virtualization/v1alpha1.VirtInterface"
			}
		}
		return ot
	}
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_ip_address", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "ipam"
		r.Kind = "IPAddress"
		r.References["interface_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/dcim/v1alpha1.DeviceInterface",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["tenant_id"] = config.Reference{
			Type:      "github.com/fire-ant/provider-netbox/apis/tenant/v1alpha1.Tenant",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		r.References["vrf_id"] = config.Reference{
			Type:      "Vrf",
			Extractor: "github.com/upbound/upjet/pkg/resource.ExtractResourceID()",
		}
		config.MarkAsRequired(r.TerraformResource, "object_type")
	})
}
