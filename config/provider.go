/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/fire-ant/provider-netbox/config/circuits/circuit"
	"github.com/fire-ant/provider-netbox/config/circuits/circuitprovider"
	"github.com/fire-ant/provider-netbox/config/circuits/circuittermination"
	"github.com/fire-ant/provider-netbox/config/circuits/circuittype"
	"github.com/fire-ant/provider-netbox/config/dcim/device"
	"github.com/fire-ant/provider-netbox/config/dcim/deviceinterface"
	"github.com/fire-ant/provider-netbox/config/dcim/devicerole"
	"github.com/fire-ant/provider-netbox/config/dcim/devicetype"
	"github.com/fire-ant/provider-netbox/config/dcim/location"
	"github.com/fire-ant/provider-netbox/config/dcim/manufacturer"
	"github.com/fire-ant/provider-netbox/config/dcim/platform"
	"github.com/fire-ant/provider-netbox/config/dcim/rack"
	"github.com/fire-ant/provider-netbox/config/dcim/rackreservation"
	"github.com/fire-ant/provider-netbox/config/dcim/rackrole"
	"github.com/fire-ant/provider-netbox/config/dcim/region"
	"github.com/fire-ant/provider-netbox/config/dcim/site"
	"github.com/fire-ant/provider-netbox/config/dcim/sitegroup"
	"github.com/fire-ant/provider-netbox/config/extras/customfield"
	"github.com/fire-ant/provider-netbox/config/extras/tag"
	"github.com/fire-ant/provider-netbox/config/ipam/aggregate"
	"github.com/fire-ant/provider-netbox/config/ipam/asn"
	"github.com/fire-ant/provider-netbox/config/ipam/availableipaddress"
	"github.com/fire-ant/provider-netbox/config/ipam/availableprefix"
	"github.com/fire-ant/provider-netbox/config/ipam/ipaddress"
	"github.com/fire-ant/provider-netbox/config/ipam/ipamrole"
	"github.com/fire-ant/provider-netbox/config/ipam/iprange"
	"github.com/fire-ant/provider-netbox/config/ipam/prefix"
	"github.com/fire-ant/provider-netbox/config/ipam/rir"
	"github.com/fire-ant/provider-netbox/config/ipam/routetarget"
	"github.com/fire-ant/provider-netbox/config/ipam/service"
	"github.com/fire-ant/provider-netbox/config/ipam/vlan"
	"github.com/fire-ant/provider-netbox/config/ipam/vlangroup"
	"github.com/fire-ant/provider-netbox/config/ipam/vrf"
	"github.com/fire-ant/provider-netbox/config/primaryip/primaryip"
	"github.com/fire-ant/provider-netbox/config/tenancy/contact"
	"github.com/fire-ant/provider-netbox/config/tenancy/contactassignment"
	"github.com/fire-ant/provider-netbox/config/tenancy/contactgroup"
	"github.com/fire-ant/provider-netbox/config/tenancy/contactrole"
	"github.com/fire-ant/provider-netbox/config/tenancy/tenant"
	"github.com/fire-ant/provider-netbox/config/tenancy/tenantgroup"
	"github.com/fire-ant/provider-netbox/config/virtualization/cluster"
	"github.com/fire-ant/provider-netbox/config/virtualization/clustergroup"
	"github.com/fire-ant/provider-netbox/config/virtualization/clustertype"
	"github.com/fire-ant/provider-netbox/config/virtualization/virtinterface"
	"github.com/fire-ant/provider-netbox/config/virtualization/virtualmachine"
)

const (
	resourcePrefix = "netbox"
	modulePath     = "github.com/fire-ant/provider-netbox"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		circuittype.Configure,
		circuit.Configure,
		circuitprovider.Configure,
		circuittermination.Configure,
		device.Configure,
		devicerole.Configure,
		devicetype.Configure,
		deviceinterface.Configure,
		location.Configure,
		manufacturer.Configure,
		platform.Configure,
		rack.Configure,
		rackreservation.Configure,
		rackrole.Configure,
		region.Configure,
		site.Configure,
		sitegroup.Configure,
		customfield.Configure,
		tag.Configure,
		aggregate.Configure,
		asn.Configure,
		availableipaddress.Configure,
		availableprefix.Configure,
		ipaddress.Configure,
		iprange.Configure,
		ipamrole.Configure,
		prefix.Configure,
		rir.Configure,
		routetarget.Configure,
		service.Configure,
		vlan.Configure,
		vlangroup.Configure,
		vrf.Configure,
		contact.Configure,
		contactassignment.Configure,
		contactgroup.Configure,
		contactrole.Configure,
		tenant.Configure,
		tenantgroup.Configure,
		cluster.Configure,
		clustergroup.Configure,
		clustertype.Configure,
		virtinterface.Configure,
		primaryip.Configure,
		virtualmachine.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
