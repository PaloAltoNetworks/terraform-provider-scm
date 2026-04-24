# Look up a basic Global Protect proxy forwarding profile by its ID.
data "scm_forwarding_profile" "gp_proxy_basic" {
  id = "a1b2c3d4-e5f6-4a5b-8c9d-0e1f2a3b4c5d"
}

# Look up a full Global Protect proxy forwarding profile with block rules by its ID.
data "scm_forwarding_profile" "gp_proxy_full" {
  id = "f1e2d3c4-b5a6-4758-9192-a3b4c5d6e7f8"
}

# Look up a PAC file based forwarding profile by its ID.
data "scm_forwarding_profile" "pac_file" {
  id = "1a2b3c4d-5e6f-7890-abcd-ef1234567890"
}

# Look up a ZTNA agent forwarding profile by its ID.
data "scm_forwarding_profile" "ztna_agent" {
  id = "9f8e7d6c-5b4a-3210-fedc-ba9876543210"
}

# Output various attributes from the found objects to verify the lookups were successful.
output "forwarding_profile_data_source_results" {
  description = "The details of the forwarding profile objects read from the data sources."
  value = {
    gp_proxy_basic_object = data.scm_forwarding_profile.gp_proxy_basic
    gp_proxy_full_object  = data.scm_forwarding_profile.gp_proxy_full
    pac_file_object       = data.scm_forwarding_profile.pac_file
    ztna_agent_object     = data.scm_forwarding_profile.ztna_agent
  }
}

