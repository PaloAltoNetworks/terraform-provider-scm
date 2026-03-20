# List all config match lists in the ngfw-shared folder
data "scm_config_match_list_list" "example" {
  folder = "ngfw-shared"
}

# Output the name and folder for each config match list
output "config_match_list_list_results" {
  description = "List of config match list names and folders"
  value = {
    for item in data.scm_config_match_list_list.example.data : item.id => {
      name   = item.name
      folder = item.folder
    }
  }
}
