package acl_token

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("acl_token", func(r *config.Resource) {
		r.ShortGroup = "acl_token"
		r.References["policy"] = config.Reference{
			TerraformName: "acl_policy",
		}
	})
}
