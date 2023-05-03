package consul

import (
	"context"
	"errors"
	"os"

	"github.com/hashicorp/consul/api"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type consulConfig struct {
	Address   *string `cty:"address"`
	Namespace *string `cty:"namespace"`
	Token     *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"address": {
		Type: schema.TypeString,
	},
	"namespace": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &consulConfig{}
}

func GetConfig(connection *plugin.Connection) consulConfig {
	if connection == nil || connection.Config == nil {
		return consulConfig{}
	}
	config, _ := connection.Config.(consulConfig)
	return config
}

func getClient(ctx context.Context, d *plugin.QueryData) (*api.Client, error) {
	consulConfig := GetConfig(d.Connection)

	address := os.Getenv("CONSUL_HTTP_ADDR")
	namespace := os.Getenv("CONSUL_NAMESPACE")
	token := os.Getenv("CONSUL_HTTP_TOKEN")

	if consulConfig.Address != nil {
		address = *consulConfig.Address
	}
	if consulConfig.Token != nil {
		token = *consulConfig.Token
	}
	if consulConfig.Namespace != nil {
		namespace = *consulConfig.Namespace
	}

	if address != "" {
		con := api.DefaultConfig()
		con.Address = address
		con.Token = token
		con.Namespace = namespace
		client, _ := api.NewClient(con)
		return client, nil
	}

	return nil, errors.New("'address' or ('address' and 'token') must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
}
