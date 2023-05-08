![image](https://hub.steampipe.io/images/plugins/turbot/consul-social-graphic.png)

# Consul Plugin for Steampipe

Use SQL to query nodes, acls, services and more from Consul.

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

You may specify the Address to authenticate:

- `address`: The address of the consul server.

```hcl
connection "consul" {
  plugin  = "consul"
  address = "http://52.14.112.248:8500"
}
```

or you may specify the Address and Token to authenticate:

- `address`: The address of the consul server.
- `token`: The ACL token.

```hcl
connection "consul" {
  plugin  = "consul"
  address = "http://52.14.112.248:8500"
  token   = "c178b810-8b18-6f38-016f-725ddec5d58"
}
```

or if you are using consul enterprise then you may specify the Address, Token, namespace and partition to authenticate:

- `address`: The address of the consul server.
- `token`: The ACL token.
- `namespace`: The consul namespace.
- `partition`: The consul partition.

```hcl
connection "consul" {
  plugin    = "consul"
  address   = "http://52.14.112.248:8500"
  token     = "c178b810-8b18-6f38-016f-725ddec5d58"
  namespace = '*'
  partition = 'default'
}
```

or through environment variables

```sh
export CONSUL_HTTP_ADDR="http://18.118.144.168:4646"
export CONSUL_NAMESPACE="*"
export CONSUL_HTTP_TOKEN="c178b810-8b18-6f38-016f-725ddec5d58"
export CONSUL_PARTITION="default"
```

Run steampipe:

```shell
steampipe query
```

List your Consul services:

```sql
select
  service_id,
  service_name
  node,
  address,
  datacenter,
  namespace
from
  consul_service;
```

```
+------------+--------+---------------+---------------------------------+-----------+
| service_id | node   | address       | datacenter                      | namespace |
+------------+--------+---------------+---------------------------------+-----------+
| consul     | consul | 172.25.34.191 | consul-quickstart-1683117303883 | default   |
+------------+--------+---------------+---------------------------------+-----------+
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
