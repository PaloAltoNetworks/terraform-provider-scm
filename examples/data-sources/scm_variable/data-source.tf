# Look up the "$example" variable by its name.
data "scm_variable" "scm_variable_ds" {
  id = "66cbe56c-0300-4905-8455-d384978a0081"
}

# Output the details of the found variable to verify they were read correctly.
output "variable_outputs" {
  value = {
    example_id    = data.scm_variable.scm_variable_ds.id
    example_name  = data.scm_variable.scm_variable_ds.name
    example_type  = data.scm_variable.scm_variable_ds.type
    example_value = data.scm_variable.scm_variable_ds.value
  }
}