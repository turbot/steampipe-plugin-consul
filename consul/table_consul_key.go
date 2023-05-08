package consul

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableConsulKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "consul_key",
		Description: "Retrieve information about your keys.",
		List: &plugin.ListConfig{
			Hydrate: listKeys,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "namespace",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("key"),
			Hydrate:    getKey,
		},
		Columns: []*plugin.Column{
			{
				Name:        "key",
				Type:        proto.ColumnType_STRING,
				Description: "Name is the name of the Key. It must be unique and must be a DNS hostname. There are also other reserved names that may not be used.",
			},
			{
				Name:        "session",
				Type:        proto.ColumnType_STRING,
				Description: "Session is a string representing the ID of the session. Any other interactions with this key over the same session must specify the same session ID.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "CreateIndex holds the index corresponding the creation of this KVPair. This is a read-only field.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "ModifyIndex is used for the Check-And-Set operations and can also be fed back into the WaitIndex of the QueryOptions in order to perform blocking queries.",
			},
			{
				Name:        "lock_index",
				Type:        proto.ColumnType_INT,
				Description: "LockIndex holds the index corresponding to a lock on this key, if any. This is a read-only field.",
			},
			{
				Name:        "partition",
				Type:        proto.ColumnType_STRING,
				Description: "Partition is the partition the KVPair is associated with.",
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "Namespace is the namespace the KVPair is associated with.",
			},
			{
				Name:        "flags",
				Type:        proto.ColumnType_INT,
				Description: "Flags are any user-defined flags on the key. It is up to the implementer to check these values, since Consul does not treat them specially.",
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_JSON,
				Description: "Value is the value for the key. This can be any value, but it will be base64 encoded upon transport.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the key.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Key"),
			},
		},
	}
}

func listKeys(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("consul_key.listKeys", "connection_error", err)
		return nil, err
	}

	input := &api.QueryOptions{}
	if d.EqualsQuals["namespace"] != nil {
		input.Namespace = d.EqualsQualString("namespace")
	}

	keys, _, err := client.KV().List("", input)
	if err != nil {
		plugin.Logger(ctx).Error("consul_key.listKeys", "api_error", err)
		return nil, err
	}

	for _, key := range keys {
		d.StreamListItem(ctx, key)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	keyName := d.EqualsQualString("key")

	// check if keyName is empty
	if keyName == "" {
		return nil, nil
	}

	// Create client
	client, err := getClient(ctx, d)
	if err != nil {
		logger.Error("consul_key.getKey", "connection_error", err)
		return nil, err
	}

	key, _, err := client.KV().Get(keyName, &api.QueryOptions{})
	if err != nil {
		logger.Error("consul_key.getKey", "api_error", err)
		return nil, err
	}

	return key, nil
}
