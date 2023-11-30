---
title: "Steampipe Table: consul_node - Query Consul Nodes using SQL"
description: "Allows users to query Nodes in Consul, providing access to information about the nodes, their health, and their associated services."
---

# Table: consul_node - Query Consul Nodes using SQL

Consul is a service networking solution to connect and secure services across any runtime platform and public or private cloud. It provides a full-featured control plane with service discovery, configuration, and segmentation functionality. Consul Nodes are the basic units in a Consul cluster, and this table provides information about these nodes.

## Table Usage Guide

The `consul_node` table provides insights into Nodes within HashiCorp's Consul service networking solution. As a DevOps engineer, you can explore node-specific details through this table, including their health, associated services, and other metadata. Utilize it to uncover information about nodes, such as their status, the services they're associated with, and their overall health within the cluster.

## Examples

### Basic info
Explore which nodes are part of your Consul datacenter to understand their addresses and partition details. This can help manage your datacenter more effectively by allowing you to determine the areas in which changes have been made recently.

```sql
select
  id,
  node,
  address,
  datacenter,
  create_index,
  modify_index,
  partition
from
  consul_node;
```

### List nodes from a specific datacenter
Analyze the settings to understand the distribution of nodes within a specific datacenter. This can help in managing and optimizing resources effectively across different datacenters.

```sql
select
  id,
  node,
  address,
  datacenter,
  create_index,
  modify_index,
  partition
from
  consul_node
where
  datacenter = 'dc1';
```

### List nodes with a specific tag
Explore which nodes are associated with a specific deployment tag. This can be useful in managing and organizing your resources based on their deployment status.

```sql
select
  id,
  node,
  address
  datacenter,
  create_index,
  modify_index,
  partition
from
  consul_node
where
  meta->> 'deployment' = '2';
```