resource "scm_anti_spyware_signature" "scm_anti_spyware_signature_1" {
  folder     = "All"
  threat_id  = 6900001
  comment    = "Managed by Terraform"
  direction  = "client2server"
  severity   = "critical"
  threatname = "Example Threat"
  default_action = {
    alert = {}
  }
  signature = {
    combination = {
      and_condition = [
        {
          name = "And Condition 1"
          or_condition = [
            {
              "name" : "Test",
              "threat_id" : "10001"
            }
          ]
        }
      ]
      order_free = false
      time_attribute = {
        interval  = 3600
        threshold = 60
        track_by  = "source"
      }
    }
  }
}
