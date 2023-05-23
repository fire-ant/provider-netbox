/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	deviceinterface "github.com/fire-ant/provider-netbox/internal/controller/device/deviceinterface"
	devicetype "github.com/fire-ant/provider-netbox/internal/controller/device/devicetype"
	role "github.com/fire-ant/provider-netbox/internal/controller/device/role"
	device "github.com/fire-ant/provider-netbox/internal/controller/netbox/device"
	location "github.com/fire-ant/provider-netbox/internal/controller/netbox/location"
	manufacturer "github.com/fire-ant/provider-netbox/internal/controller/netbox/manufacturer"
	platform "github.com/fire-ant/provider-netbox/internal/controller/netbox/platform"
	rack "github.com/fire-ant/provider-netbox/internal/controller/netbox/rack"
	providerconfig "github.com/fire-ant/provider-netbox/internal/controller/providerconfig"
	reservation "github.com/fire-ant/provider-netbox/internal/controller/rack/reservation"
	rolerack "github.com/fire-ant/provider-netbox/internal/controller/rack/role"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		deviceinterface.Setup,
		devicetype.Setup,
		role.Setup,
		device.Setup,
		location.Setup,
		manufacturer.Setup,
		platform.Setup,
		rack.Setup,
		providerconfig.Setup,
		reservation.Setup,
		rolerack.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
