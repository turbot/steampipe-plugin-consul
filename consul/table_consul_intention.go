package consul

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableConsulIntention(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "consul_intention",
		Description: "Retrieve information about your intentions.",
		List: &plugin.ListConfig{
			Hydrate: listIntentions,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "source_name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "ID is the UUID-based ID for the intention, always generated by Consul.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "Description is a human-friendly description of this intention.",
			},
			{
				Name:        "source_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the source service.",
			},
			{
				Name:        "destination_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the destination service.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "CreateIndex holds the index corresponding the creation of this intention.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "ModifyIndex is the latest Raft index at which the intention was modified.",
			},
			{
				Name:        "source_partition",
				Type:        proto.ColumnType_STRING,
				Description: "The source partition of the intention.",
			},
			{
				Name:        "destination_partition",
				Type:        proto.ColumnType_STRING,
				Description: "The destination partition of the intention.",
			},
			{
				Name:        "source_ns",
				Type:        proto.ColumnType_STRING,
				Description: "The source namespace of the intention.",
				Transform:   transform.FromField("SourceNS"),
			},
			{
				Name:        "destination_ns",
				Type:        proto.ColumnType_STRING,
				Description: "The destination namespace of the intention.",
				Transform:   transform.FromField("DestinationNS"),
			},
			{
				Name:        "source_peer",
				Type:        proto.ColumnType_STRING,
				Description: "The source peer of the intention.",
			},
			{
				Name:        "source_type",
				Type:        proto.ColumnType_STRING,
				Description: "The source type of the intention.",
			},
			{
				Name:        "action",
				Type:        proto.ColumnType_STRING,
				Description: "Action is whether this is an allowlist or denylist intention.",
			},
			{
				Name:        "precedence",
				Type:        proto.ColumnType_INT,
				Description: "Precedence is the order that the intention will be applied, with larger numbers being applied first. This is a read-only field, on any intention update it is updated.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The create timestamp of the intention.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The update timestamp of the intention.",
			},
			{
				Name:        "permissions",
				Type:        proto.ColumnType_JSON,
				Description: "Permissions is the list of additional L7 attributes that extend the intention definition.",
			},
			{
				Name:        "meta",
				Type:        proto.ColumnType_STRING,
				Description: "Meta is arbitrary metadata associated with the intention.",
			},
			{
				Name:        "hash",
				Type:        proto.ColumnType_STRING,
				Description: "Hash of the contents of the intention.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the intention.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
		}),
	}
}

func listIntentions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("consul_intention.listIntentions", "connection_error", err)
		return nil, err
	}

	input := &api.QueryOptions{}
	if d.EqualsQualString("source_name") != "" {
		filter := fmt.Sprintf("SourceName== %q\n", d.EqualsQualString("source_name"))
		input.Filter = filter
	}

	intentions, _, err := client.Connect().Intentions(input)
	if err != nil {
		plugin.Logger(ctx).Error("consul_intention.listIntentions", "api_error", err)
		return nil, err
	}

	for _, intention := range intentions {
		d.StreamListItem(ctx, intention)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
