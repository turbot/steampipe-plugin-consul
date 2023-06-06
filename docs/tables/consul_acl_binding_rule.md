# Table: consul_acl_binding_rule

An ACL binding rule is a configuration setting that associates a token with a set of policies that define what the token is allowed to do. The binding rule specifies which resources the token has access to and what actions it can perform on those resources.

You need to specify the `token` parameter in the `consul.spc` file to be able to query this table.

## Examples

### Basic info

```sql
select
  id,
  auth_method,
  bind_name,
  bind_type,
  create_index,
  namespace,
  partition
from
  consul_acl_binding_rule;
```

### List rules that are present in default namespace

```sql
select
  id,
  auth_method,
  bind_name,
  bind_type,
  create_index,
  namespace,
  partition
from
  consul_acl_binding_rule
where
  namespace = 'default';
```

### List service type binding rules

```sql
select
  id,
  auth_method,
  bind_name,
  bind_type,
  create_index,
  namespace,
  partition
from
  consul_acl_binding_rule
where
  bind_type = 'service';
```

### Show auth methods related to the binding rule

```sql
select
  a.name as auth_method_name,
  a.type as auth_method_type,
  a.create_index as auth_method_create_index,
  b.id as binding_rule_id,
  b.bind_type as bind_type
from
  consul_acl_binding_rule as b
  left join consul_acl_auth_method as a on b.auth_method = a.name;
```
