#
# Data source to retrieve a single SCM QoS Profile object by its ID.
#

# Replace the ID with the UUID of the QoS Profile you want to find.
data "scm_qos_profile" "scm_qos_prof" {
  id = "cffecf78-b3b1-4b01-ad31-c69bf839850b"
}

# Output the details of the single QoS Profile object found.
output "scm_qos_profile_details" {
  value = {
    id     = data.scm_qos_profile.scm_qos_prof.id
    folder = data.scm_qos_profile.scm_qos_prof.folder
    name   = data.scm_qos_profile.scm_qos_prof.name
  }
}