package acl_token

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("acl_token", func(r *config.Resource) {
		r.ShortGroup = "acl_token"
		r.References["policy"] = config.Reference{
			Type: "github.com/daanvinken/provider-consul/apis/policy/v1alpha1/v1alpha1.policy",
		}
	})
}
