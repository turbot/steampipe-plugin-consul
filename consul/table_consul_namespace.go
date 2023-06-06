package consul

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableConsulNamespace(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "consul_namespace",
		Description: "Retrieve information about your namespaces.",
		List: &plugin.ListConfig{
			Hydrate: listNamespaces,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "create_index",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getNamespace,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name is the name of the Namespace. It must be unique and must be a DNS hostname. There are also other reserved names that may not be used.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description is where the user puts any information they want about the namespace. It is not used internally.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "CreateIndex is the Raft index at which the Namespace was created.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "ModifyIndex is the latest Raft index at which the Namespace was modified.",
			},
			{
				Name:        "partition",
				Type:        proto.ColumnType_STRING,
				Description: "Partition which contains the Namespace.",
			},
			{
				Name:        "deleted_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "DeletedAt is the time when the Namespace was marked for deletion.",
			},
			{
				Name:        "acls",
				Type:        proto.ColumnType_JSON,
				Description: "ACLs is the configuration of ACLs for this namespace.",
				Transform:   transform.FromField("ACLs"),
			},
			{
				Name:        "meta",
				Type:        proto.ColumnType_JSON,
				Description: "Meta is a map that can be used to add kv metadata to the namespace definition.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the namespace.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listNamespaces(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("consul_namespace.listNamespaces", "connection_error", err)
		return nil, err
	}

	input := &api.QueryOptions{}
	if d.EqualsQuals["create_index"] != nil {
		filter := fmt.Sprintf("CreateIndex== %q\n", d.EqualsQuals["create_index"].GetStringValue())
		input.Filter = filter
	}

	namespaces, _, err := client.Namespaces().List(input)
	if err != nil {
		plugin.Logger(ctx).Error("consul_namespace.listNamespaces", "api_error", err)
		return nil, err
	}

	for _, namespace := range namespaces {
		d.StreamListItem(ctx, namespace)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getNamespace(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	name := d.EqualsQualString("name")

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("consul_namespace.getNamespace", "connection_error", err)
		return nil, err
	}

	namespace, _, err := client.Namespaces().Read(name, &api.QueryOptions{})
	if err != nil {
		logger.Error("consul_namespace.getNamespace", "api_error", err)
		return nil, err
	}

	return namespace, nil
}
