# Table: consul_namespace

With Consul Enterprise 1.7.0+, data for different users or teams can be isolated from each other with the use of namespaces. Namespaces help reduce operational challenges by removing restrictions around uniqueness of resource names across distinct teams, and enable operators to provide self-service through delegation of administrative privileges.

You need to specify the `namespace` and `partition` parameters in the `consul.spc` file to be able to query this table.

## Examples

### Basic info

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