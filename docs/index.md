---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/consul.svg"
brand_color: "#E03875"
display_name: "Consul"
short_name: "consul"
description: "Steampipe plugin to query nodes, ACLs, services and more from Consul."
og_description: "Query Consul with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/consul-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Consul + Steampipe

[Consul](https://www.consul.io/) is a service networking solution to automate network configurations, discover services, and enable secure connectivity across any cloud or runtime.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

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

  # Address is required for requests. Required.
  # This can also be set via the CONSUL_HTTP_ADDR environment variable.
  # address = "http://52.14.112.248:8500"

  # An ACL token is required for ACL-enabled Consul servers. Optional.
  # For more information on the ACL Token, please see https://developer.hashicorp.com/consul/docs/security/acl/acl-tokens.
  # This can also be set via the CONSUL_HTTP_TOKEN environment variable.
  # token = "c178b810-8b18-6f38-016f-725ddec5d58"

  # Namespace is required for Consul Enterprise access. Optional.
  # API will execute with default namespace if this parameter is not set.
  # This can also be set via the CONSUL_NAMESPACE environment variable.
  # "*" indicates all the namespaces available.
  # namespace = "*"

  # Partition is required for Consul Enterprise access. Optional.
  # API will execute with default partition if this parameter is not set.
  # This can also be set via the CONSUL_PARTITION environment variable.
  # partition = "default"
}
```

- `token` parameter is only required to query the ACL tables like `consul_acl_auth_method`, `consul_acl_binding_rule`, `consul_acl_policy`, `consul_acl_role` and `consul_acl_token` tables.
- `namespace`, and `partition` parameters are only required to query the `consul_namespace` table.

Alternatively, you can also use the standard Consul environment variable to obtain credentials **only if other arguments (`address`, `token`, `namespace`, and `partition`) are not specified** in the connection:

```sh
export CONSUL_HTTP_ADDR=http://18.118.144.168:4646
export CONSUL_HTTP_TOKEN=c178b810-8b18-6f38-016f-725ddec5d58
export CONSUL_NAMESPACE=*
export CONSUL_PARTITION=default
```


