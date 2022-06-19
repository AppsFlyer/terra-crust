resource "helm_release" "consul" {
  count            = var.is_enabled ? 1 : 0
  chart            = var.chart
  name             = var.name
  repository       = var.repository
  namespace        = var.namespace
  version          = var.chart_version
  create_namespace = var.create_namespace
  atomic           = var.atomic
  wait             = var.wait
  wait_for_jobs    = var.wait_for_jobs
  timeout          = var.timeout
  values = [
    templatefile("${path.module}/helm-consule-values-template.tftpl", {
      # Consul Configuration
      consul_image           = var.consul_image
      consul_datacenter      = var.consul_datacenter
      consul_servers_address = var.consul_servers_address
      chart_version          = var.chart_version

      # EKS cluster name
      eks_cluster = var.eks_cluster

      # Affinity
      kafka_node_affinity_key       = local.affinity.kafka_node_key
      kafka_node_affinity_value     = local.affinity.kafka_node_value
      zookeeper_key_affinity_key    = local.affinity.zookeeper_node_key
      zookeeper_node_affinity_value = local.affinity.zookeeper_node_value

      # Resources
      limit_cpu      = local.resources.limit_cpu
      limit_memory   = local.resources.limit_memory
      request_cpu    = local.resources.requested_cpu
      request_memory = local.resources.requested_memory
    })
  ]
}

# Need to add aws security groups for TCP and UDP - port 8301 for consul service

