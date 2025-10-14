# Look up anti-spyware-profile by ID
data "scm_anti_spyware_signature" "scm_anti_spyware_signature_ds" {
  id = "7720ab77-d9fe-42c1-8001-6ef2202aae8c"
}

# Output values of anti-spyware-profile
output "scm_anti_spyware_signature_output" {
  value = {
    thread_id = data.scm_anti_spyware_signature.scm_anti_spyware_signature_ds.id
    folder    = data.scm_anti_spyware_signature.scm_anti_spyware_signature_ds.folder
    name      = data.scm_anti_spyware_signature.scm_anti_spyware_signature_ds.name
    comment   = data.scm_anti_spyware_signature.scm_anti_spyware_signature_ds.comment
    severity  = data.scm_anti_spyware_signature.scm_anti_spyware_signature_ds.severity
    signature = data.scm_anti_spyware_signature.scm_anti_spyware_signature_ds.signature
  }
}