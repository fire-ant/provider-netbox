/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/fire-ant/provider-netbox/config/device"
	"github.com/fire-ant/provider-netbox/config/device_interface"
	"github.com/fire-ant/provider-netbox/config/device_role"
	"github.com/fire-ant/provider-netbox/config/device_type"
	"github.com/fire-ant/provider-netbox/config/location"
	"github.com/fire-ant/provider-netbox/config/manufacturer"
	"github.com/fire-ant/provider-netbox/config/platform"
	"github.com/fire-ant/provider-netbox/config/rack"
	"github.com/fire-ant/provider-netbox/config/rack_group"
	"github.com/fire-ant/provider-netbox/config/rack_region"
	"github.com/fire-ant/provider-netbox/config/rack_reservation"
	"github.com/fire-ant/provider-netbox/config/rack_role"
	"github.com/fire-ant/provider-netbox/config/rack_site"
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
		rack_region.Configure,
		rack_site.Configure,
		rack_group.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
