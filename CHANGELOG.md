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
