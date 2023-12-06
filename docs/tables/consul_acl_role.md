---
title: "Steampipe Table: consul_acl_role - Query OCI Consul ACL Roles using SQL"
description: "Allows users to query Consul ACL Roles, specifically the role ID, name, and description, providing insights into access control and permissions."
---

# Table: consul_acl_role - Query OCI Consul ACL Roles using SQL

Consul ACL Roles are a feature in Oracle Cloud Infrastructure's Consul service. They are used to manage permissions and access control. ACL Roles can be assigned to tokens to grant the token the permissions of the role.

## Table Usage Guide

The `consul_acl_role` table provides insights into ACL Roles within OCI Consul. As a system administrator, explore role-specific details through this table, including role ID, name, and description. Utilize it to manage and monitor access control and permissions within your OCI environment.

**Important Notes**
- You need to specify the `token` parameter in the `consul.spc` file to be able to query this table.

## Examples

### Basic info
Explore the roles within your Consul ACL system to gain insights into their creation and modification indices, as well as their associated namespaces and partitions. This is useful for understanding the structure and organization of your access control system.

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
Discover the roles that are not linked to any service identities. This can help in identifying unused roles and aid in system optimization by removing unnecessary elements.

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
Determine the access control list (ACL) policies linked to a specific ACL role. This can be helpful in managing and understanding the permissions associated with different roles within your system.

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
      consul_acl_role,
      jsonb_array_elements(policies) as p
    where
      name = 'aclRole'
  );
```

### List roles which are attached to ACL tokens
Discover the segments that have roles attached to ACL tokens to understand the user permissions and security settings in your system. This can help in managing access control and identifying potential security risks.

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