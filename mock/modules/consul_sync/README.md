<!-- BEGIN_TF_DOCS -->
### Requirements

No requirements.

### Providers

| Name | Version |
|------|---------|
| <a name="provider_helm"></a> [helm](#provider\_helm) | n/a |

### Modules

No modules.

### Resources

| Name | Type |
|------|------|
| [helm_release.consul](https://registry.terraform.io/providers/hashicorp/helm/latest/docs/resources/release) | resource |

### Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_affinity"></a> [affinity](#input\_affinity) | (Required) Labels and Taints that are used on nodes with nvme storage. | `map(string)` | n/a | yes |
| <a name="input_eks_cluster"></a> [eks\_cluster](#input\_eks\_cluster) | (Required) The name of the EKS Cluster. Needed for the consul clients nodes names to be attached to the consul server. | `string` | n/a | yes |
| <a name="input_atomic"></a> [atomic](#input\_atomic) | (Optional) If set, installation process purges chart on fail. The wait flag will be set automatically if atomic is used. | `bool` | `true` | no |
| <a name="input_chart"></a> [chart](#input\_chart) | (Optional) Chart name to be installed. The chart name can be local path, a URL to a chart, or the name of the chart if repository is specified. It is also possible to use the <repository>/<chart> format here if you are running Terraform on a system that the repository has been added to with helm repo add but this is not recommended. | `string` | `"consul"` | no |
| <a name="input_chart_version"></a> [chart\_version](#input\_chart\_version) | (Optional) Helm chart version. | `string` | `"0.39.0"` | no |
| <a name="input_cleanup_on_fail"></a> [cleanup\_on\_fail](#input\_cleanup\_on\_fail) | (Optional) Allow deletion of new resources created in this upgrade when upgrade fails. | `bool` | `true` | no |
| <a name="input_consul_datacenter"></a> [consul\_datacenter](#input\_consul\_datacenter) | (Optional) The consul datacenter name. | `string` | `"dev-euw1-general"` | no |
| <a name="input_consul_image"></a> [consul\_image](#input\_consul\_image) | (Optional) The consul image:tag name | `string` | `"consul:1.11.0"` | no |
| <a name="input_consul_servers_address"></a> [consul\_servers\_address](#input\_consul\_servers\_address) | (Optional) The consul server addresses to which the consul clients will connect to. | `string` | `"dev-euw1-general.consul.appsflyer.platform"` | no |
| <a name="input_create_consul_sync"></a> [create\_consul\_sync](#input\_create\_consul\_sync) | (Optional) Create the resource only if variable set to true, default = true | `bool` | `true` | no |
| <a name="input_create_namespace"></a> [create\_namespace](#input\_create\_namespace) | (Optional) Create the namespace if it does not yet exist | `bool` | `true` | no |
| <a name="input_name"></a> [name](#input\_name) | (Optional) The helm chart Release name. | `string` | `"consul"` | no |
| <a name="input_namespace"></a> [namespace](#input\_namespace) | (Optional) The namespace to install the release into. | `string` | `"consul"` | no |
| <a name="input_repository"></a> [repository](#input\_repository) | (Optional) Repository URL where to locate the requested chart. | `string` | `"https://helm.releases.hashicorp.com"` | no |
| <a name="input_resources"></a> [resources](#input\_resources) | (Optional) Resource requests/limit (CPU and Memory) to specify the maximum resources that can be consumed. | `map(string)` | <pre>{<br>  "limit_cpu": "2",<br>  "limit_memory": "2Gi",<br>  "requested_cpu": "0.5",<br>  "requested_memory": "512Mi"<br>}</pre> | no |
| <a name="input_timeout"></a> [timeout](#input\_timeout) | (Optional) Time in seconds to wait for any individual kubernetes operation (like Jobs for hooks). | `number` | `300` | no |
| <a name="input_wait"></a> [wait](#input\_wait) | (Optional) Will wait until all resources are in a ready state before marking the release as successful. It will wait for as long as timeout. | `bool` | `true` | no |
| <a name="input_wait_for_jobs"></a> [wait\_for\_jobs](#input\_wait\_for\_jobs) | (Optional) will wait until all Jobs have been completed before marking the release as successful. | `bool` | `true` | no |
| <a name="input_watch_any_namespace"></a> [watch\_any\_namespace](#input\_watch\_any\_namespace) | (Optional) Tells strimzi to watch all namespaces , if used 1 strimzi operator on eks cluster then should be set to true. | `bool` | `true` | no |

### Outputs

No outputs.
<!-- END_TF_DOCS -->