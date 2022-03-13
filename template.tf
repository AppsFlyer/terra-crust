# auto generated
variable "cruise_control" {
  description = <<EOF
  (Optional) Cruise Control Module will be used by default.
  [Readme](https://gitlab.appsflyer.com/real-time-platform/af-rti-iac/modules/strimzi/-/blob/feature/creating-terraform-modules/terraform/modules/cruise_control/README.md)
  EOF
  type = object({
    is_enabled      = optional(bool)
    root_log_level  = optional(string)
    resources       = optional(map(string))
    jvm             = optional(map(string))
    broker_capacity = optional(map(string))
    default_goals   = optional(list(string))
    hard_goals      = optional(list(string))
  })
  default = {
    enabled        = true
    root_log_level = "INFO"
    resources = {
      requested_memory = "512Mi"
      requested_cpu    = "0.5"
      limit_cpu        = "2"
      limit_memory     = "2Gi"
    }
    jvm = {
      xms = "512m"
      xmx = "2048m"
    }
    broker_capacity = {
      cpu_utilization  = 100
      inbound_network  = "10000KB/s"
      outbound_network = "10000KB/s"
    }
    default_goals = ["com.linkedin.kafka.cruisecontrol.analyzer.goals.DiskCapacityGoal"]
    hard_goals    = ["com.linkedin.kafka.cruisecontrol.analyzer.goals.DiskCapacityGoal"]
  }
  # TODO OPA Policy on Dev Side... can't be use because we use object that override keys...
  # validation {
  #   condition     = alltrue([for key in keys(var.cruise_control) : contains(["is_enabled", "root_log_level", "resources", "jvm", "broker_capacity", "default_goals", "hard_goals"], key)])
  #   error_message = "Cruise Control variable include invalid key, keys can be use: ['is_enabled', 'root_log_level', 'resources', 'jvm', 'broker_capacity', 'default_goals', 'hard_goals'] ."
  # }
}
