/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.

var ExternalNameConfigs = map[string]config.ExternalName{
	"netbox_device":           config.NameAsIdentifier,
	"netbox_device_role":      config.NameAsIdentifier,
	"netbox_device_type":      config.NameAsIdentifier,
	"netbox_device_interface": config.NameAsIdentifier,
	"netbox_location":         config.NameAsIdentifier,
	"netbox_manufacturer":     config.NameAsIdentifier,
	"netbox_platform":         config.NameAsIdentifier,
	"netbox_rack":             config.NameAsIdentifier,
	"netbox_rack_reservation": config.NameAsIdentifier,
	"netbox_rack_role":        config.NameAsIdentifier,
	"netbox_region":           config.NameAsIdentifier,
	"netbox_site":             config.NameAsIdentifier,
	"netbox_site_group":       config.NameAsIdentifier,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
