locals {
  affinity = merge(
    tomap({
      kafka_node_key       = "dedicated"
      kafka_node_value     = "Kafka"
      zookeeper_node_key   = "dedicated"
      zookeeper_node_value = "Zookeeper"
    }),
    var.affinity,
  )

  resources = merge(
    tomap({
      requested_memory = "512Mi"
      requested_cpu    = "0.5"
      limit_cpu        = "2"
      limit_memory     = "2Gi"
    }),
    var.resources,
  )
}