# List all globalprotect match lists in the ngfw-shared folder
data "scm_globalprotect_match_list_list" "example" {
  folder = "ngfw-shared"
}

# Output the name and folder for each globalprotect match list
output "globalprotect_match_list_list_results" {
  description = "List of globalprotect match list names and folders"
  value = {
    for item in data.scm_globalprotect_match_list_list.example.data : item.id => {
      name   = item.name
      folder = item.folder
    }
  }
}
