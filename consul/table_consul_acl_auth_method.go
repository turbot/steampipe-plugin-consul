package consul

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableConsulACLAuthMethod(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "consul_acl_auth_method",
		Description: "Retrieve information about your ACL auth methods.",
		List: &plugin.ListConfig{
			Hydrate: listACLAuthMethods,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getACLAuthMethod,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the acl auth method.",
			},
			{
				Name:        "display_name",
				Type:        proto.ColumnType_STRING,
				Description: "The display name of the acl auth method.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the acl auth method.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Type is the SSO identifier of this auth method.",
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "Namespace is the namespace the ACL auth method is associated with.",
			},
			{
				Name:        "partition",
				Type:        proto.ColumnType_STRING,
				Description: "Partition is the partition the ACL auth method is associated with.",
			},
			{
				Name:        "token_locality",
				Type:        proto.ColumnType_STRING,
				Description: "Defines whether the auth method creates a local or global token when performing SSO login.",
			},
			{
				Name:        "max_token_ttl",
				Type:        proto.ColumnType_STRING,
				Description: "The maximum life of a token created by this method.",
				Transform:   transform.FromField("MaxTokenTTL"),
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "Create index of the auth method.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "Modify index of the auth method.",
			},
			{
				Name:        "config",
				Type:        proto.ColumnType_JSON,
				Description: "Config contains the detailed configuration which is specific to the auth method.",
				Hydrate:     getACLAuthMethod,
			},
			{
				Name:        "namespace_rules",
				Type:        proto.ColumnType_JSON,
				Description: "Namespace rules apply only on auth methods defined in the default namespace.",
				Hydrate:     getACLAuthMethod,
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the acl auth method.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listACLAuthMethods(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("consul_acl_auth_method.listACLAuthMethods", "connection_error", err)
		return nil, err
	}

	input := &api.QueryOptions{
		//	PerPage: int32(maxLimit),
	}

	authMethods, _, err := client.ACL().AuthMethodList(input)
	if err != nil {
		plugin.Logger(ctx).Error("consul_acl_auth_method.listACLAuthMethods", "api_error", err)
		return nil, err
	}

	for _, authMethod := range authMethods {
		d.StreamListItem(ctx, authMethod)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getACLAuthMethod(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var name string
	if h.Item != nil {
		name = h.Item.(*api.ACLAuthMethodListEntry).Name
	} else {
		name = d.EqualsQualString("name")
	}

	// check if name is empty
	if name == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("consul_acl_auth_method.getACLAuthMethod", "connection_error", err)
		return nil, err
	}

	authMethod, _, err := client.ACL().AuthMethodRead(name, &api.QueryOptions{})
	if err != nil {
		logger.Error("consul_acl_auth_method.getACLAuthMethod", "api_error", err)
		return nil, err
	}

	return authMethod, nil
}
