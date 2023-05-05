# Table: consul_acl_token

Tokens are artifacts in the ACL system used to authenticate users, services, and Consul agents. When ACLs are enabled, entities requesting access to a resource must include a token that has been linked with a policy, service identity, or node identity that grants permission to the resource. The ACL system checks the token and grants or denies access to resource based on the associated permissions.

## Examples

### Basic info

```sql
select
  accessor_id,
  secret_id,
  auth_method,
  local,
  create_time,
  expiration_ttl,
  namespace,
  partition
from
  consul_acl_token;
```

### List local tokens

```sql
select
  accessor_id,
  secret_id,
  auth_method,
  local,
  create_time,
  expiration_ttl
from
  consul_acl_token
where
  local;
```

### List tokens which will not expire

```sql
select
  accessor_id,
  secret_id,
  auth_method,
  local,
  create_time,
  expiration_ttl
from
  consul_acl_token
where
  expiration_time is null;
```

### List tokens which are not associated with any role

```sql
select
  accessor_id,
  secret_id,
  auth_method,
  local,
  create_time,
  expiration_ttl
from
  consul_acl_token
where
  roles is null;
```

### List tokens which are not associated with any auth method

```sql
select
  accessor_id,
  secret_id,
  auth_method,
  local,
  create_time,
  expiration_ttl
from
  consul_acl_token
where
  auth_method = '';
```
