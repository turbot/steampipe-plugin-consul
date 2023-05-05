# Table: consul_acl_role

A role is a collection of policies that your ACL administrator can link to a token. They enable you to reuse policies by decoupling the policies from the token distributed to team members. Instead, the token is linked to the role, which is able to hold several policies that can be updated asynchronously without distributing new tokens to users. As a result, roles can provide a more convenient authentication infrastructure than creating unique policies and tokens for each requester.

## Examples

### Basic info

```sql
select
  id,
  name,
  description,
  create_index,
  modify_index,
  namespace,
  partition
from
  consul_acl_role;
```

### List roles which are not attached to any service identities

```sql
select
  id,
  name,
  description,
  create_index,
  modify_index,
  namespace,
  partition
from
  consul_acl_role
where
  service_identities is null;
```

### Show ACL policies attached to a particular ACL role

```sql
select
  id,
  name,
  rules,
  description,
  create_index,
  modify_index,
from
  consul_acl_policy
where
  id in
  (
    select
      p ->> 'ID'
    from
      consul_acl_role,
      jsonb_array_elements(policies) as p
    where
      name = 'aclRole'
  );
```

### List roles which are attached to ACL tokens

```sql
select
  id,
  name,
  description,
  create_index,
  modify_index
from
  consul_acl_role
where
  id in
  (
    select
      r ->> 'ID'
    from
      consul_acl_token,
      jsonb_array_elements(roles) as r
  );
```
