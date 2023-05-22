/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	device "github.com/fire-ant/provider-netbox/internal/controller/device/device"
	manufacturer "github.com/fire-ant/provider-netbox/internal/controller/manufacturer/manufacturer"
	providerconfig "github.com/fire-ant/provider-netbox/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		device.Setup,
		manufacturer.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
