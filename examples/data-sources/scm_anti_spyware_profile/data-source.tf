# Look up anti-spyware-profile by ID
data "scm_anti_spyware_profile" "scm_anti_spyware_profile_ds" {
  id = "7720ab77-d9fe-42c1-8001-6ef2202aae8c"
}

# Output values of anti-spyware-profile
output "scm_anti_spyware_profile_output" {
  value = {
    profile_id            = data.scm_anti_spyware_profile.scm_anti_spyware_profile_ds.id
    folder                = data.scm_anti_spyware_profile.scm_anti_spyware_profile_ds.folder
    name                  = data.scm_anti_spyware_profile.scm_anti_spyware_profile_ds.name
    description           = data.scm_anti_spyware_profile.scm_anti_spyware_profile_ds.description
    cloud_inline_analysis = data.scm_anti_spyware_profile.scm_anti_spyware_profile_ds.cloud_inline_analysis
  }
}