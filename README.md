# Terra Crust

Sub module wrapper was created after encountering major issue in terraform which is not solved, terraform best practice is to create a flat hierchy but thats create a problem when we want to export the new module as a single module, terraform wrapper generate module out of folder of modules wrap them all under a single terraform module, exporting the variables of any type including maps and regular defaults supporting in the ability to create:
* local file - create module_local.tf that export all of the locals of each sub module.
* variables file - create module_variables.tf that export all of the variables of each sub module.
* main file - create module_main.tf , exporting all of required variables as commented values in order to fill in and optional values are genereated automatically by default, also supports a external main template for development.


## Project Folder Structure
Folder structure is simple, in order to add another functionality, api related things will go into `cmd` folder while all the buisness logic is in `internal`.
Using Drivers design for everything that is 3rd party such as consul and kafka-manager(cmak).


## Commands:
### Create Main:
`go run main.go terraform-main --destination-path="." --source-path=".../modules"`

* Main has additional flag: ``main-template-path`` to support external main templates like in examples/templates/main.tf.tmpl

### Create Variables:
`go run main.go terraform-variables --destination-path="." --source-path=".../modules"`

### Create Locals:
`go run main.go terraform-locals --destination-path="." --source-path=".../modules"`

### Create All:
`go run main.go terraform-all --destination-path="." --source-path=".../modules"`

* same as Main has additional flag: ```main-template-path``` to support external main templates like in examples/templates/main.tf.tmpl

## Contribution
In order to contribute, please make sure to test and validate that everything is working, including lint ci, open a merge request and wait for approval.