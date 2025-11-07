# -----------------------------------------------------------------------------
# 6. DATA SOURCE: Fetch a list of Interface Management Profiles
# This data source retrieves multiple Interface Management Profiles from SCM.
# -----------------------------------------------------------------------------
data "scm_interface_management_profile_list" "all_mgmt_profiles" {
  # Use 'folder' to limit the scope of the search.
  limit  = 50
  folder = "All"
}

# -----------------------------------------------------------------------------
# 7. OUTPUT: Display the fetched Interface Management Profiles list data
# -----------------------------------------------------------------------------
output "fetched_mgmt_profile_list_summary" {
  description = "Summary of Interface Management Profiles retrieved."
  value = {
    # 'total' gives the total number of profiles matching the criteria
    count_of_profiles_fetched = data.scm_interface_management_profile_list.all_mgmt_profiles.total
    # 'data' is the list of profile objects. Access the 'name' of the first one.
    first_profile_name = data.scm_interface_management_profile_list.all_mgmt_profiles.data[0].name
    data               = data.scm_interface_management_profile_list.all_mgmt_profiles.data
  }
}