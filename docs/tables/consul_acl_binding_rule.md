---
title: "Steampipe Table: consul_acl_binding_rule - Query Consul ACL Binding Rules using SQL"
description: "Allows users to query ACL Binding Rules in Consul, specifically the ACL binding rule properties, providing insights into the rule configuration and status."
---

# Table: consul_acl_binding_rule - Query Consul ACL Binding Rules using SQL

Consul ACL Binding Rules is a feature in HashiCorp Consul that allows you to define how to translate trusted identities into ACL Tokens. It provides a flexible way to manage service-to-service authorization and enforce security policies across your network. Consul ACL Binding Rules helps you maintain the integrity and security of your services by ensuring that only authorized identities can interact with your system.

## Table Usage Guide

The `consul_acl_binding_rule` table provides insights into ACL Binding Rules within HashiCorp Consul. As a Security Engineer, explore rule-specific details through this table, including rule properties, associated service, and rule status. Utilize it to uncover information about rules, such as the binding rule configuration, the service associated with each rule, and the current status of each rule.

**Important Notes**
- You need to specify the `token` parameter in the `consul.spc` file to be able to query this table.

## Examples

### Basic info
Discover the segments that use different authentication methods within your system by analyzing the settings of your ACL binding rules. This can help you pinpoint specific locations where certain types of binding rules are used, aiding in system security and configuration management.

```sql+postgres
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

```sql+sqlite
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
Uncover the details of access control list (ACL) binding rules that are set in the default namespace. This is useful for auditing security configurations and ensuring that the default namespace is not overly exposed.

```sql+postgres
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

```sql+sqlite
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
Analyze the settings to understand the binding rules associated with a particular service type. This can help in managing access control lists (ACLs) more effectively by pinpointing specific services.

```sql+postgres
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

```sql+sqlite
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
Analyze the settings to understand the relationship between authorization methods and binding rules, which can help in identifying any inconsistencies or discrepancies in access control. This is particularly useful in enhancing security measures and ensuring proper access management.

```sql+postgres
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

```sql+sqlite
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