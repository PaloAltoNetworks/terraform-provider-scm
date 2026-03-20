# List all hipmatch match lists in the ngfw-shared folder
data "scm_hipmatch_match_list_list" "example" {
  folder = "ngfw-shared"
}

# Output the name and folder for each hipmatch match list
output "hipmatch_match_list_list_results" {
  description = "List of hipmatch match list names and folders"
  value = {
    for item in data.scm_hipmatch_match_list_list.example.data : item.id => {
      name   = item.name
      folder = item.folder
    }
  }
}
