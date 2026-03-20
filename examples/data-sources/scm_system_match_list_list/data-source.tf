# List all system match lists in the ngfw-shared folder
data "scm_system_match_list_list" "example" {
  folder = "ngfw-shared"
}

# Output the name and folder for each system match list
output "system_match_list_list_results" {
  description = "List of system match list names and folders"
  value = {
    for item in data.scm_system_match_list_list.example.data : item.id => {
      name   = item.name
      folder = item.folder
    }
  }
}
