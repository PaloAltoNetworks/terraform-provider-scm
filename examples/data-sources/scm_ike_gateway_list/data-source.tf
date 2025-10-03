/**
 * Data source to retrieve a list of IKE Gateways.
 * You can filter the list by folder, snippet, or device.
 * This example retrieves all gateways in the "Shared" folder.
 */
data "scm_ike_gateway_list" "example_list_ike_gateway_ds" {
  folder = "Remote Networks"
}

/**
 * Output the total number of IKE Gateways found by the data source.
 */
output "ike_gateway_list_total" {
  value     = data.scm_ike_gateway_list.example_list_ike_gateway_ds.total
  sensitive = true
}

/**
 * Output the list of data for all retrieved IKE Gateways.
 * This provides the configuration details for each gateway in the list.
 * $ terraform output -json ike_gateway_list_data can help you view the details.
 */
output "ike_gateway_list_data" {
  value     = data.scm_ike_gateway_list.example_list_ike_gateway_ds.data
  sensitive = true
}