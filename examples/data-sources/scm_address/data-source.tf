# Look up the IP Netmask address object by its ID.
data "scm_address" "scm_address_1_ds" {
  id = "f23e1c22-de94-44cd-b67f-36f2516618a7"
}

# Look up the IP Range address object by its ID.
data "scm_address" "scm_address_2_ds" {
  id = "662ef9a5-80d4-40d8-b51a-ec11915895d8"
}

# Look up the FQDN address object by its ID.
data "scm_address" "scm_address_3_ds" {
  id = "1b996b9d-d350-4565-8ace-b319e6ae5a34"
}

# Look up the class_c_wildcard address object by its ID.
data "scm_address" "scm_address_4_ds" {
  id = "933a1646-21fa-4edc-a0ca-868ead783ac0"
}

# Look up the multi-tag address object by its ID.
data "scm_address" "scm_address_5_ds" {
  id = "f10df339-61aa-42a0-aab8-85424bfb2a8f"
}

# Output various attributes from the found objects to verify the lookups were successful.
output "address_data_source_results" {
  description = "The details of the address objects read from the data sources."
  value = {
    netmask_object          = data.scm_address.scm_address_1_ds
    ip_range_object         = data.scm_address.scm_address_2_ds
    fqdn_object             = data.scm_address.scm_address_3_ds
    class_c_wildcard_object = data.scm_address.scm_address_4_ds
    multi_tag_test_object   = data.scm_address.scm_address_5_ds
  }
}