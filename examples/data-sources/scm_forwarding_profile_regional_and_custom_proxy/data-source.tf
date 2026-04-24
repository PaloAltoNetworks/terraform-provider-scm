# Look up the gp-and-pac basic proxy by its ID.
data "scm_forwarding_profile_regional_and_custom_proxy" "scm_forwarding_profile_regional_and_custom_proxy_1_ds" {
  id = "0e1c887a-ecf7-4a12-990c-bf1213ef81fd"
}

# Look up the gp-and-pac dual proxy by its ID.
data "scm_forwarding_profile_regional_and_custom_proxy" "scm_forwarding_profile_regional_and_custom_proxy_2_ds" {
  id = "64ecc1e6-8bb6-4689-8f69-67a4553d4da8"
}

# Look up the ztna-agent basic proxy by its ID.
data "scm_forwarding_profile_regional_and_custom_proxy" "scm_forwarding_profile_regional_and_custom_proxy_3_ds" {
  id = "8a3401e6-01dc-4fd4-b273-b0070c8fe9bb"
}

# Look up the complete ztna-agent by its ID.
data "scm_forwarding_profile_regional_and_custom_proxy" "scm_forwarding_profile_regional_and_custom_proxy_4_ds" {
  id = "7de1edaf-b9dc-4b65-b8c4-c44a14f1095b"
}

# Output various attributes from the found objects to verify the lookups were successful.
output "forwarding_profile_regional_and_custom_proxy_data_source_results" {
  description = "The details of the forwarding profile regional and custom proxy objects read from the data sources."
  value = {
    gp_and_pac_basic_proxy = data.scm_forwarding_profile_regional_and_custom_proxy.scm_forwarding_profile_regional_and_custom_proxy_1_ds
    gp_and_pac_dual_proxy  = data.scm_forwarding_profile_regional_and_custom_proxy.scm_forwarding_profile_regional_and_custom_proxy_2_ds
    ztna_agent_basic_proxy = data.scm_forwarding_profile_regional_and_custom_proxy.scm_forwarding_profile_regional_and_custom_proxy_3_ds
    ztna_agent_full_proxy  = data.scm_forwarding_profile_regional_and_custom_proxy.scm_forwarding_profile_regional_and_custom_proxy_4_ds
  }
}
