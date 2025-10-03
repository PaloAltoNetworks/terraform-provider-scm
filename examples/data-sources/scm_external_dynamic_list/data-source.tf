# Data source to look up a single external dynamic list by its ID.
data "scm_external_dynamic_list" "scm_external_dynamic_list_ds" {
  id = "ce39b7b5-f5bc-4276-9fe5-be56613e37ad"
}

# Output the details of the looked-up external dynamic list.
output "external_dynamic_list_details" {
  value = {
    id          = data.scm_external_dynamic_list.scm_external_dynamic_list_ds.id
    name        = data.scm_external_dynamic_list.scm_external_dynamic_list_ds.name
    folder      = data.scm_external_dynamic_list.scm_external_dynamic_list_ds.folder
    description = data.scm_external_dynamic_list.scm_external_dynamic_list_ds.type.domain.description
    url         = data.scm_external_dynamic_list.scm_external_dynamic_list_ds.type.domain.url
    recurring   = data.scm_external_dynamic_list.scm_external_dynamic_list_ds.type.domain.recurring
  }
}