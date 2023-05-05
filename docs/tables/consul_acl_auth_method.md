# Table: consul_acl_auth_method

An auth method is a component in Consul that performs authentication against a trusted external party to authorize the creation of an ACL tokens usable within the local datacenter.

## Examples

### Basic info

```sql
select
  name,
  type,
  namespace,
  create_index,
  modify_index,
  partition
from
  consul_acl_auth_method;
```

### List auth methods present in default namespace

```sql
select
  name,
  type,
  namespace,
  create_index,
  modify_index,
  partition
from
  consul_acl_auth_method;
where
  namespace = 'default';
```

### List auth methods with global token locality

```sql
select
  name,
  type,
  namespace,
  create_index,
  modify_index,
  partition
from
  consul_acl_auth_method
where
  token_locality = 'global';
```

### Get config details of auth methods

```sql
select
  name,
  namespace,
  partition,
  jsonb_pretty(config) as config
from
  consul_acl_auth_method;
```
