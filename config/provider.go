/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/fire-ant/provider-netbox/config/circuits/circuit"
	"github.com/fire-ant/provider-netbox/config/circuits/circuit_provider"
	"github.com/fire-ant/provider-netbox/config/circuits/circuit_termination"
	"github.com/fire-ant/provider-netbox/config/circuits/circuit_type"
	"github.com/fire-ant/provider-netbox/config/dcim/device"
	"github.com/fire-ant/provider-netbox/config/dcim/device_interface"
	"github.com/fire-ant/provider-netbox/config/dcim/device_role"
	"github.com/fire-ant/provider-netbox/config/dcim/device_type"
	"github.com/fire-ant/provider-netbox/config/dcim/location"
	"github.com/fire-ant/provider-netbox/config/dcim/manufacturer"
	"github.com/fire-ant/provider-netbox/config/dcim/platform"
	"github.com/fire-ant/provider-netbox/config/dcim/rack"
	"github.com/fire-ant/provider-netbox/config/dcim/rack_reservation"
	"github.com/fire-ant/provider-netbox/config/dcim/rack_role"
	"github.com/fire-ant/provider-netbox/config/dcim/region"
	"github.com/fire-ant/provider-netbox/config/dcim/site"
	"github.com/fire-ant/provider-netbox/config/dcim/site_group"
	"github.com/fire-ant/provider-netbox/config/extras/custom_field"
	"github.com/fire-ant/provider-netbox/config/extras/tag"
	"github.com/fire-ant/provider-netbox/config/ipam/aggregate"
	"github.com/fire-ant/provider-netbox/config/ipam/asn"
	"github.com/fire-ant/provider-netbox/config/ipam/available_ip_address"
	"github.com/fire-ant/provider-netbox/config/ipam/available_prefix"
	"github.com/fire-ant/provider-netbox/config/ipam/ip_address"
	"github.com/fire-ant/provider-netbox/config/ipam/ip_range"
	"github.com/fire-ant/provider-netbox/config/ipam/ipam_role"
	"github.com/fire-ant/provider-netbox/config/ipam/prefix"
	"github.com/fire-ant/provider-netbox/config/ipam/rir"
	"github.com/fire-ant/provider-netbox/config/ipam/route_target"
	"github.com/fire-ant/provider-netbox/config/ipam/service"
	"github.com/fire-ant/provider-netbox/config/ipam/vlan"
	"github.com/fire-ant/provider-netbox/config/ipam/vlan_group"
	"github.com/fire-ant/provider-netbox/config/ipam/vrf"
	"github.com/fire-ant/provider-netbox/config/tenancy/contact"
	"github.com/fire-ant/provider-netbox/config/tenancy/contact_assignment"
	"github.com/fire-ant/provider-netbox/config/tenancy/contact_group"
	"github.com/fire-ant/provider-netbox/config/tenancy/tenant"
	"github.com/fire-ant/provider-netbox/config/tenancy/tenant_group"
	"github.com/fire-ant/provider-netbox/config/tenancy/tenants"
	"github.com/fire-ant/provider-netbox/config/virtualization/cluster"
	"github.com/fire-ant/provider-netbox/config/virtualization/cluster_group"
	"github.com/fire-ant/provider-netbox/config/virtualization/cluster_type"
	"github.com/fire-ant/provider-netbox/config/virtualization/primary_ip"
	"github.com/fire-ant/provider-netbox/config/virtualization/virt_interface"
	"github.com/fire-ant/provider-netbox/config/virtualization/virtual_machine"
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
		circuit_type.Configure,
		circuit.Configure,
		circuit_provider.Configure,
		circuit_termination.Configure,
		device.Configure,
		device_role.Configure,
		device_type.Configure,
		device_interface.Configure,
		location.Configure,
		manufacturer.Configure,
		platform.Configure,
		rack.Configure,
		rack_reservation.Configure,
		rack_role.Configure,
		region.Configure,
		site.Configure,
		site_group.Configure,
		custom_field.Configure,
		tag.Configure,
		aggregate.Configure,
		asn.Configure,
		available_ip_address.Configure,
		available_prefix.Configure,
		ip_address.Configure,
		ip_range.Configure,
		ipam_role.Configure,
		prefix.Configure,
		rir.Configure,
		route_target.Configure,
		service.Configure,
		vlan.Configure,
		vlan_group.Configure,
		vrf.Configure,
		contact.Configure,
		contact_assignment.Configure,
		contact_group.Configure,
		tenant.Configure,
		tenant_group.Configure,
		tenants.Configure,
		cluster.Configure,
		cluster_group.Configure,
		cluster_type.Configure,
		virt_interface.Configure,
		primary_ip.Configure,
		virtual_machine.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
