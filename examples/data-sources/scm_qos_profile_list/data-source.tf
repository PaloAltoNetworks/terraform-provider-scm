#
# Data source to retrieve a list of SCM QoS Profile Signature objects.
#

# Example 1: Fetch a list of all SCM QoS Profile in the "Service Connections" folder.
# Folder must be one of [Remote Networks, Service Connections]"
data "scm_qos_profile_list" "all_shared" {
  folder = "Service Connections"
}

# Output the list of all SSCM QoS Profiles from the "Service Connections" folder.
output "scm_qos_profile_list_all_shared" {
  description = "A list of all SCM QoS Profiles from the Service Connections folder."
  value       = data.scm_qos_profile_list.all_shared.data
}
