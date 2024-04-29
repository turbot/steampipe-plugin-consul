package consul

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableConsulACLToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "consul_acl_token",
		Description: "Retrieve information about your ACL tokens.",
		List: &plugin.ListConfig{
			Hydrate: listACLTokens,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "namespace",
					Require: plugin.Optional,
				},
				{
					Name:    "create_index",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("accessor_id"),
			Hydrate:    getACLToken,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "accessor_id",
				Type:        proto.ColumnType_STRING,
				Description: "The accessor ID of the acl token.",
				Transform:   transform.FromField("AccessorID"),
			},
			{
				Name:        "secret_id",
				Type:        proto.ColumnType_STRING,
				Description: "The secret ID of the acl token.",
				Transform:   transform.FromField("SecretID"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the acl token.",
			},
			{
				Name:        "local",
				Type:        proto.ColumnType_BOOL,
				Description: "Check whether the token is local or not.",
			},
			{
				Name:        "auth_method",
				Type:        proto.ColumnType_STRING,
				Description: "The auth method of the acl token.",
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "The namespace of the acl token.",
			},
			{
				Name:        "partition",
				Type:        proto.ColumnType_STRING,
				Description: "The partition of the acl token.",
			},
			{
				Name:        "auth_method_namespace",
				Type:        proto.ColumnType_STRING,
				Description: "The auth method namespace of the acl token.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The creation time of the acl token.",
			},
			{
				Name:        "expiration_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The expiration time of the acl token.",
			},
			{
				Name:        "expiration_ttl",
				Type:        proto.ColumnType_STRING,
				Description: "The maximum life of the acl token.",
				Transform:   transform.FromField("ExpirationTTL"),
				Hydrate:     getACLToken,
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "Create index of the acl token.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "Modify index of the acl token.",
			},
			{
				Name:        "policies",
				Type:        proto.ColumnType_JSON,
				Description: "Policies attached to the acl token.",
			},
			{
				Name:        "roles",
				Type:        proto.ColumnType_JSON,
				Description: "Roles attached to the acl token.",
			},
			{
				Name:        "service_identities",
				Type:        proto.ColumnType_JSON,
				Description: "Service identities attached to the acl token.",
			},
			{
				Name:        "node_identities",
				Type:        proto.ColumnType_JSON,
				Description: "Node identities attached to the acl token.",
			},
			{
				Name:        "hash",
				Type:        proto.ColumnType_JSON,
				Description: "The acl token hash.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the acl token.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("AccessorID"),
			},
		}),
	}
}

func listACLTokens(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("consul_acl_token.listACLTokens", "connection_error", err)
		return nil, err
	}

	input := &api.QueryOptions{}
	if d.EqualsQuals["namespace"] != nil {
		input.Namespace = d.EqualsQualString("namespace")
	}
	if d.EqualsQuals["create_index"] != nil {
		filter := fmt.Sprintf("CreateIndex== %q\n", d.EqualsQuals["create_index"].GetStringValue())
		input.Filter = filter
	}

	tokens, _, err := client.ACL().TokenList(input)
	if err != nil {
		plugin.Logger(ctx).Error("consul_acl_token.listACLTokens", "api_error", err)
		return nil, err
	}

	for _, token := range tokens {
		d.StreamListItem(ctx, token)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getACLToken(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var accessorID string
	if h.Item != nil {
		accessorID = h.Item.(*api.ACLTokenListEntry).AccessorID
	} else {
		accessorID = d.EqualsQualString("accessor_id")
	}

	// check if accessorID is empty
	if accessorID == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("consul_acl_token.getACLToken", "connection_error", err)
		return nil, err
	}

	token, _, err := client.ACL().TokenRead(accessorID, &api.QueryOptions{})
	if err != nil {
		logger.Error("consul_acl_token.getACLToken", "api_error", err)
		return nil, err
	}

	return token, nil
}
