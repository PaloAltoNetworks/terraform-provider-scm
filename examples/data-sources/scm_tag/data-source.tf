# Look up the "production" tag by its name.
data "scm_tag" "scm_tag_ds" {
  id = "66cbe56c-0300-4905-8455-d384978a0082"
}

# Output the details of the found tags to verify they were read correctly.
output "tag_outputs" {
  value = {
    production_id      = data.scm_tag.scm_tag_ds.id
    production_color   = data.scm_tag.scm_tag_ds.color
    production_folder  = data.scm_tag.scm_tag_ds.folder
    production_snippet = data.scm_tag.scm_tag_ds.snippet
    production_device  = data.scm_tag.scm_tag_ds.device
  }
}