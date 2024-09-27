package repository

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("acl_role", func(r *config.Resource) {
		r.ShortGroup = "acl_role"
	})
}
