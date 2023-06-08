/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	circuit "github.com/fire-ant/provider-netbox/internal/controller/circuits/circuit"
	circuitprovider "github.com/fire-ant/provider-netbox/internal/controller/circuits/circuitprovider"
	circuittype "github.com/fire-ant/provider-netbox/internal/controller/circuits/circuittype"
	termination "github.com/fire-ant/provider-netbox/internal/controller/circuits/termination"
	device "github.com/fire-ant/provider-netbox/internal/controller/dcim/device"
	deviceinterface "github.com/fire-ant/provider-netbox/internal/controller/dcim/deviceinterface"
	devicerole "github.com/fire-ant/provider-netbox/internal/controller/dcim/devicerole"
	devicetype "github.com/fire-ant/provider-netbox/internal/controller/dcim/devicetype"
	group "github.com/fire-ant/provider-netbox/internal/controller/dcim/group"
	location "github.com/fire-ant/provider-netbox/internal/controller/dcim/location"
	manufacturer "github.com/fire-ant/provider-netbox/internal/controller/dcim/manufacturer"
	platform "github.com/fire-ant/provider-netbox/internal/controller/dcim/platform"
	rack "github.com/fire-ant/provider-netbox/internal/controller/dcim/rack"
	region "github.com/fire-ant/provider-netbox/internal/controller/dcim/region"
	reservation "github.com/fire-ant/provider-netbox/internal/controller/dcim/reservation"
	role "github.com/fire-ant/provider-netbox/internal/controller/dcim/role"
	site "github.com/fire-ant/provider-netbox/internal/controller/dcim/site"
	field "github.com/fire-ant/provider-netbox/internal/controller/extras/field"
	tag "github.com/fire-ant/provider-netbox/internal/controller/extras/tag"
	aggregate "github.com/fire-ant/provider-netbox/internal/controller/ipam/aggregate"
	asn "github.com/fire-ant/provider-netbox/internal/controller/ipam/asn"
	availableipadrress "github.com/fire-ant/provider-netbox/internal/controller/ipam/availableipadrress"
	availableprefix "github.com/fire-ant/provider-netbox/internal/controller/ipam/availableprefix"
	groupipam "github.com/fire-ant/provider-netbox/internal/controller/ipam/group"
	ipaddress "github.com/fire-ant/provider-netbox/internal/controller/ipam/ipaddress"
	ipamrole "github.com/fire-ant/provider-netbox/internal/controller/ipam/ipamrole"
	iprange "github.com/fire-ant/provider-netbox/internal/controller/ipam/iprange"
	prefix "github.com/fire-ant/provider-netbox/internal/controller/ipam/prefix"
	rir "github.com/fire-ant/provider-netbox/internal/controller/ipam/rir"
	service "github.com/fire-ant/provider-netbox/internal/controller/ipam/service"
	target "github.com/fire-ant/provider-netbox/internal/controller/ipam/target"
	vlan "github.com/fire-ant/provider-netbox/internal/controller/ipam/vlan"
	vrf "github.com/fire-ant/provider-netbox/internal/controller/ipam/vrf"
	providerconfig "github.com/fire-ant/provider-netbox/internal/controller/providerconfig"
	assignment "github.com/fire-ant/provider-netbox/internal/controller/tenant/assignment"
	contact "github.com/fire-ant/provider-netbox/internal/controller/tenant/contact"
	grouptenant "github.com/fire-ant/provider-netbox/internal/controller/tenant/group"
	tenant "github.com/fire-ant/provider-netbox/internal/controller/tenant/tenant"
	tenantgroup "github.com/fire-ant/provider-netbox/internal/controller/tenant/tenantgroup"
	cluster "github.com/fire-ant/provider-netbox/internal/controller/virtualization/cluster"
	clustertype "github.com/fire-ant/provider-netbox/internal/controller/virtualization/clustertype"
	groupvirtualization "github.com/fire-ant/provider-netbox/internal/controller/virtualization/group"
	ip "github.com/fire-ant/provider-netbox/internal/controller/virtualization/ip"
	machine "github.com/fire-ant/provider-netbox/internal/controller/virtualization/machine"
	virtinterface "github.com/fire-ant/provider-netbox/internal/controller/virtualization/virtinterface"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		circuit.Setup,
		circuitprovider.Setup,
		circuittype.Setup,
		termination.Setup,
		device.Setup,
		deviceinterface.Setup,
		devicerole.Setup,
		devicetype.Setup,
		group.Setup,
		location.Setup,
		manufacturer.Setup,
		platform.Setup,
		rack.Setup,
		region.Setup,
		reservation.Setup,
		role.Setup,
		site.Setup,
		field.Setup,
		tag.Setup,
		aggregate.Setup,
		asn.Setup,
		availableipadrress.Setup,
		availableprefix.Setup,
		groupipam.Setup,
		ipaddress.Setup,
		ipamrole.Setup,
		iprange.Setup,
		prefix.Setup,
		rir.Setup,
		service.Setup,
		target.Setup,
		vlan.Setup,
		vrf.Setup,
		providerconfig.Setup,
		assignment.Setup,
		contact.Setup,
		grouptenant.Setup,
		tenant.Setup,
		tenantgroup.Setup,
		cluster.Setup,
		clustertype.Setup,
		groupvirtualization.Setup,
		ip.Setup,
		machine.Setup,
		virtinterface.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
