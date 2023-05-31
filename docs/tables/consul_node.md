# Table: consul_node

A node refers to a physical or virtual machine that runs a Consul agent. Each Consul node runs a Consul agent process, which is responsible for communicating with other nodes in the Consul cluster to exchange information about available services, their health status, and other metadata. The agent also listens for requests from external clients and applications, and provides various APIs and interfaces to allow for interaction with the Consul cluster.

## Examples

### Basic info

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
