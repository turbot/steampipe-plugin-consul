---
title: "Steampipe Table: consul_key - Query Consul Key-Value Store using SQL"
description: "Allows users to query Consul Key-Value Store, specifically keys and their corresponding values, providing insights into the configuration and service discovery data."
---

# Table: consul_key - Query Consul Key-Value Store using SQL

Consul is a service networking solution to automate network configurations, discover services, and enable secure connectivity across any cloud or runtime. The key-value store feature in Consul is a flexible and dynamic configuration store that can be used for a wide variety of purposes, including storing configuration for service discovery and orchestration, leader election, distributed semaphore, and more. It is an integral part of Consul and is installed with the Consul agent.

## Table Usage Guide

The `consul_key` table provides insights into the keys and their corresponding values stored in the Consul Key-Value Store. As a DevOps engineer, explore key-specific details through this table, including the key name, its corresponding value, and associated metadata. Utilize it to uncover information about keys, such as their current values, flags, and the sessions they are associated with.

## Examples

### Basic info
Explore which sessions are associated with a specific key in your system. This can help you understand and manage the distribution and assignment of resources within your network.

```sql+postgres
select
  key,
  session,
  create_index,
  lock_index,
  namespace,
  partition
from
  consul_key;
```

```sql+sqlite
select
  key,
  session,
  create_index,
  lock_index,
  namespace,
  partition
from
  consul_key;
```

### List keys present in default namespace
Explore the keys present within the default namespace to understand their session contexts, creation indices, lock indices, and partitions. This can be beneficial in assessing the elements within the default namespace and understanding their configuration.

```sql+postgres
select
  key,
  session,
  create_index,
  lock_index,
  namespace,
  partition
from
  consul_key
where
  namespace = 'default';
```

```sql+sqlite
select
  key,
  session,
  create_index,
  lock_index,
  namespace,
  partition
from
  consul_key
where
  namespace = 'default';
```

### Show key value details
Explore which key-value pairs are being used in your Consul sessions. This can help you understand the configuration and data distribution within your system, allowing for more efficient management and troubleshooting.

```sql+postgres
select
  key,
  session,
  namespace,
  jsonb_pretty(value) as value
from
  consul_key;
```

```sql+sqlite
select
  key,
  session,
  namespace,
  value
from
  consul_key;
```