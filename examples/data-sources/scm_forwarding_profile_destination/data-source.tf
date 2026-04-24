# Look up the forwarding profile destination with FQDN by its ID.
data "scm_forwarding_profile_destination" "scm_forwarding_profile_destination_1_ds" {
  id = "47c43e8b-eedf-483a-bfd8-5dfcf7a2462b"
}

# Look up the forwarding profile destination with IP addresses by its ID.
data "scm_forwarding_profile_destination" "scm_forwarding_profile_destination_2_ds" {
  id = "7796420b-c45e-4ca2-a609-91ab83bbd219"
}

# Look up the forwarding profile destination with mixed types by its ID.
data "scm_forwarding_profile_destination" "scm_forwarding_profile_destination_3_ds" {
  id = "b3730821-0692-4245-96b5-3474651ed2b5"
}

# Output various attributes from the found objects to verify the lookups were successful.
output "forwarding_profile_destination_data_source_results" {
  description = "The details of the forwarding profile destination objects read from the data sources."
  value = {
    fqdn_destination  = data.scm_forwarding_profile_destination.scm_forwarding_profile_destination_1_ds
    ip_destination    = data.scm_forwarding_profile_destination.scm_forwarding_profile_destination_2_ds
    mixed_destination = data.scm_forwarding_profile_destination.scm_forwarding_profile_destination_3_ds
  }
}