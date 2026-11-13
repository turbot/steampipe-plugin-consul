package consul

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v6/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v6/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "address_url",
			Description: "The address URL.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getAddressUrl,
		},
	}, c...)
}

var getAddressUrlMemoize = plugin.HydrateFunc(getAddressUrlUncached).Memoize(memoize.WithCacheKeyFunction(getAddressUrlCacheKey))

func getAddressUrlCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheKey := "getAddressUrl"
	return cacheKey, nil
}

func getAddressUrl(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	config, err := getAddressUrlMemoize(ctx, d, h)
	if err != nil {
		return nil, err
	}

	c := config.(consulConfig)

	return c.Address, nil
}

func getAddressUrlUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cacheKey := "getAddressUrl"

	var consultData consulConfig

	if cachedData, ok := d.ConnectionCache.Get(ctx, cacheKey); ok {
		consultData = cachedData.(consulConfig)
	} else {
		consultData = GetConfig(d.Connection)

		_ = d.ConnectionCache.Set(ctx, cacheKey, consultData)
	}

	return consultData, nil
}
