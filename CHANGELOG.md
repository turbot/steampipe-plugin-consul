## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#27](https://github.com/turbot/steampipe-plugin-consul/pull/27))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#27](https://github.com/turbot/steampipe-plugin-consul/pull/27))

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#19](https://github.com/turbot/steampipe-plugin-consul/pull/19))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#19](https://github.com/turbot/steampipe-plugin-consul/pull/19))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-consul/blob/main/docs/LICENSE). ([#19](https://github.com/turbot/steampipe-plugin-consul/pull/19))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#18](https://github.com/turbot/steampipe-plugin-consul/pull/18))

## v0.1.2 [2023-12-07]

_Bug fixes_

- Fixed the invalid Go module path of the plugin. ([#15](https://github.com/turbot/steampipe-plugin-consul/pull/15))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#8](https://github.com/turbot/steampipe-plugin-consul/pull/8))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#6](https://github.com/turbot/steampipe-plugin-consul/pull/6))
- Recompiled plugin with Go version `1.21`. ([#6](https://github.com/turbot/steampipe-plugin-consul/pull/6))

## v0.0.1 [2023-06-06]

_What's new?_

- New tables added
  - [consul_acl_auth_method](https://hub.steampipe.io/plugins/turbot/consul/tables/consul_acl_auth_method)
  - [consul_acl_binding_rule](https://hub.steampipe.io/plugins/turbot/consul/tables/consul_acl_binding_rule)
  - [consul_acl_policy](https://hub.steampipe.io/plugins/turbot/consul/tables/consul_acl_policy)
  - [consul_acl_role](https://hub.steampipe.io/plugins/turbot/consul/tables/consul_acl_role)
  - [consul_acl_token](https://hub.steampipe.io/plugins/turbot/consul/tables/consul_acl_token)
  - [consul_intention](https://hub.steampipe.io/plugins/turbot/consul/tables/consul_intention)
  - [consul_key](https://hub.steampipe.io/plugins/turbot/consul/tables/consul_key)
  - [consul_namespace](https://hub.steampipe.io/plugins/turbot/consul/tables/consul_namespace)
  - [consul_node](https://hub.steampipe.io/plugins/turbot/consul/tables/consul_node)
  - [consul_service](https://hub.steampipe.io/plugins/turbot/consul/tables/consul_service)
