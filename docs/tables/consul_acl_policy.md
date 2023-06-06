# Table: consul_acl_policy

A policy is a group of one or more ACL rules that are linked to ACL tokens.

You need to specify the `token` parameter in the `consul.spc` file to be able to query this table.

## Examples

### Basic info

```sql
select
  id,
  name,
  rules,
  description,
  create_index,
  modify_index,
  namespace,
  partition
from
  consul_acl_policy;
```

### List policies that are present in default namespace

```sql
select
  id,
  name,
  rules,
  description,
  create_index,
  modify_index,
  namespace,
  partition
from
  consul_acl_policy
where
  namespace = 'default';
```

### List policies which are attached to ACL tokens

```sql
select
  id,
  name,
  rules,
  description,
  create_index,
  modify_index
from
  consul_acl_policy
where
  id in
  (
    select
      p ->> 'ID'
    from
      consul_acl_token,
      jsonb_array_elements(policies) as p
  );
```
