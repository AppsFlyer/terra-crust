variable "create_consul_sync" {
  type        = bool
  description = "(Optional) Create the resource only if variable set to true, default = true"
  default     = true
}

variable "name" {
  type        = string
  description = "(Optional) The helm chart Release name."
  default     = "consul"
}

variable "chart" {
  type        = string
  description = "(Optional) Chart name to be installed. The chart name can be local path, a URL to a chart, or the name of the chart if repository is specified. It is also possible to use the <repository>/<chart> format here if you are running Terraform on a system that the repository has been added to with helm repo add but this is not recommended."
  default     = "consul"
}

variable "repository" {
  type        = string
  description = "(Optional) Repository URL where to locate the requested chart."
  default     = "https://helm.releases.hashicorp.com"
}

variable "chart_version" {
  type        = string
  description = "(Optional) Helm chart version."
  default     = "0.39.0"
}

variable "namespace" {
  type        = string
  description = "(Optional) The namespace to install the release into."
  default     = "consul"
}

variable "create_namespace" {
  type        = bool
  description = "(Optional) Create the namespace if it does not yet exist"
  default     = true
}

variable "atomic" {
  type        = bool
  description = "(Optional) If set, installation process purges chart on fail. The wait flag will be set automatically if atomic is used."
  default     = true
}

variable "wait" {
  type        = bool
  description = "(Optional) Will wait until all resources are in a ready state before marking the release as successful. It will wait for as long as timeout."
  default     = true
}

variable "wait_for_jobs" {
  type        = bool
  description = "(Optional) will wait until all Jobs have been completed before marking the release as successful."
  default     = true
}

variable "timeout" {
  type        = number
  description = "(Optional) Time in seconds to wait for any individual kubernetes operation (like Jobs for hooks)."
  default     = 300
}

variable "cleanup_on_fail" {
  type        = bool
  description = "(Optional) Allow deletion of new resources created in this upgrade when upgrade fails."
  default     = true
}

variable "watch_any_namespace" {
  type        = bool
  description = "(Optional) Tells strimzi to watch all namespaces , if used 1 strimzi operator on eks cluster then should be set to true."
  default     = true
}

variable "consul_image" {
  type        = string
  description = "(Optional) The consul image:tag name"
  default     = "consul:1.11.0"
}

variable "consul_datacenter" {
  type        = string
  description = "(Optional) The consul datacenter name."
  default     = "dev-euw1-general"
}

variable "consul_servers_address" {
  type        = string
  description = "(Optional) The consul server addresses to which the consul clients will connect to."
  default     = "dev-euw1-general.consul.appsflyer.platform"
}

variable "eks_cluster" {
  type        = string
  description = "(Required) The name of the EKS Cluster. Needed for the consul clients nodes names to be attached to the consul server."
}

variable "affinity" {
  type        = map(string)
  description = "(Required) Labels and Taints that are used on nodes with nvme storage."
}

variable "resources" {
  type        = map(string)
  description = <<EOT
  (Optional) Resource requests/limit (CPU and Memory) to specify the maximum resources that can be consumed. 
  EOT
  default = {
    requested_memory = "512Mi"
    requested_cpu    = "0.5"
    limit_cpu        = "2"
    limit_memory     = "2Gi"
  }
}
