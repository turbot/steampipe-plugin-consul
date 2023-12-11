![image](https://hub.steampipe.io/images/plugins/turbot/consul-social-graphic.png)

# Consul Plugin for Steampipe

Use SQL to query nodes, acls, services and more from Consul.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/consul)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/consul/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-consul/issues)

## Quick start

### Install

Download and install the latest Consul plugin:

```bash
steampipe plugin install consul
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/consul#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/consul#configuration).

Add your configuration details in `~/.steampipe/config/consul.spc`:

```hcl
connection "consul" {
  plugin    = "consul"
  # Authentication information
  address   = "http://52.14.112.248:8500"
  token     = "c178b810-8b18-6f38-016f-725ddec5d58"
}
```

- `token` parameter is only required to query the ACL tables like `consul_acl_auth_method`, `consul_acl_binding_rule`, `consul_acl_policy`, `consul_acl_role` and `consul_acl_token` tables.
- `namespace`, and `partition` parameters are only required to query the `consul_namespace` table.

Or through environment variables:

```sh
export CONSUL_HTTP_ADDR=http://18.118.144.168:4646
export CONSUL_HTTP_TOKEN=c178b810-8b18-6f38-016f-725ddec5d58
```

Run steampipe:

```shell
steampipe query
```

List your Consul services:

```sql
select
  service_id,
  service_name,
  node,
  address,
  datacenter,
  namespace
from
  consul_service;
```

```
+------------------------------------------------+--------------+------------------+---------------+------------+-----------+
| service_id                                     | service_name | node             | address       | datacenter | namespace |
+------------------------------------------------+--------------+------------------+---------------+------------+-----------+
| consul                                         | consul       | ip-172-31-30-170 | 172.31.30.170 | dc1        | default   |
+------------------------------------------------+--------------+------------------+---------------+------------+-----------+
```

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/index) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs//steampipe_sqlite/index) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/index) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-consul.git
cd steampipe-plugin-consul
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/consul.spc
```

Try it!

```
steampipe query
> .inspect consul
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-consul/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-consul/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Consul Plugin](https://github.com/turbot/steampipe-plugin-consul/labels/help%20wanted)
