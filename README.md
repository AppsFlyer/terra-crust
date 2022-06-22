![Build](https://github.com/danielvrog/terra-crust/actions/workflows/golang-build-test-coverage.yml/badge.svg?branch=main) [![Go Report Card](https://goreportcard.com/badge/github.com/AppsFlyer/terra-crust)](https://goreportcard.com/report/github.com/AppsFlyer/terra-crust) [![codecov](https://codecov.io/gh/AppsFlyer/terra-crust/branch/main/graph/badge.svg)](https://codecov.io/gh/AppsFlyer/terra-crust)

<img src="./docs/images/logo/terra-crust.png" width="400">

----
Terra crust was created to allow platform teams to expose terraform as the main API communication with the developers, by giving a simple interface in Terraform.
Terraform was adopted as our primary language because it is intuitive, declarative, and community driven.
We wanted to find a way to have an excellent user experience while increasing the platform productivity, moreover to keep it accessible to use and decrease management overhead. 
Our goal in developing terra crust was to hide the system's complexity and provide a simple interface to the client based on Facade Pattern and KISS principles.

<img src="./docs/images/diagram/terra-crust-levels.png" width="800">


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

## Example Usage
Further guidelines as well as examples, can be shown [here](./examples/example.md).

## Contribution
In order to contribute, please make sure to test and validate that everything is working, including lint ci, open a merge request and wait for approval.