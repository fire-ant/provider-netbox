/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	ipaddress "github.com/fire-ant/provider-netbox/internal/controller/available/ipaddress"
	prefix "github.com/fire-ant/provider-netbox/internal/controller/available/prefix"
	circuittype "github.com/fire-ant/provider-netbox/internal/controller/circuit/circuittype"
	provider "github.com/fire-ant/provider-netbox/internal/controller/circuit/provider"
	termination "github.com/fire-ant/provider-netbox/internal/controller/circuit/termination"
	clustertype "github.com/fire-ant/provider-netbox/internal/controller/cluster/clustertype"
	group "github.com/fire-ant/provider-netbox/internal/controller/cluster/group"
	assignment "github.com/fire-ant/provider-netbox/internal/controller/contact/assignment"
	groupcontact "github.com/fire-ant/provider-netbox/internal/controller/contact/group"
	field "github.com/fire-ant/provider-netbox/internal/controller/custom/field"
	deviceinterface "github.com/fire-ant/provider-netbox/internal/controller/device/deviceinterface"
	devicetype "github.com/fire-ant/provider-netbox/internal/controller/device/devicetype"
	role "github.com/fire-ant/provider-netbox/internal/controller/device/role"
	address "github.com/fire-ant/provider-netbox/internal/controller/ip/address"
	iprange "github.com/fire-ant/provider-netbox/internal/controller/ip/iprange"
	roleipam "github.com/fire-ant/provider-netbox/internal/controller/ipam/role"
	aggregate "github.com/fire-ant/provider-netbox/internal/controller/netbox/aggregate"
	asn "github.com/fire-ant/provider-netbox/internal/controller/netbox/asn"
	circuit "github.com/fire-ant/provider-netbox/internal/controller/netbox/circuit"
	cluster "github.com/fire-ant/provider-netbox/internal/controller/netbox/cluster"
	contact "github.com/fire-ant/provider-netbox/internal/controller/netbox/contact"
	device "github.com/fire-ant/provider-netbox/internal/controller/netbox/device"
	location "github.com/fire-ant/provider-netbox/internal/controller/netbox/location"
	manufacturer "github.com/fire-ant/provider-netbox/internal/controller/netbox/manufacturer"
	platform "github.com/fire-ant/provider-netbox/internal/controller/netbox/platform"
	prefixnetbox "github.com/fire-ant/provider-netbox/internal/controller/netbox/prefix"
	rack "github.com/fire-ant/provider-netbox/internal/controller/netbox/rack"
	region "github.com/fire-ant/provider-netbox/internal/controller/netbox/region"
	rir "github.com/fire-ant/provider-netbox/internal/controller/netbox/rir"
	service "github.com/fire-ant/provider-netbox/internal/controller/netbox/service"
	site "github.com/fire-ant/provider-netbox/internal/controller/netbox/site"
	tag "github.com/fire-ant/provider-netbox/internal/controller/netbox/tag"
	tenant "github.com/fire-ant/provider-netbox/internal/controller/netbox/tenant"
	virtinterface "github.com/fire-ant/provider-netbox/internal/controller/netbox/virtinterface"
	vlan "github.com/fire-ant/provider-netbox/internal/controller/netbox/vlan"
	vrf "github.com/fire-ant/provider-netbox/internal/controller/netbox/vrf"
	ip "github.com/fire-ant/provider-netbox/internal/controller/primary/ip"
	providerconfig "github.com/fire-ant/provider-netbox/internal/controller/providerconfig"
	reservation "github.com/fire-ant/provider-netbox/internal/controller/rack/reservation"
	rolerack "github.com/fire-ant/provider-netbox/internal/controller/rack/role"
	target "github.com/fire-ant/provider-netbox/internal/controller/route/target"
	grouptenant "github.com/fire-ant/provider-netbox/internal/controller/tenant/group"
	machine "github.com/fire-ant/provider-netbox/internal/controller/virtual/machine"
	groupvlan "github.com/fire-ant/provider-netbox/internal/controller/vlan/group"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		ipaddress.Setup,
		prefix.Setup,
		circuittype.Setup,
		provider.Setup,
		termination.Setup,
		clustertype.Setup,
		group.Setup,
		assignment.Setup,
		groupcontact.Setup,
		field.Setup,
		deviceinterface.Setup,
		devicetype.Setup,
		role.Setup,
		address.Setup,
		iprange.Setup,
		roleipam.Setup,
		aggregate.Setup,
		asn.Setup,
		circuit.Setup,
		cluster.Setup,
		contact.Setup,
		device.Setup,
		location.Setup,
		manufacturer.Setup,
		platform.Setup,
		prefixnetbox.Setup,
		rack.Setup,
		region.Setup,
		rir.Setup,
		service.Setup,
		site.Setup,
		tag.Setup,
		tenant.Setup,
		virtinterface.Setup,
		vlan.Setup,
		vrf.Setup,
		ip.Setup,
		providerconfig.Setup,
		reservation.Setup,
		rolerack.Setup,
		target.Setup,
		grouptenant.Setup,
		machine.Setup,
		groupvlan.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
