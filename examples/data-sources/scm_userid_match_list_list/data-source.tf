# List all userid match lists in the ngfw-shared folder
data "scm_userid_match_list_list" "example" {
  folder = "ngfw-shared"
}

# Output the name and folder for each userid match list
output "userid_match_list_list_results" {
  description = "List of userid match list names and folders"
  value = {
    for item in data.scm_userid_match_list_list.example.data : item.id => {
      name   = item.name
      folder = item.folder
    }
  }
}
