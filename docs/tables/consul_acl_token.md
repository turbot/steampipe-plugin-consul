---
title: "Steampipe Table: consul_acl_token - Query Consul ACL Tokens using SQL"
description: "Allows users to query Consul ACL Tokens, specifically the Access Control List (ACL) tokens, providing insights into security rules and permissions."
---

# Table: consul_acl_token - Query Consul ACL Tokens using SQL

Consul ACL Tokens are a security feature in HashiCorp Consul that provides a flexible way to control access to data and APIs. The ACL system is a Capability-based system that relies on tokens which can have fine-grained rules and policies. ACL Tokens allow you to restrict which data and APIs a client can access, ensuring secure and controlled access within your Consul environment.

## Table Usage Guide

The `consul_acl_token` table provides insights into ACL Tokens within HashiCorp Consul. As a security engineer, explore token-specific details through this table, including permissions, policies, and associated metadata. Utilize it to uncover information about ACL Tokens, such as those with unrestricted permissions, the policies associated with each token, and the verification of access rules.

**Important Notes**
- You need to specify the `token` parameter in the `consul.spc` file to be able to query this table.

## Examples

### Basic info
Explore which accessors have been authorized, when they were created, and their expiration timeline. This can help you manage access control and understand potential security risks in your system.

```sql+postgres
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

```sql+sqlite
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
Explore which access control list (ACL) tokens are locally stored within the Consul service. This is useful for managing security and access controls, particularly in understanding which tokens might expire soon.

```sql+postgres
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

```sql+sqlite
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
  local = 1;
```

### List tokens which will never expire
Identify instances where certain access tokens are set to never expire. This can be useful in managing security and access control, as perpetual tokens may pose a potential risk.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that consist of tokens not linked to any role, which can be useful to identify potential security risks or unused resources. This information can aid in streamlining your system's security and efficiency.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which tokens are not associated with any authentication methods. This can be beneficial in identifying potential security vulnerabilities or gaps in your system's access control.

```sql+postgres
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

```sql+sqlite
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