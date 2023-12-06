---
title: "Steampipe Table: consul_namespace - Query Consul Namespaces using SQL"
description: "Allows users to query Consul Namespaces, specifically details related to each namespace including their name, description, ACLs, and metadata."
---

# Table: consul_namespace - Query Consul Namespaces using SQL

A Consul Namespace is a resource within HashiCorp Consul that allows you to isolate Consul resources and configurations into separate, distinct entities. It provides a way to segment the Consul ecosystem into smaller, manageable parts for different teams, applications, or environments. Consul Namespaces helps in ensuring the right level of access and control over the Consul resources.

## Table Usage Guide

The `consul_namespace` table provides insights into Consul Namespaces within HashiCorp Consul. As a DevOps engineer, you can explore namespace-specific details through this table, including their names, descriptions, ACLs, and associated metadata. Utilize it to uncover information about namespaces, such as their access controls, configurations, and the metadata associated with each namespace.

**Important Notes**
- You need to specify the `namespace` and partition parameters in the `consul.spc` file to be able to query this table.

## Examples

### Basic info
Explore the namespaces within your Consul environment to understand their creation and modification indices, which can help in tracking changes and managing your resources effectively.

```sql
select
  name,
  create_index,
  description,
  modify_index,
  partition
from
  consul_namespace;
```

### List deleted namespaces
Discover the segments that were previously created but have since been removed. This is beneficial in assessing the changes in your system's organization and structure over time.

```sql
select
  name,
  create_index,
  description,
  modify_index,
  partition
from
  consul_namespace
where
  deleted_at is not null;
```

### Show ACLs of each namespace
Assess the access control lists (ACLs) for each namespace to understand their policy and role defaults, which can be useful for auditing security configurations and permissions.

```sql
select
  name,
  create_index,
  partition,
  jsonb_pretty(acls -> 'PolicyDefaults') as policy_defaults,
  jsonb_pretty(acls -> 'RoleDefaults') as role_defaults
from
  consul_namespace;
```