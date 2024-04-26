connection "consul" {
  plugin = "consul"

  # Address is required for requests. Required.
  # This can also be set via the CONSUL_HTTP_ADDR environment variable.
  # address = "http://52.14.112.248:8500"

  # An ACL token is required for ACL-enabled Consul servers. Optional.
  # For more information on the ACL Token, please see https://developer.hashicorp.com/consul/docs/security/acl/tokens/create/create-a-service-token.
  # This can also be set via the CONSUL_HTTP_TOKEN environment variable.
  # token = "c178b810-8b18-6f38-016f-725ddec5d58"

  # Namespace is required for Consul Enterprise access. Optional.
  # API will execute with default namespace if this parameter is not set.
  # This can also be set via the CONSUL_NAMESPACE environment variable.
  # "*" indicates all the namespaces available.
  # namespace = "*"

  # Partition is required for Consul Enterprise access. Optional.
  # API will execute with default partition if this parameter is not set.
  # This can also be set via the CONSUL_PARTITION environment variable.
  # partition = "default"
}
