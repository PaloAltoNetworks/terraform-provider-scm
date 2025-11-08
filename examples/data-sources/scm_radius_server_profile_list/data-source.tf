data "scm_radius_server_profile_list" "paged_profiles_list" {
  # Filters the list of profiles to those within the 'All' folder scope.
  folder = "All"
  # Sets the maximum number of results per page/fetch.
  limit = 10
}

output "fetched_profile_list_summary" {
  description = "Summary of profiles retrieved by the SCM Radius Profile List data source."
  value = {
    # CORRECTED: Accessing 'paged_profiles_list' and its 'total' attribute.
    count_of_profiles_fetched = data.scm_radius_server_profile_list.paged_profiles_list.total
    first_profile_name        = try(data.scm_radius_server_profile_list.paged_profiles_list.data[0].name, "List is empty")
  }
}