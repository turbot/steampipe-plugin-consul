package consul

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableConsulACLRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "consul_acl_role",
		Description: "Retrieve information about your ACL roles.",
		List: &plugin.ListConfig{
			Hydrate: listACLRoles,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getACLRole,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the ACL role.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the ACL role.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "A human-readable, operator set description that can provide additional context about the ACL role.",
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "Namespace is the namespace the ACL role is associated with.",
			},
			{
				Name:        "partition",
				Type:        proto.ColumnType_STRING,
				Description: "Partition is the partition the ACL role is associated with.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the ACL role was created.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the ACL role was last modified.",
			},
			{
				Name:        "policies",
				Type:        proto.ColumnType_JSON,
				Description: "An array of ACL policy links.",
			},
			{
				Name:        "service_identities",
				Type:        proto.ColumnType_JSON,
				Description: "Service identities attached to the acl role.",
			},
			{
				Name:        "node_identities",
				Type:        proto.ColumnType_JSON,
				Description: "Node identities attached to the acl role.",
			},
			{
				Name:        "hash",
				Type:        proto.ColumnType_JSON,
				Description: "The hash of the acl role.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the acl role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listACLRoles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("consul_acl_role.listACLRoles", "connection_error", err)
		return nil, err
	}

	input := &api.QueryOptions{
		//PerPage: int32(maxLimit),
	}

	roles, _, err := client.ACL().RoleList(input)
	if err != nil {
		plugin.Logger(ctx).Error("consul_acl_role.listACLRoles", "api_error", err)
		return nil, err
	}

	for _, role := range roles {
		d.StreamListItem(ctx, role)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getACLRole(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	id := d.EqualsQualString("id")

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("consul_acl_role.getACLRole", "connection_error", err)
		return nil, err
	}

	role, _, err := client.ACL().RoleRead(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("consul_acl_role.getACLRole", "api_error", err)
		return nil, err
	}

	return role, nil
}
