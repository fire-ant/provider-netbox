/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"netbox_device":               config.NameAsIdentifier,
	"netbox_device_role":          config.NameAsIdentifier,
	"netbox_device_type":          config.NameAsIdentifier,
	"netbox_device_interface":     config.NameAsIdentifier,
	"netbox_location":             config.NameAsIdentifier,
	"netbox_manufacturer":         config.NameAsIdentifier,
	"netbox_platform":             config.NameAsIdentifier,
	"netbox_rack":                 config.NameAsIdentifier,
	"netbox_rack_reservation":     config.NameAsIdentifier,
	"netbox_rack_role":            config.NameAsIdentifier,
	"netbox_region":               config.NameAsIdentifier,
	"netbox_site":                 config.NameAsIdentifier,
	"netbox_custom_field":         config.NameAsIdentifier,
	"netbox_tag":                  config.NameAsIdentifier,
	"netbox_aggregate":            config.NameAsIdentifier,
	"netbox_asn":                  config.NameAsIdentifier,
	"netbox_available_ip_address": config.NameAsIdentifier,
	"netbox_available_prefix":     config.NameAsIdentifier,
	"netbox_ip_address":           config.NameAsIdentifier,
	"netbox_ip_range":             config.NameAsIdentifier,
	"netbox_ipam_role":            config.NameAsIdentifier,
	"netbox_prefix":               config.NameAsIdentifier,
	"netbox_rir":                  config.NameAsIdentifier,
	"netbox_route_target":         config.NameAsIdentifier,
	"netbox_service":              config.NameAsIdentifier,
	"netbox_vlan":                 config.NameAsIdentifier,
	"netbox_vlan_group":           config.NameAsIdentifier,
	"netbox_vrf":                  config.NameAsIdentifier,
	"netbox_contact":              config.NameAsIdentifier,
	"netbox_contact_assignment":   config.NameAsIdentifier,
	"netbox_contact_group":        config.NameAsIdentifier,
	"netbox_tenant":               config.NameAsIdentifier,
	"netbox_tenant_group":         config.NameAsIdentifier,
	"netbox_tenants":              config.NameAsIdentifier,
	"netbox_cluster":              config.NameAsIdentifier,
	"netbox_cluster_group":        config.NameAsIdentifier,
	"netbox_cluster_type":         config.NameAsIdentifier,
	"netbox_interface":            config.NameAsIdentifier,
	"netbox_primary_ip":           config.NameAsIdentifier,
	"netbox_virtual_machine":      config.NameAsIdentifier,
	"netbox_circuit":              config.NameAsIdentifier,
	"netbox_circuit_provider":     config.NameAsIdentifier,
	"netbox_circuit_termination":  config.NameAsIdentifier,
	"netbox_circuit_type":         config.NameAsIdentifier,
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
