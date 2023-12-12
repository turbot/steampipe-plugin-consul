---
title: "Steampipe Table: consul_acl_auth_method - Query Consul ACL Auth Methods using SQL"
description: "Allows users to query Consul ACL Auth Methods, specifically the configuration and rules of each method, providing insights into access control within the Consul service."
---

# Table: consul_acl_auth_method - Query Consul ACL Auth Methods using SQL

Consul ACL Auth Method is a feature within HashiCorp Consul that enables authentication of entities through different methods such as Kubernetes, JWT, or OIDC. It provides a way to define how to authenticate an entity and produce a set of Consul ACL Tokens upon successful authentication. Consul ACL Auth Method aids in managing and controlling user access and privileges within the Consul service.

## Table Usage Guide

The `consul_acl_auth_method` table provides insights into the ACL Auth Methods within HashiCorp Consul. As a security engineer, explore method-specific details through this table, including the type of authentication method, its configuration, and associated rules. Utilize it to uncover information about methods, such as those using Kubernetes or JWT, the configuration parameters for each method, and the verification of rules associated with each method.

**Important Notes**
- You need to specify the `token` parameter in the `consul.spc` file to be able to query this table.

## Examples

### Basic info
Explore the authorization methods used within your Consul ACL system. This helps to assess the security settings and identify any modifications or partitions, enhancing overall system management and integrity.

```sql+postgres
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

```sql+sqlite
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
Explore the authentication methods currently active within the default namespace. This information can be useful for assessing security configuration and identifying potential vulnerabilities.

```sql+postgres
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
  namespace = 'default';
```

```sql+sqlite
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
  namespace = 'default';
```

### List auth methods with global token locality
Discover the authentication methods that have a global scope. This can be useful for understanding the distribution and reach of different authorization methods across your network.

```sql+postgres
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

```sql+sqlite
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
Assess the configuration details of authentication methods to gain insights into their settings, which can help in managing and enhancing security protocols.

```sql+postgres
select
  name,
  namespace,
  partition,
  jsonb_pretty(config) as config
from
  consul_acl_auth_method;
```

```sql+sqlite
select
  name,
  namespace,
  partition,
  config
from
  consul_acl_auth_method;
```