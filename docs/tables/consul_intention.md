# Table: consul_intention

Intentions control traffic communication between services at the network layer, also called L4 traffic, or the application layer, also called L7 traffic. The protocol that the destination service uses to send and receive traffic determines the type of authorization the intention can enforce.

## Examples

### Basic info

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

### List intentions with highest precedence

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
