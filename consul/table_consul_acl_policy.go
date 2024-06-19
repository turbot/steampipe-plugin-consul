package consul

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableConsulACLPolicy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "consul_acl_policy",
		Description: "Retrieve information about your ACL policies.",
		List: &plugin.ListConfig{
			Hydrate: listACLPolicies,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "namespace",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getACLPolicy,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The id of the acl policy.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the acl policy.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the acl policy.",
			},
			{
				Name:        "rules",
				Type:        proto.ColumnType_STRING,
				Description: "The set of rules of the acl policy.",
				Hydrate:     getACLPolicy,
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "Namespace is the namespace the ACL policy is associated with.",
			},
			{
				Name:        "partition",
				Type:        proto.ColumnType_STRING,
				Description: "Partition is the partition the ACL policy is associated with.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the acl policy was created.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "The index when the acl policy was last modified.",
			},
			{
				Name:        "datacenters",
				Type:        proto.ColumnType_JSON,
				Description: "The datacenters of the acl policy.",
			},
			{
				Name:        "hash",
				Type:        proto.ColumnType_JSON,
				Description: "The hash of the acl policy.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the acl policy.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

func listACLPolicies(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("consul_acl_policy.listACLPolicies", "connection_error", err)
		return nil, err
	}

	input := &api.QueryOptions{}
	if d.EqualsQuals["namespace"] != nil {
		input.Namespace = d.EqualsQualString("namespace")
	}

	policies, _, err := client.ACL().PolicyList(input)
	if err != nil {
		plugin.Logger(ctx).Error("consul_acl_policy.listACLPolicies", "api_error", err)
		return nil, err
	}

	for _, policy := range policies {
		d.StreamListItem(ctx, policy)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getACLPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	var id string
	if h.Item != nil {
		id = h.Item.(*api.ACLPolicyListEntry).ID
	} else {
		id = d.EqualsQualString("id")
	}

	// check if id is empty
	if id == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("consul_acl_policy.getACLPolicy", "connection_error", err)
		return nil, err
	}

	policy, _, err := client.ACL().PolicyRead(id, &api.QueryOptions{})
	if err != nil {
		logger.Error("consul_acl_policy.getACLPolicy", "api_error", err)
		return nil, err
	}

	return policy, nil
}
