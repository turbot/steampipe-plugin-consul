![image](https://hub.steampipe.io/images/plugins/turbot/consul-social-graphic.png)

# Consul Plugin for Steampipe

Use SQL to query nodes, jobs, deployments and more from Consul.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/consul)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/consul/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-consul/issues)

## Quick start

Download and install the latest Consul plugin:

```bash
steampipe plugin install consul
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/consul#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/consul#configuration).

### Configuring Consul Credentials

Configure your account details in `~/.steampipe/config/consul.spc`:

You may specify the Address and Namespace to authenticate:

- `address`: The address of the consul server.
- `namespace`: The Consul Cluster namespace.

```hcl
connection "consul" {
  plugin    = "consul"
  address   = "http://18.118.144.168:4646"
  namespace = "*"
}
```

or you may specify the Address, Namespace and SecretID to authenticate:

- `address`: The address of the consul server.
- `namespace`: The Consul Cluster namespace.
- `secret_id`: The SecretID of an ACL token.

```hcl
connection "consul" {
  plugin    = "consul"
  address   = "http://18.118.144.168:4646"
  namespace = "*"
  secret_id = "c178b810-8b18-6f38-016f-725ddec5d58"
}
```

or through environment variables

```sh
export NOMAD_ADDR="http://18.118.144.168:4646"
export NOMAD_NAMESPACE="*"
export NOMAD_TOKEN="c178b810-8b18-6f38-016f-725ddec5d58"
```

Run steampipe:

```shell
steampipe query
```

List your Consul jobs:

```sql
select
  id,
  name,
  status,
  dispatched,
  namespace,
  priority,
  region
from
  nomad_job;
```

```
+------+------+---------+------------+-----------+----------+--------+
| id   | name | status  | dispatched | namespace | priority | region |
+------+------+---------+------------+-----------+----------+--------+
| docs | docs | pending | false      | default   | 50       | global |
+------+------+---------+------------+-----------+----------+--------+
```

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

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-consul/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Consul Plugin](https://github.com/turbot/steampipe-plugin-consul/labels/help%20wanted)
