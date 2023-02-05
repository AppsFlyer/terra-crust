![Build](https://github.com/danielvrog/terra-crust/actions/workflows/golang-build-test-coverage.yml/badge.svg?branch=main) [![Go Report Card](https://goreportcard.com/badge/github.com/AppsFlyer/terra-crust)](https://goreportcard.com/report/github.com/AppsFlyer/terra-crust) [![codecov](https://codecov.io/gh/AppsFlyer/terra-crust/branch/main/graph/badge.svg)](https://codecov.io/gh/AppsFlyer/terra-crust)

<img src="./docs/images/logo/terra-crust.png" width="400">

----
Terra crust was created to allow platform teams to expose terraform as the main API communication with the developers, by giving a simple interface in Terraform.
Terraform was adopted as our primary language because it is intuitive, declarative, and community driven.
We wanted to find a way to have an excellent user experience while increasing the platform productivity, moreover to keep it accessible to use and decrease management overhead. 
Our goal in developing terra crust was to hide the system's complexity and provide a simple interface to the client based on Facade Pattern and KISS principles.

<img src="./docs/images/diagram/terra-crust-levels.png" width="800">

## HashiConf Europe 

Listen to a 12 minutes talk about Terra Crust at HashiConf conference Amsterdam June 2022

[![Watch the video](https://img.youtube.com/vi/_LhbL0ZRz_c/maxresdefault.jpg)](https://www.youtube.com/watch?v=_LhbL0ZRz_c)

## AWS reinvent 2022

Check the brakeout session **"Running high-throughput, real-time ad platforms in the cloud (ADM302)"** at AWS reinvent 2022 presenting Terra crust and the way we empower our developers.

[![Watch the video](https://img.youtube.com/vi/iHNgS1YKtH8/maxresdefault.jpg)](https://www.youtube.com/watch?v=iHNgS1YKtH8&t=1348s)

## How it works

Terra crust creates a **root module** from existing modules folders inside your terraform project, and exposes all the `variables`,`locals` and `main` used in each module automatically. \
It takes the following components and exports it through 3 files, `local.tf` , `variables.tf`, `main.tf`.
* **local** -  module_local.tf exports the locals of each sub module.
* **variables** - module_variables.tf exports the variables of each sub module.
* **main** - module_main.tf , exports the required variables as commented values in order to fill in by their logical connection. Optional values are genereated automatically by default, also supports an external main template for development.

Terra crust is going over the module folders and extracting from the `Variables.tf` files the defaults values for each variable. \
Every variable that is `Optional` must have a default value in it's variable block, even if there is a merge on the locals that already contains the default. \
Terra crust will release 3 files that wraps all of the existing modules folder under a general module. \
At the end of the flow it will run Terraform FMT to format the files.

## Output Examples:

### Variables:
![](/docs/images/vars.jpg)

### Locals:
![](/docs/images/locals.jpg)

### Main:
![](/docs/images/main.jpg)

## Commands:
### Create Main:
```
terra-crust terraform-main --destination-path="." --source-path=".../modules"
```
* terra-crust has additional flag: ``main-template-path`` to support external main templates like in examples/templates/main.tf.tmpl
### Create Variables:
```
terra-crust terraform-variables --destination-path="." --source-path=".../modules"
```
### Create Locals:
```
terra-crust terraform-locals --destination-path="." --source-path=".../modules"
```
### Create All:
```
terra-crust terraform-all --destination-path="." --source-path=".../modules"
```
* same as Main has additional flag: ``main-template-path`` to support external main templates like in examples/templates/main.tf.tmpl
* Upon failing on create one of the files , It wont fail the entire flow(Will keep on to the next files).

## Remote TerraCrust:

### Description
From version 2.0.0 TerraCrust will support fetching remote modules 
and composite them into the root module that is being created.
### How to use
in order to activate this feature all you have to do is to set `fetch-remote` to true like so:  
```terra-crust terraform-all  --main-template-path=./terraform/main.tmpl  --destination-path="." --source-path=".../modules" fetch-remote=true```  
When activating `FetchRemote` you must use Custom main template, terracrust will look for all sources
that are from git it will look for this pattern:
1. `"git::REPOSITORY.git/PATH/MODULE"`
2. `"git::REPOSITORY"`  
it will start with git:: then the repository has to end with .git if it has
more than 1 module inside it, if the repo is just the root module then no need to add .git to the end.
it will also support versioning/ branch reference so feel free to add them ,
more examples could be found under `examples/templates`.

Output Example for this module: 
```"git::https://github.com/terraform-aws-modules/terraform-aws-iam.git/modules/iam-account"```
```
module "iam-account" {
  source = "git::https://github.com/terraform-aws-modules/terraform-aws-iam.git/modules/iam-account"

  # account_alias = module. TODO: Add Required Field 


  hard_expiry                    = local.iam-account.hard_expiry
  require_numbers                = local.iam-account.require_numbers
  require_symbols                = local.iam-account.require_symbols
  create_account_password_policy = local.iam-account.create_account_password_policy
  allow_users_to_change_password = local.iam-account.allow_users_to_change_password
  minimum_password_length        = local.iam-account.minimum_password_length
  password_reuse_prevention      = local.iam-account.password_reuse_prevention
  require_lowercase_characters   = local.iam-account.require_lowercase_characters
  require_uppercase_characters   = local.iam-account.require_uppercase_characters
  get_caller_identity            = local.iam-account.get_caller_identity
  max_password_age               = local.iam-account.max_password_age

}
```

this will propagate all the default and required variables using the Templates Api

## Template API:
The main.tmpl exposing Template API that includes for now 2 functions:
1. GetRequired - Will expose the require variables with option for you to fill
2. GetDefaults - Will expose the optional variables - without needing to fill them up

For example:
```
module "iam-account" {
  source            = "git::https://github.com/terraform-aws-modules/terraform-aws-iam.git/modules/iam-account"

 {{(GetRequired "iam-account" .)}}

 {{(GetDefaults "iam-account" .)}}
}
```

Results can be seen at the section above, basically after you fill the required variables,
you want to get rid of `{{(GetRequired "iam-account" .)}}` because it will keep overwriting your changes
so once you filled the required variables you can drop it instead and put the required variables you filled in the template.

## Example Usage
Further guidelines as well as examples, can be shown [here](./examples/example.md).

## Contribution
In order to contribute, please make sure to test and validate that everything is working, including lint ci, open a merge request and wait for approval.
