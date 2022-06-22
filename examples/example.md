## Example Usage

Let's look at following terraform project structure:

```
    .               
    ├── main.tf           
    └── modules                   
        ├── moduleA     
        ├── moduleB 
        └── ...
``` 
Terracrust can automatically extract all used varialbes and locals block from the modules folder, to generate a new layer that combines all modules to a single, simplified, root module.
```
    .
    ├── developers_main.tf
    └── root_module
        ├── module_variables.tf                 
        ├── main.tf.tmpl      
        ├── module_locals.tf
        ├── module_main.tf                    
        └── modules                   
            ├── moduleA    
            ├── moduleB 
            └── ...
``` 

### TerraCrust Usage
let's start from the modules layer. Let's say, both modules A and B have the same following variables in them:
```hcl
variable "foo" {
  type        = string
  description = "Optional variable."
  default     = "bar"
}

variable "required" {
  type        = string
  description = "Required variable."
}
```

The main layer above will look samiliar to the following:
```hcl
module "moduleA" {
  source   = "./modules/moduleA"
  required = var.moduleA_requried_value
}

module "moduleB" {
  source   = "./modules/moduleB"
  required = var.moduleB_requried_value
}
```

When you wish to extract this structure to another layer, meaning, have a root module that contains both module A and B respectfully. \
You will need to extract their varialbes to the same main layer, including optional values configured at the module layer, so that their logical connection will not break. 

When looking at this example, it's pretty simple to do this manually, but what about cases when you have over 10 modules to manage, and each module contains 10 different/similiar variables. This can lead to managing over 100 varialbes in total, copying and pasting them with the same logical structure.

This is where TerraCrust comes in place, by providing a path to the modules folder, TerraCrust will automatically extract and organize a `module_variables.tf`, a `module_locals.tf` as well as a `module_main.tf`, with all the modules default variables in them.

Each module's variables is looked as an object, containing all it's optional variables and their optional values respectfully under `module_varialbes.tf`:
```hcl
variable "moduleA" {
  type = object({
    foo = optional(string) # The variable name and type
  })
  default = {
    foo = "bar" # The default found in the module folder
  }
}

variable "moduleB" {
  type = object({
    foo = optional(string)
  })
  default = {
    foo = "bar"
  }
}
```

`module_locals.tf`, each variable in every module object has a coalesce block that either uses the default value provided at the module level, or if a user entered a different default value, the locals will take it to account:
```hcl
locals {
  moduleA = {
    foo = coalesce(var.drain_cleaner.foo, "bar")
  }
  moduleB = {
    foo = coalesce(var.drain_cleaner.foo, "bar")
  }
}
```
`module_main.tf`, which contains all the connection to the default values, as well the required variables needed to be filled before apply will also be generated:
```hcl
module "moduleA" {
  source   = "./modules/moduleA"
  required = #TO DO: fill requried variable

  foo = local.moduleA.foo
}

module "moduleB" {
  source   = "./modules/moduleB"
  required = #TO DO: fill requried variable

  foo = local.moduleB.foo
}
```

Leaving the users to use a single ready module, combining everything in a simple way using the following structure like `developers_main.tf`:
```hcl
module "root_module" {
  source = "./root_module"
  moduleA {
    requried = "user_input"
  }
  moduleB {
    requried = "user_input"
    foo      = "foobar" # example of changing default value.
  }
}
```
Saving time for platform teams extracting every module variable manually, and providing developers a simple, singular module to use without the need to deep dive on how each module works.

### Map type values usage
When using map type variables. Use the following structure.
under `variables.tf` in the module level:
```hcl
variable "map_varialbe" {
  type = map(string)
  default = {
    "keyA" = "valueA",
    "keyB" = "ValueB",
    "KeyC" = "ValueC"
  }
}
```
Create a map type variables with it's default values.

under `locals.tf`:
```hcl
locals {
  map_varialbe = merge(
    tomap({
      "keyA" = "valueA",
      "keyB" = "ValueB",
      "KeyC" = "ValueC"
    }),
    var.map_varialbe,
  )
}
```
Create a tomap function, so when a user changes a value at the top level, it will not earase all other keys in the map.


### Working with main template usage

TerraCrust also give the ability to use a ready main.tf template, that contains all module required variables dependencies, without changing them for every run.

Simply take your ready made main.tf, with it's logical dependencies, and change the end of the file to `main.tmpl`. Add the following lines in the place you wish to add the default values to:

```hcl
module "moduleA" {
  source = "./modules/moduleA"
  required = var.moduleA_User_Input
  
  {{(GetDefaults "moduleA" .)}}
}

module "moduleB" {
  source   = "./modules/moduleB"
  required = module.moduleA.output

  {{(GetDefaults "moduleB" .)}}
}
```
The function GetDefaults, will add the optional values to each module block, without chainging the required values. It uses the name of the module folder give under `source`.

Example, this `module_main.tf` will get generated:
```hcl
module "moduleA" {
  source = "./modules/moduleA"  
  required = var.moduleA_User_Input  # Remains the same
  
  foo = local.moduleA.foo # Gets added automatically
}

module "moduleB" {
  source   = "./modules/moduleB"
  required = module.moduleA.output  # Remains the same

  foo = local.moduleB.foo # Gets added automatically
}
``` 