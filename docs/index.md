---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/consul.svg"
brand_color: "#E03875"
display_name: "Consul"
short_name: "consul"
description: "Steampipe plugin to query nodes, acls, services and more from Consul."
og_description: "Query Consul with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/consul-social-graphic.png"
---

# Consul + Steampipe

[Consul](https://www.consul.io/) is a service networking solution to automate network configurations, discover services, and enable secure connectivity across any cloud or runtime.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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

## Documentation

- **[Table definitions & examples →](/plugins/turbot/consul/tables)**

## Quick start

### Install

Download and install the latest Consul plugin:

```sh
steampipe plugin install consul
```

### Credentials

| Item        | Description                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Consul requires an Address or Address and [Token](https://developer.hashicorp.com/consul/docs/security/acl/acl-tokens) for all requests.                                                              |
| Permissions | The permission scope of tokens is set by the Admin at the creation time of the ACL tokens.                                                                                                            |
| Radius      | Each connection represents a single Consul Installation.                                                                                                                                              |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/consul.spc`)<br />2. Credentials specified in environment variables, e.g., `CONSUL_HTTP_ADDR` and `CONSUL_HTTP_TOKEN`. |

### Configuration

Installing the latest consul plugin will create a config file (`~/.steampipe/config/consul.spc`) with a single connection named `consul`:

Configure your account details in `~/.steampipe/config/consul.spc`:

```hcl
connection "consul" {
  plugin = "consul"

  # `address`(required) - The address of the Consul server.
  # Can also be set with the CONSUL_HTTP_ADDR environment variable.
  # address = "http://52.14.112.248:8500"

  # `token`(optional) - The ACL token. It is required for ACL-enabled Consul servers.
  # For more information on the ACL Token, please see https://developer.hashicorp.com/consul/docs/security/acl/acl-tokens.
  # Can also be set with the CONSUL_HTTP_TOKEN or CONSUL_TOKEN environment variable.
  # token = "c178b810-8b18-6f38-016f-725ddec5d58"

  # `namespace`(optional) - This feature requires HashiCorp Cloud Platform (HCP) or self-managed Consul Enterprise. This parameter is not required in case of non-Enterprise.
  # API will execute with default namespace if this parameter is not set.
  # Can also be set with the CONSUL_NAMESPACE environment variable.
  # "*" indicates all the namespaces available.
  namespace = "*"

  # `partition`(optional) - This feature requires HashiCorp Cloud Platform (HCP) or self-managed Consul Enterprise. This parameter is not required in case of non-Enterprise.
  # API will execute with default partition if this parameter is not set.
  # Can also be set with the CONSUL_PARTITION environment variable.
  # partition = "default"
}
```

## Configuring Consul Credentials

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

Alternatively, you can also use the standard Consul environment variable to obtain credentials **only if the `address`, `token`, `namespace`, and `partition` is not specified** in the connection:

```sh
export CONSUL_HTTP_ADDR=http://18.118.144.168:4646
export CONSUL_NAMESPACE=*
export CONSUL_HTTP_TOKEN=c178b810-8b18-6f38-016f-725ddec5d58
export CONSUL_PARTITION=default
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-consul
- Community: [Slack Channel](https://steampipe.io/community/join)
