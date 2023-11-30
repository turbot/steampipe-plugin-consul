---
title: "Steampipe Table: consul_intention - Query Consul Intentions using SQL"
description: "Allows users to query Consul Intentions, providing insights into service access control and potential communication paths."
---

# Table: consul_intention - Query Consul Intentions using SQL

Consul Intentions are a resource within HashiCorp Consul that allows you to define access controls, which dictate what services may communicate. They are used to control which services may establish connections, providing a way to manage service-to-service communication in a microservices architecture. Intentions are a crucial component of Consul's service mesh capabilities.

## Table Usage Guide

The `consul_intention` table provides insights into Consul Intentions within HashiCorp Consul. As a network engineer or a security administrator, explore intention-specific details through this table, including source and destination services, action, and associated metadata. Utilize it to uncover information about intentions, such as those allowing or denying certain communication paths, and the verification of service-to-service access controls.

## Examples

### Basic info
Gain insights into the communication intentions between different services in your network. This query helps identify potential areas of improvement or points of failure, by analyzing the source and destination of each interaction.

```sql
select
  id,
  created_at,
  source_name,
  source_ns,
  destination_name,
  destination_ns
from
  consul_intention;
```

### List intentions with default source namespace
Explore which intentions have been set with the default source namespace. This can be useful for understanding the default configurations and identifying areas for potential adjustment or optimization.

```sql
select
  id,
  created_at,
  source_name,
  source_ns,
  destination_name,
  destination_ns
from
  consul_intention
where
  source_ns = 'default';
```

### List intentions in order of highest precedence
Explore the priorities of different intentions in your system by arranging them in descending order of importance. This can help you understand the hierarchy and manage your resources more effectively.

```sql
select
  id,
  created_at,
  source_name,
  source_ns,
  destination_name,
  destination_ns
from
  consul_intention
order by
  precedence desc;
```

### List intentions with destination applied to all namespaces
Discover the intentions that have a destination applied to all namespaces. This is useful for understanding the broad application of policies and permissions across your system.

```sql
select
  id,
  created_at,
  source_name,
  source_ns,
  destination_name,
  destination_ns
from
  consul_intention
where
  destination_ns = '*';
```

### List allowlist intentions
Discover the segments that have been given access permissions. This query is useful in identifying and analyzing the areas where access has been explicitly granted for better security management.

```sql
select
  id,
  created_at,
  source_name,
  source_ns,
  destination_name,
  destination_ns
from
  consul_intention
where
  action = 'allow';
```

### List intentions with deny permission
Discover the segments that have been denied access within your network infrastructure. This can be useful for security audits, identifying potential weak points, or understanding the overall security structure.

```sql
select
  id,
  created_at,
  source_name,
  source_ns,
  destination_name,
  destination_ns,
  jsonb_pretty(p) as permission
from
  consul_intention,
  jsonb_array_elements(permissions) as p
where
  p ->> 'Action' = 'deny';
```