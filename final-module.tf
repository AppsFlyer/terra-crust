variable "zookeeper" {
  description = <<EOT
	(Optional) zookeeper Module will be used by default.
	[Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/master/terraform/modules/zookeeper/README.md)
	EOT
  type = object({
    jvm                 = optional(map(string))
    pod_annotations     = optional(map(string))
    replicas            = optional(number)
    resources           = optional(map(string))
    root_log_level      = optional(string)
    service_annotations = optional(map(string))
  })
  default = {
    jvm = {
      xms = "4096m"
      xmx = "4096m"
    }
    pod_annotations = {
      "ad.datadoghq.com/zookeeper.check_names" : "['zk']"
      "ad.datadoghq.com/zookeeper.init_configs" : "[{}]"
      "ad.datadoghq.com/zookeeper.instances" : "[{ 'host' : '%%host%%', 'port' : '2181' }]"
      "ad.datadoghq.com/zookeeper.logs" : "[{ 'source' : 'zookeeper', 'service' : 'zookeeper' }]"
    }
    replicas = 3
    resources = {
      requested_memory = "5Gi"
      requested_cpu    = "1"
      limit_cpu        = "2"
      limit_memory     = "6Gi"
    }
    root_log_level = "INFO"
    service_annotations = {
      "consul.hashicorp.com/service-sync" : "true"
    }
  }
}
variable "kafka_exporter" {
  description = <<EOT
	(Optional) kafka_exporter Module will be used by default.
	[Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/master/terraform/modules/kafka_exporter/README.md)
EOT
  type = object({
    group_regex            = optional(string)
    is_sarama_logs_enabled = optional(bool)
    resources              = optional(map(string))
    root_log_level         = optional(string)
    topic_regex            = optional(string)
  })
  default = {
    group_regex            = ".*"
    is_sarama_logs_enabled = true
    resources = {
      requested_memory = "256Mi"
      requested_cpu    = "0.5"
      limit_cpu        = "1"
      limit_memory     = "512Mi"
    }
    root_log_level = "debug"
    topic_regex    = ".*"
  }
}
variable "cruise_control" {
  description = <<EOT
	(Optional) cruise_control Module will be used by default.
	[Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/master/terraform/modules/cruise_control/README.md)
EOT
  type = object({
    broker_capacity     = optional(map(string))
    default_goals       = optional(list(string))
    global_config       = optional(map(string))
    hard_goals          = optional(list(string))
    is_enabled          = optional(bool)
    jvm                 = optional(map(string))
    resources           = optional(map(string))
    root_log_level      = optional(string)
    service_annotations = optional(map(string))
  })
  default = {
    broker_capacity = {
      cpu_utilization : 100
      inbound_network : "10000KB/s"
      outbound_network : "10000KB/s"
    }
    default_goals = [
      "com.linkedin.kafka.cruisecontrol.analyzer.goals.DiskCapacityGoal"
    ]
    global_config = {
      "webserver.http.cors.enabled"       = "true",
      "webserver.http.cors.origin"        = "'*'",
      "webserver.http.cors.exposeheaders" = "User-Task-ID,Content-Type"
      "webserver.security.enable"         = "false"
      "webserver.ssl.enable"              = "false"
      "cpu.balance.threshold"             = "1.1"
      "metadata.max.age.ms"               = "60000"
      "send.buffer.bytes"                 = "131072"
    }
    hard_goals = [
      "com.linkedin.kafka.cruisecontrol.analyzer.goals.DiskCapacityGoal"
    ]
    is_enabled = true
    jvm = {
      xms = "512m"
      xmx = "2048m"
    }
    resources = {
      requested_memory = "512Mi"
      requested_cpu    = "0.5"
      limit_cpu        = "2"
      limit_memory     = "2Gi"
    }
    root_log_level = "INFO"
    service_annotations = {
      "consul.hashicorp.com/service-sync" : "true"
    }
  }
}
variable "entity_operator" {
  description = <<EOT
	(Optional) entity_operator Module will be used by default.
	[Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/master/terraform/modules/entity_operator/README.md)
EOT
  type = object({
    reconciliation_interval_seconds = optional(number)
    resources                       = optional(map(string))
    root_log_level                  = optional(string)
  })
  default = {
    reconciliation_interval_seconds = 120
    resources = {
      requested_memory = "256Mi"
      requested_cpu    = "0.5"
      limit_cpu        = "1"
      limit_memory     = "512Mi"
    }
    root_log_level = "INFO"
  }
}
variable "kafka" {
  description = <<EOT
	(Optional) kafka Module will be used by default.
	[Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/master/terraform/modules/kafka/README.md)
EOT
  type = object({
    bootstrap_node_port = optional(number)
    jvm                 = optional(map(string))
    kafka_config        = optional(map(string))
    kafka_version       = optional(string)
    pod_annotations     = optional(map(string))
    ports               = optional(map(string))
    replicas            = optional(number)
    resources           = optional(map(string))
    root_log_level      = optional(string)
    service_annotations = optional(map(string))
  })
  default = {
    bootstrap_node_port = 32300
    jvm = {
      xms = "8192m"
      xmx = "8192m"
    }
    kafka_config = {
      "auto.create.topics.enable" : "true"
      "delete.topic.enable" : "true"
      "controlled.shutdown.enable" : "true"
      "num.io.threads" : "8"
      "num.network.threads" : "8"
      "num.recovery.threads.per.data.dir" : "1"
      "num.replica.fetchers" : "4"
      "socket.send.buffer.bytes" : "1048576"
      "socket.receive.buffer.bytes" : "1048576"
      "socket.request.max.bytes" : "104857600"
      "offsets.topic.replication.factor" : "3"
      "transaction.state.log.replication.factor" : "3"
      "transaction.state.log.min.isr" : "1"
      "default.replication.factor" : "1"
      "min.insync.replicas" : "1"
      "inter.broker.protocol.version" : "'3.0'"
      "replica.selector.class" : "org.apache.kafka.common.replica.RackAwareReplicaSelector"
      "log.flush.scheduler.interval.ms" : "2000"
    }
    kafka_version = "3.0.0"
    pod_annotations = {
      "ad.datadoghq.com/kafka.check_names" : "['kafka']"
      "ad.datadoghq.com/kafka.init_configs" : "[{}]"
      "ad.datadoghq.com/kafka.instances" : "[{ 'host' : '%%host%%', 'port' : '9999' }]"
      "ad.datadoghq.com/kafka.logs" : "[{ 'source' : 'kafka', 'service' : 'kafka' }]"
    }
    ports = {
      "internal_port" : 9092
      "external_port" : 9094
    }
    replicas = 3
    resources = {
      requested_memory = "60Gi"
      requested_cpu    = "6"
      limit_cpu        = "8"
      limit_memory     = "63Gi"
    }
    root_log_level = "INFO"
    service_annotations = {
      "consul.hashicorp.com/service-sync" : "true"
    }
  }
}
variable "consul_sync" {
  description = <<EOT
	(Optional) consul_sync Module will be used by default.
	[Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/master/terraform/modules/consul_sync/README.md)
EOT
  type = object({
    atomic                 = optional(bool)
    chart                  = optional(string)
    chart_version          = optional(string)
    cleanup_on_fail        = optional(bool)
    consul_datacenter      = optional(string)
    consul_image           = optional(string)
    consul_servers_address = optional(string)
    create_consul_sync     = optional(bool)
    create_namespace       = optional(bool)
    name                   = optional(string)
    namespace              = optional(string)
    repository             = optional(string)
    resources              = optional(map(string))
    timeout                = optional(number)
    wait                   = optional(bool)
    wait_for_jobs          = optional(bool)
    watch_any_namespace    = optional(bool)
  })
  default = {
    atomic                 = true
    chart                  = "consul"
    chart_version          = "0.39.0"
    cleanup_on_fail        = true
    consul_datacenter      = "dev-euw1-general"
    consul_image           = "consul:1.11.0"
    consul_servers_address = "dev-euw1-general.consul.appsflyer.platform"
    create_consul_sync     = true
    create_namespace       = true
    name                   = "consul"
    namespace              = "consul"
    repository             = "https://helm.releases.hashicorp.com"
    resources = {
      requested_memory = "512Mi"
      requested_cpu    = "0.5"
      limit_cpu        = "2"
      limit_memory     = "2Gi"
    }
    timeout             = 300
    wait                = true
    wait_for_jobs       = true
    watch_any_namespace = true
  }
}
variable "kowl" {
  description = <<EOT
	(Optional) kowl Module will be used by default.
	[Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/master/terraform/modules/kowl/README.md)
EOT
  type = object({
    atomic              = optional(bool)
    chart               = optional(string)
    chart_version       = optional(string)
    cleanup_on_fail     = optional(bool)
    create_kowl         = optional(bool)
    create_namespace    = optional(bool)
    name                = optional(string)
    repository          = optional(string)
    root_log_level      = optional(string)
    service_annotations = optional(map(string))
    timeout             = optional(number)
    wait                = optional(bool)
  })
  default = {
    atomic           = true
    chart            = "kowl"
    chart_version    = "2.3.0"
    cleanup_on_fail  = true
    create_kowl      = true
    create_namespace = true
    name             = "kowl"
    repository       = "https://raw.githubusercontent.com/cloudhut/charts/master/archives"
    root_log_level   = "info"
    service_annotations = {
      "consul.hashicorp.com/service-sync" : "true"
    }
    timeout = 300
    wait    = true
  }
}
variable "storage_local_provisioner" {
  description = <<EOT
	(Optional) storage_local_provisioner Module will be used by default.
	[Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/master/terraform/modules/storage_local_provisioner/README.md)
EOT
  type = object({
    atomic                           = optional(bool)
    chart                            = optional(string)
    cleanup_on_fail                  = optional(bool)
    create_local_storage_provisioner = optional(bool)
    create_namespace                 = optional(bool)
    local_provisioner_hostdir        = optional(string)
    name                             = optional(string)
    namespace                        = optional(string)
    storage_class                    = optional(string)
    timeout                          = optional(number)
    wait                             = optional(bool)
  })
  default = {
    atomic                           = true
    chart                            = "vendors/provisioner"
    cleanup_on_fail                  = true
    create_local_storage_provisioner = true
    create_namespace                 = true
    local_provisioner_hostdir        = "/mnt"
    name                             = "local-storage-provisioner"
    namespace                        = "storage-provisioner"
    storage_class                    = "nvme-ssd"
    timeout                          = 300
    wait                             = true
  }
}

