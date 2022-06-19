output "template" {
  value = templatefile("${path.module}/zk-config-template.yaml.tftpl", {
    # General
    zookeeper_replicas       = var.replicas
    zookeeper_root_log_level = var.root_log_level

    # Resources
    limit_cpu      = local.resources.limit_cpu
    limit_memory   = local.resources.limit_memory
    request_cpu    = local.resources.requested_cpu
    request_memory = local.resources.requested_memory

    # JVM
    jvm_xmx = local.jvm.xmx
    jvm_xms = local.jvm.xms

    # Annotations
    pod_annotations = var.pod_annotations

    #Node Affinity
    node_affinity_key   = local.affinity.node_affinity_key
    node_affinity_value = local.affinity.node_affinity_value

    #Service Annotations
    service_annotations = local.service_annotations

    #Storage
    storage_class        = local.storage.class
    storage_size         = local.storage.size
    storage_delete_claim = local.storage.delete_claim

    # ConfigMap
    configmap_name = var.configmap_name
    configmap_key  = var.configmap_key
  })
}


