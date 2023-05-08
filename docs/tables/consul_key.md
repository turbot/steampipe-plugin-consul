# Table: consul_key

Consul KV is a core feature of Consul and is installed with the Consul agent. Once installed with the agent, it will have reasonable defaults. Consul KV allows users to store indexed objects, though its main uses are storing configuration parameters and metadata.

## Examples

### Basic info

```sql
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

```sql
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

```sql
select
  key,
  session,
  namespace,
  jsonb_pretty(value) as value
from
  consul_key;
```
