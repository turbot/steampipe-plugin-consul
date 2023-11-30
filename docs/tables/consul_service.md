---
title: "Steampipe Table: consul_service - Query OCI Consul Services using SQL"
description: "Allows users to query Consul Services in OCI, specifically the service details, providing insights into service configurations and status."
---

# Table: consul_service - Query OCI Consul Services using SQL

Consul is a service networking solution to connect and secure services across any runtime platform and public or private cloud. It provides a full featured control plane with service discovery, configuration, and segmentation functionality. Given that, Consul enables rapid deployment, configuration, and maintenance of service-oriented architectures at a massive scale.

## Table Usage Guide

The `consul_service` table provides insights into Consul Services within Oracle Cloud Infrastructure (OCI). As a network administrator, explore service-specific details through this table, including service configurations, status, and associated metadata. Utilize it to uncover information about services, such as those with specific configurations, the relationships between services, and the verification of service status.

## Examples

### Basic info
Explore the various services within your network infrastructure to identify their associated data centers and namespaces. This can help you better understand the overall structure and organization of your services, which is useful in managing and optimizing your network resources.

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
Explore which services are present in the default namespace. This can help in understanding the configuration and organization of your services, and identify any areas for optimization or restructuring.

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
Uncover the details of services that are not linked with any health checks. This can be useful for identifying potential vulnerabilities in your system, as these services may not be monitored for issues or failures.

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
Explore which services are associated with a specific tag to better manage and organize your resources. This can be particularly useful for identifying patterns or issues related to a specific service category.

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
Explore which services are actively running within a specific node for effective management and monitoring. This is particularly useful for identifying any potential issues or inconsistencies within your network infrastructure.

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
Explore which services allow for tag overrides, useful for identifying areas where service tagging behavior can be customized to suit specific requirements. This can be beneficial in managing and organizing your services more effectively.

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