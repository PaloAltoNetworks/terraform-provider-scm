# --- Step 1: Make configuration changes ---

resource "scm_address" "web_server" {
  folder      = "Prisma Access"
  name        = "tf_web_server"
  description = "Web server address"
  ip_netmask  = "10.0.1.100/32"
}

# --- Step 2: Push configuration to Prisma Access ---

# Declare the config push action. This does NOT execute automatically
# on "terraform apply" — it must be explicitly invoked.
action "scm_config_push" "deploy" {
  config {
    devices     = ["Service Connection"]
    description = "Push after Terraform changes"
  }
}

# Alternatively, attach the push action to resources so it triggers
# automatically after create/update operations.
#
# resource "scm_address" "web_server" {
#   ...
#   lifecycle {
#     action_trigger {
#       events  = [after_create, after_update]
#       actions = [action.scm_config_push.deploy]
#     }
#   }
# }

# --- Usage ---
#
# Apply resources only (no push):
#   terraform apply
#
# Push configuration after applying:
#   terraform apply -invoke=action.scm_config_push.deploy
#
# Push without any other changes:
#   terraform apply -invoke=action.scm_config_push.deploy -refresh=false
