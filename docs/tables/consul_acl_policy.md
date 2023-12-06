---
title: "Steampipe Table: consul_acl_policy - Query Consul ACL Policies using SQL"
description: "Allows users to query Consul ACL Policies, specifically the details of each policy, providing insights into the permissions and rules set for each policy."
---

# Table: consul_acl_policy - Query Consul ACL Policies using SQL

Consul ACL (Access Control List) Policies are a set of rules that control the actions that a client can perform in a Consul cluster. These policies are used to restrict the operations a token can perform based on its assigned policies. ACL Policies provide a flexible way to control access to data and APIs.

## Table Usage Guide

The `consul_acl_policy` table provides insights into ACL Policies within HashiCorp's Consul. As a DevOps engineer, explore policy-specific details through this table, including the rules, descriptions, and associated metadata. Utilize it to uncover information about policies, such as their specific rules, the actions they allow or deny, and to verify the overall security configuration of your Consul cluster.

**Important Notes**
- You need to specify the `token` parameter in the `consul.spc` file to be able to query this table.

## Examples

### Basic info
Explore the specific policies within your ACL system to understand their rules, descriptions, and indices. This can help in managing access control and ensuring security within your system.

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
Explore which policies are present in the default namespace, allowing you to assess the elements within your system's default settings. This can help you maintain better control over your system's security and access rules.

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
Discover the segments that consist of policies linked to ACL tokens. This is useful for understanding the security measures in place and managing access control within your system.

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