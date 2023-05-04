package consul

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableConsulACLBindingRule(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "consul_acl_binding_rule",
		Description: "Retrieve information about your ACL binding rules.",
		List: &plugin.ListConfig{
			ParentHydrate: listACLAuthMethods,
			Hydrate:       listACLBindingRules,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getACLBindingRule,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "An internally generated UUID for this rule and is controlled by Consul.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the acl binding rule.",
			},
			{
				Name:        "auth_method",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the auth method for which this rule applies to.",
			},
			{
				Name:        "selector",
				Type:        proto.ColumnType_STRING,
				Description: "An expression that matches against verified identity attributes returned from the auth method during login.",
			},
			{
				Name:        "bind_type",
				Type:        proto.ColumnType_STRING,
				Description: "The binding type of the ACL binding rule.",
			},
			{
				Name:        "bind_name",
				Type:        proto.ColumnType_STRING,
				Description: "The binding name of the ACL binding rule.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "Create index of the ACL binding rule.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "Modify index of the ACL binding rule.",
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "Namespace is the namespace the ACL binding rule is associated with.",
			},
			{
				Name:        "partition",
				Type:        proto.ColumnType_STRING,
				Description: "Partition is the partition the ACL binding rule is associated with.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the acl binding rule.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
		},
	}
}

func listACLBindingRules(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	method := h.Item.(*api.ACLAuthMethodListEntry)

	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("consul_acl_binding_rule.listACLBindingRules", "connection_error", err)
		return nil, err
	}

	input := &api.QueryOptions{
		//PerPage: int32(maxLimit),
	}

	bindingRules, _, err := client.ACL().BindingRuleList(method.Name, input)
	if err != nil {
		plugin.Logger(ctx).Error("consul_acl_binding_rule.listACLBindingRules", "api_error", err)
		return nil, err
	}

	for _, bindingRule := range bindingRules {
		d.StreamListItem(ctx, bindingRule)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getACLBindingRule(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("consul_acl_binding_rule.getACLBindingRule", "connection_error", err)
		return nil, err
	}

	bindingRule, _, err := client.ACL().BindingRuleRead(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("consul_acl_binding_rule.getACLBindingRule", "api_error", err)
		return nil, err
	}

	return bindingRule, nil
}
