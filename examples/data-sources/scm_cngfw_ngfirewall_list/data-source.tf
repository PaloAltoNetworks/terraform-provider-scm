data "scm_cngfw_ngfirewall_list" "list_example" {
  region = var.aws_region
}

# Environment variable to pick AWS region from the environment.
variable "aws_region" {
  type        = string
  description = "The AWS region where the firewall is deployed (e.g. us-east-1)."
}

output "ngfirewall" {
  description = "The Cloud NGFW firewall list."
  value       = data.scm_cngfw_ngfirewall_list.list_example
}
