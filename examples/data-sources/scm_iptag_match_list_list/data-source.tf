# List all iptag match lists in the ngfw-shared folder
data "scm_iptag_match_list_list" "example" {
  folder = "ngfw-shared"
}

# Output the name and folder for each iptag match list
output "iptag_match_list_list_results" {
  description = "List of iptag match list names and folders"
  value = {
    for item in data.scm_iptag_match_list_list.example.data : item.id => {
      name   = item.name
      folder = item.folder
    }
  }
}
