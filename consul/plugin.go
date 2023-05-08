package consul

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-consul",
		DefaultTransform: transform.FromCamel(),
		DefaultIgnoreConfig: &plugin.IgnoreConfig{
			ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"404"}),
		},
		DefaultRetryConfig: &plugin.RetryConfig{
			ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"consul_acl_auth_method":  tableConsulACLAuthMethod(ctx),
			"consul_acl_binding_rule": tableConsulACLBindingRule(ctx),
			"consul_acl_policy":       tableConsulACLPolicy(ctx),
			"consul_acl_role":         tableConsulACLRole(ctx),
			"consul_acl_token":        tableConsulACLToken(ctx),
			"consul_intention":        tableConsulIntention(ctx),
			"consul_key":              tableConsulKey(ctx),
			"consul_namespace":        tableConsulNamespace(ctx),
			"consul_node":             tableConsulNode(ctx),
			"consul_service":          tableConsulService(ctx),
		},
	}
	return p
}
