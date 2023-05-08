# Table: consul_service

In Consul, a service refers to a logical endpoint that represents a single instance or a group of instances of a specific application or microservice. Services are a key concept in Consul's service discovery functionality, allowing clients and applications to locate and communicate with other services within the Consul cluster.

## Examples

### Basic info

```sql
select
  id,
  node,
  address,
  datacenter,
  service_id,
  service_name,
  namespace
from
  consul_service;
```

### List services present in default namespace

```sql
select
  id,
  node,
  address,
  datacenter,
  service_id,
  service_name,
  namespace
from
  consul_service
where
  namespace = 'default';
```

### List services which are not associated with any health checks

```sql
select
  id,
  node,
  address,
  datacenter,
  service_id,
  service_name,
  namespace
from
  consul_service
where
  checks is null;
```

### List services with a specific tag

```sql
select
  id,
  node,
  address,
  datacenter,
  service_id,
  service_name,
  namespace
from
  consul_service
where
  service_tag = 'http';
```

### List services running in a specific node

```sql
select
  id,
  node,
  address,
  datacenter,
  service_id,
  service_name,
  namespace
from
  consul_service
where
  node = 'ip-172-25-34-191';
```

### List services where service tags can be overridden

```sql
select
  id,
  node,
  address,
  datacenter,
  service_id,
  service_name,
  namespace
from
  consul_service
where
  service_enable_tag_override;
```
