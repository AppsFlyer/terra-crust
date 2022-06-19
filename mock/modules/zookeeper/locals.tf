locals {

  resources = merge(
    tomap({
      requested_memory = "5Gi"
      requested_cpu    = "1"
      limit_cpu        = "2"
      limit_memory     = "6Gi"
    }),
    var.resources,
  )

  jvm = merge(
    tomap({
      xms = "4096m"
      xmx = "4096m"
    }),
    var.jvm,
  )

  pod_annotations = merge(
    tomap({
      "ad.datadoghq.com/zookeeper.check_names" : "['zk']"
      "ad.datadoghq.com/zookeeper.init_configs" : "[{}]"
      "ad.datadoghq.com/zookeeper.instances" : "[{ 'host' : '%%host%%', 'port' : '2181' }]"
      "ad.datadoghq.com/zookeeper.logs" : "[{ 'source' : 'zookeeper', 'service' : 'zookeeper' }]"
    }),
  var.pod_annotations)

  affinity = merge(
    tomap({
      node_affinity_key   = "dedicated"
      node_affinity_value = "Zookeeper"
    }),
    var.affinity,
  )

  service_annotations = merge(
    tomap({
      "consul.hashicorp.com/service-sync" : "true"
      "consul.hashicorp.com/service-name" : "strimzi-zookeeper-${var.cluster_name}"
    }),
    var.service_annotations,
  )

  storage = merge(
    tomap({
      size         = "69Gi"
      class        = "nvme-ssd"
      delete_claim = "true"
    }),
    var.storage,
  )

}
