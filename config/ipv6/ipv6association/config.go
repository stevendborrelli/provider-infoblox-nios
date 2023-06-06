package ipv6association

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infoblox_ipv6_association", func(r *config.Resource) {

		r.ShortGroup = "ipv6"
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.OmittedFields = []string{
			"name",
		}
	})
}
