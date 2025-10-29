# Look up zone by its ID.
data "scm_zone" "scm_zone_ds" {
  id = "50f1f0f3-a420-4989-9770-c927f1467a9a"
}

# Output various attributes from the found zone to verify the lookups were successful.
output "zone_data_source_results" {
  description = "The details of the zone read from the data source."
  value = {
    id                           = data.scm_zone.scm_zone_ds.id
    name                         = data.scm_zone.scm_zone_ds.name
    network                      = data.scm_zone.scm_zone_ds.network
    enable_device_identification = data.scm_zone.scm_zone_ds.enable_device_identification
    enable_user_identification   = data.scm_zone.scm_zone_ds.enable_user_identification
    user_acl                     = data.scm_zone.scm_zone_ds.user_acl
    device_acl                   = data.scm_zone.scm_zone_ds.device_acl
    folder                       = data.scm_zone.scm_zone_ds.folder
  }
}