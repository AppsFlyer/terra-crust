# Terra Crust

## What It Does 
Terra crust is a tool that creates a [Root Module](https://metallb.universe.tf) from existing Modules and exposes all of the modules `variables`,`locals` and `main` 
Terra crust is a tool to create generic module from existing Terraform modules and exporting its 3 components, Local.tf , Variables.tf, Main.tf.
* local file - create module_local.tf that export all of the locals of each sub module.
* variables file - create module_variables.tf that export all of the variables of each sub module.
* main file - create module_main.tf , exporting all of required variables as commented values in order to fill in and optional values are genereated automatically by default, also supports a external main template for development.


## How it works
Terra crust is going all over the modules folder that is given, while extracting from the `Variables.tf` files the defaults values, so make sure every variable that is `Optional` has a default in the variables section, even if there is a merge on the locals that already contains the default.
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