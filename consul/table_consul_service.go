package consul

import (
	"context"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableConsulService(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "consul_service",
		Description: "Retrieve information about your services.",
		List: &plugin.ListConfig{
			Hydrate: listServices,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "namespace",
					Require: plugin.Optional,
				},
				{
					Name:    "node",
					Require: plugin.Optional,
				},
				{
					Name:    "service_id",
					Require: plugin.Optional,
				},
				{
					Name:    "service_name",
					Require: plugin.Optional,
				},
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "A unique identifier for the service.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "node",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the node where the service is running.",
			},
			{
				Name:        "address",
				Type:        proto.ColumnType_STRING,
				Description: "The IP address or hostname of the service.",
			},
			{
				Name:        "datacenter",
				Type:        proto.ColumnType_STRING,
				Description: "The datacenter where the service is running.",
			},
			{
				Name:        "service_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique identifier for the service instance.",
				Transform:   transform.FromField("ServiceID"),
			},
			{
				Name:        "service_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the service.",
			},
			{
				Name:        "service_tag",
				Type:        proto.ColumnType_STRING,
				Description: "The service tag.",
				Transform:   transform.FromField("ServiceTags").Transform(fetchServiceTag),
			},
			{
				Name:        "service_address",
				Type:        proto.ColumnType_STRING,
				Description: "The IP address or hostname of the service instance.",
			},
			{
				Name:        "service_port",
				Type:        proto.ColumnType_INT,
				Description: "The port on which the service instance is listening.",
			},
			{
				Name:        "service_enable_tag_override",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean indicating whether service tags can be overridden.",
			},
			{
				Name:        "create_index",
				Type:        proto.ColumnType_INT,
				Description: "The index of the create operation that created the service.",
			},
			{
				Name:        "modify_index",
				Type:        proto.ColumnType_INT,
				Description: "The index of the last modify operation that modified the service.",
			},
			{
				Name:        "namespace",
				Type:        proto.ColumnType_STRING,
				Description: "The namespace of the service.",
			},
			{
				Name:        "partition",
				Type:        proto.ColumnType_STRING,
				Description: "The partition of the service.",
			},
			{
				Name:        "tagged_addresses",
				Type:        proto.ColumnType_JSON,
				Description: "A map of service tags to IP addresses or hostnames.",
			},
			{
				Name:        "node_meta",
				Type:        proto.ColumnType_JSON,
				Description: "A map of metadata associated with the node where the service is running.",
			},
			{
				Name:        "service_tagged_addresses",
				Type:        proto.ColumnType_JSON,
				Description: "A map of service tags to service addresses.",
			},
			{
				Name:        "service_meta",
				Type:        proto.ColumnType_JSON,
				Description: "A map of metadata associated with the service.",
			},
			{
				Name:        "service_weights",
				Type:        proto.ColumnType_JSON,
				Description: "The weights associated with the service.",
			},
			{
				Name:        "service_proxy",
				Type:        proto.ColumnType_JSON,
				Description: "The proxy configuration for the service.",
			},
			{
				Name:        "checks",
				Type:        proto.ColumnType_JSON,
				Description: "The health checks associated with the service.",
			},

			/// Steampipe standard columns
			{
				Name:        "title",
				Description: "The title of the service.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ServiceName"),
			},
		},
	}
}

func listServices(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("consul_service.listServices", "connection_error", err)
		return nil, err
	}

	input := &api.QueryOptions{}
	if d.EqualsQuals["namespace"] != nil {
		input.Namespace = d.EqualsQualString("namespace")
	}
	if d.EqualsQuals["service_id"] != nil {
		filter := "ServiceID==" + d.EqualsQualString("service_id")
		input.Filter = filter
	}
	if d.EqualsQuals["service_name"] != nil {
		filter := "ServiceName==" + d.EqualsQualString("service_name")
		input.Filter = filter
	}
	if d.EqualsQuals["node"] != nil {
		filter := "Node==" + d.EqualsQualString("node")
		input.Filter = filter
	}

	serviceTags, _, err := client.Catalog().Services(input)
	if err != nil {
		plugin.Logger(ctx).Error("consul_service.listServices", "api_error", err)
		return nil, err
	}

	for service, tags := range serviceTags {
		if len(tags) > 0 {
			for _, tag := range tags {
				services, _, err := client.Catalog().Service(service, tag, &api.QueryOptions{})
				if err != nil {
					plugin.Logger(ctx).Error("consul_service.Service_with_tag", "api_error", err)
					return nil, err
				}
				if services != nil {
					d.StreamListItem(ctx, services[0])

					// Context can be cancelled due to manual cancellation or the limit has been hit
					if d.RowsRemaining(ctx) == 0 {
						return nil, nil
					}
				}
			}
		} else {
			services, _, err := client.Catalog().Service(service, "", &api.QueryOptions{})
			if err != nil {
				plugin.Logger(ctx).Error("consul_service.Service_without_tag", "api_error", err)
				return nil, err
			}
			if services != nil {
				d.StreamListItem(ctx, services[0])

				// Context can be cancelled due to manual cancellation or the limit has been hit
				if d.RowsRemaining(ctx) == 0 {
					return nil, nil
				}
			}
		}
	}

	return nil, nil
}

func fetchServiceTag(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	tags := d.Value.([]string)
	if len(tags) == 0 {
		return nil, nil
	}
	return tags[0], nil
}
