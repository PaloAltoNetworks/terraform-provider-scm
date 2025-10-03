/**
 * Data source to retrieve a single IKE Gateway by its ID.
 * Replace the placeholder ID with the actual UUID of the gateway you want to fetch.
 */
data "scm_ike_gateway" "example_singular_ike_gateway_ds" {
  id = "1ba42513-2985-4783-8bdf-c83cf20d6dd1"
}

/**
 * Output the configuration of the fetched IKE Gateway.
 * This will display all the attributes of the specific gateway.
 * $ terraform output -json ike_gateway_singular_example can help you view it
 */
output "ike_gateway_singular_example" {
  value     = data.scm_ike_gateway.example_singular_ike_gateway_ds
  sensitive = true
}