/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

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
		region.Configure,
		site.Configure,
		site_group.Configure,
		custom_field.Configure,
		tag.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
