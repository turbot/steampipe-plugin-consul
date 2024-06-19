package consul

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableConsulNode(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "consul_node",
		Description: "Retrieve information about your nodes.",
		List: &plugin.ListConfig{
			Hydrate: listNodes,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "address",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("node"),
			Hydrate:    getNode,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "A string representing the unique identifier of the node.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "node",
				Type:        proto.ColumnType_STRING,
				Description: "A string representing the name of the node.",
			},
			{
				Name:        "address",
				Type:        proto.ColumnType_STRING,
				Description: "The node address.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "An unsigned 64-bit integer representing the index at which the node was created.",
			},
			{
				Name:        "datacenter",
				Type:        proto.ColumnType_STRING,
				Description: "A string specifying the datacenter in which the node is located.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "An unsigned 64-bit integer representing the index at which the node was last modified.",
			},
			{
				Name:        "partition",
				Type:        proto.ColumnType_STRING,
				Description: "The partition of the node.",
			},
			{
				Name:        "peer_name",
				Type:        proto.ColumnType_STRING,
				Description: "The node peer name.",
			},
			{
				Name:        "tagged_addresses",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing links to related tagged addresses associated with the node.",
			},
			{
				Name:        "meta",
				Type:        proto.ColumnType_JSON,
				Description: "A map containing additional metadata associated with the node.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the node.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Node"),
			},
		}),
	}
}

func listNodes(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("consul_node.listNodes", "connection_error", err)
		return nil, err
	}

	input := &api.QueryOptions{}
	if d.EqualsQualString("address") != "" {
		filter := fmt.Sprintf("Address== %q\n", d.EqualsQualString("address"))
		input.Filter = filter
	}

	nodes, _, err := client.Catalog().Nodes(input)
	if err != nil {
		plugin.Logger(ctx).Error("consul_node.listNodes", "api_error", err)
		return nil, err
	}

	for _, node := range nodes {
		d.StreamListItem(ctx, node)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getNode(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	node := d.EqualsQualString("node")

	// check if node is empty
	if node == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("consul_node.getNode", "connection_error", err)
		return nil, err
	}

	catalogNode, _, err := client.Catalog().Node(node, &api.QueryOptions{})
	if err != nil {
		logger.Error("consul_node.getNode", "api_error", err)
		return nil, err
	}

	return catalogNode.Node, nil
}
