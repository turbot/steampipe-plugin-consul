connection "consul" {
  plugin = "consul"

  # `address`(required) - The address of the Consul server.
  # Can also be set with the CONSUL_HTTP_ADDR environment variable.
  # address = "http://52.14.112.248:8500"

  # `token`(optional) - The ACL token. It is required for ACL-enabled Consul servers.
  # For more information on the ACL Token, please see https://developer.hashicorp.com/consul/docs/security/acl/acl-tokens.
  # Can also be set with the CONSUL_HTTP_TOKEN environment variable.
  # token = "c178b810-8b18-6f38-016f-725ddec5d58"

  # `namespace`(optional) - This feature requires HashiCorp Cloud Platform (HCP) or self-managed Consul Enterprise. This parameter is not required in case of non-Enterprise access.
  # API will execute with default namespace if this parameter is not set.
  # Can also be set with the CONSUL_NAMESPACE environment variable.
  # "*" indicates all the namespaces available.
  # namespace = "*"

  # `partition`(optional) - This feature requires HashiCorp Cloud Platform (HCP) or self-managed Consul Enterprise. This parameter is not required in case of non-Enterprise access.
  # API will execute with default partition if this parameter is not set.
  # Can also be set with the CONSUL_PARTITION environment variable.
  # partition = "default"
}
