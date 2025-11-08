# 1. Resource: Create the Interface Management Profile
# This block creates the profile with your specified configuration.
resource "scm_interface_management_profile" "test_inf_mgmt_profile" {
  # Required Field
  name = "test_inf_mgmt_profile_ds_1"

  # Contextual Field
  folder = "All"

  # Permitted IP Addresses
  permitted_ip = [
    {
      name = "10.0.0.0/24"
    },
    {
      name = "10.0.0.0/32"
    }
  ]

  # Service Flags (Booleans)
  http                       = true
  https                      = false
  telnet                     = false
  ssh                        = true
  ping                       = false
  http_ocsp                  = true
  userid_service             = true
  userid_syslog_listener_ssl = true
  userid_syslog_listener_udp = true
  response_pages             = false
}

# --------------------------------------------------------------------------------

# 2. Data Source: Retrieve the Interface Management Profile by ID
# We use the resource's generated 'id' attribute to fetch the profile.
data "scm_interface_management_profile" "single_profile_by_id" {
  # The data source type needs to match the resource type, 
  # but the name is arbitrary (e.g., single_profile_by_id).
  id = scm_interface_management_profile.test_inf_mgmt_profile.id
}

# --------------------------------------------------------------------------------

# 3. Output: Display a value from the fetched data source
output "fetched_profile_name" {
  description = "The name of the single Interface Management Profile fetched by ID."
  # Access the 'name' attribute from the data source
  value = data.scm_interface_management_profile.single_profile_by_id.name
}

output "fetched_profile" {
  description = "The PING status retrieved from the fetched data source."
  # Access the 'ping' attribute from the data source
  value = data.scm_interface_management_profile.single_profile_by_id
}