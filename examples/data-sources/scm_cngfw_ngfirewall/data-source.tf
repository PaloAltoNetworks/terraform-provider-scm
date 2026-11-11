
data "scm_cngfw_ngfirewall" "example" {
  id     = var.firewall_id
  region = var.aws_region
}

# Environment variable to pick firewall ID from the environment.
variable "firewall_id" {
  type        = string
  description = "The ID of the Cloud NGFW firewall to read."
}

# Environment variable to pick AWS region from the environment.
variable "aws_region" {
  type        = string
  description = "The AWS region where the firewall is deployed (e.g. us-east-1)."
}

output "ngfirewall" {
  description = "The Cloud NGFW firewall details."
  value       = data.scm_cngfw_ngfirewall.example
}
