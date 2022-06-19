<img src="./docs/images/logo/terra-crust.png" width="400">

----
Terra crust was created to allow platform teams to expose terraform as the main API communication with the developers, it gives a simple interface in Terraform.
Terraform was adopted as our primary language because it is intuitive, declarative, and community driven.
We wanted to find a way to have an excellent user experience while increasing the platform productivity, moreover to keep it accessible to use and decrease management overhead. 
Our goal in developing terra crust was to hide the system's complexity and provide a simple interface to the client based on Facade Pattern and KISS principles.

<img src="./docs/images/diagram/terra-crust-levels.png" width="800">


## How it works

Terra crust creates a **root module** from existing modules and exposes all the `variables`,`locals` and `main`, it creates generic module from existing Terraform modules and exporting its 3 components, Local.tf , Variables.tf, Main.tf.
* **local** -  module_local.tf exports the locals of each sub module.
* **variables** - module_variables.tf exports the variables of each sub module.
* **main** - module_main.tf , exports the required variables as commented values in order to fill in and optional values are genereated automatically by default, also supports a external main template for development.

Terra crust is going over the module folders and extracting from the `Variables.tf` files the defaults values, so make sure every variable that is `Optional` has a default in the variables section, even if there is a merge on the locals that already contains the default.
Terra crust will release 3 files that wraps all of the existing modules folder under a general module, at the end of the flow it will run Terraform FMT to format the files.

## Output Examples:

### Variables:
![](/img/vars.jpg)

### Locals:
![](/img/locals.jpg)

### Main:
![](/img/main.jpg)

## Commands:
### Create Main:
```
./main.go terraform-main --destination-path="." --source-path=".../modules"
```
* Main has additional flag: ``main-template-path`` to support external main templates like in examples/templates/main.tf.tmpl
### Create Variables:
```
./main terraform-variables --destination-path="." --source-path=".../modules"
```
### Create Locals:
```
./main terraform-locals --destination-path="." --source-path=".../modules"
```
### Create All:
```
./main terraform-all --destination-path="." --source-path=".../modules"
```
* same as Main has additional flag: ``main-template-path`` to support external main templates like in examples/templates/main.tf.tmpl
* Upon failing on create one of the files , It wont fail the entire flow(Will keep on to the next files).
## Contribution
In order to contribute, please make sure to test and validate that everything is working, including lint ci, open a merge request and wait for approval.