# Look up a ZTNA subnet rule by its ID.
data "ztna_subnet" "example" {
  oid = "44444444-4444-4444-4444-444444444444"
}

output "subnet_output" {
  value = {
    oid         = data.ztna_subnet.example.oid
    name        = data.ztna_subnet.example.name
    description = data.ztna_subnet.example.description
    ip_subnets  = data.ztna_subnet.example.ip_subnets
    group       = data.ztna_subnet.example.group
    app_enabled = data.ztna_subnet.example.app_enabled
  }
}
