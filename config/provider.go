/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/fire-ant/provider-netbox/config/device"
	"github.com/fire-ant/provider-netbox/config/manufacturer"
	"github.com/fire-ant/provider-netbox/config/role"
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
		manufacturer.Configure,
		role.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
