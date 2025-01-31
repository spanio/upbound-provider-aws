/*
Copyright 2024 Upbound Inc.
*/

package opensearchserverless

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure adds configurations for the opensearchserverless group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_opensearchserverless_security_config", func(r *config.Resource) {
		r.SchemaElementOptions.SetEmbeddedObject("saml_options")
	})
	p.AddResourceConfigurator("aws_opensearchserverless_security_policy", func(r *config.Resource) {
		r.TerraformConfigurationInjector = config.CanonicalizeJSONParameters("policy")
	})
	p.AddResourceConfigurator("aws_opensearchserverless_lifecycle_policy", func(r *config.Resource) {
		r.TerraformConfigurationInjector = config.CanonicalizeJSONParameters("policy")
	})
	p.AddResourceConfigurator("aws_opensearchserverless_access_policy", func(r *config.Resource) {
		r.TerraformConfigurationInjector = config.CanonicalizeJSONParameters("policy")
	})
}
