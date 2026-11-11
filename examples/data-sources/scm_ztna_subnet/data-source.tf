# Look up a ZTNA subnet rule by its ID.
data "scm_ztna_subnet" "example" {
  oid = "44444444-4444-4444-4444-444444444444"
}

output "subnet_output" {
  value = {
    oid         = data.scm_ztna_subnet.example.oid
    name        = data.scm_ztna_subnet.example.name
    description = data.scm_ztna_subnet.example.description
    ip_subnets  = data.scm_ztna_subnet.example.ip_subnets
    group       = data.scm_ztna_subnet.example.group
    app_enabled = data.scm_ztna_subnet.example.app_enabled
  }
}
