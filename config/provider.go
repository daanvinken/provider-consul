/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/daanvinken/provider-consul/config/acl_policy"
	"github.com/daanvinken/provider-consul/config/acl_role"
	"github.com/daanvinken/provider-consul/config/acl_token"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "consul"
	modulePath     = "github.com/daanvinken/provider-consul"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("daanvinken.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		acl_role.Configure,
		acl_token.Configure,
		acl_policy.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
