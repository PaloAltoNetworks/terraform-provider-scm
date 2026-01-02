resource "scm_qos_profile" "scm_qos_profile_1" {
  folder = "Remote Networks"
  name   = "scm-qos-profile-base"
}

resource "scm_qos_profile" "scm_qos_profile_2" {
  folder = "Remote Networks"
  name   = "scm-qos-profile-2"
  aggregate_bandwidth = {
    egress_max        = 200
    egress_guaranteed = 10000
  }
}

resource "scm_qos_profile" "scm_qos_profile_3" {
  folder = "Service Connections"
  name   = "scm-qos-profile-3"
  aggregate_bandwidth = {
    egress_guaranteed = 20
  }
  class_bandwidth_type = {
    mbps = {
      class = [
        {
          name = "class1"
        },
        {
          name     = "class2"
          priority = "high"
        },
        {
          name     = "class3"
          priority = "real-time"
          class_bandwidth = {
            egress_max = 500
          }
        },
        {
          name     = "class4"
          priority = "low"
          class_bandwidth = {
            egress_guaranteed = 60000
          }
        },
        {
          name     = "class5"
          priority = "medium"
          class_bandwidth = {
            egress_max        = 955
            egress_guaranteed = 50000
          }
        }
      ]
    }
  }
}

resource "scm_qos_profile" "scm_qos_profile_4" {
  folder = "Service Connections"
  name   = "scm-qos-profile-4"
  aggregate_bandwidth = {
    egress_max = 1400
  }
  class_bandwidth_type = {
    percentage = {
      class = [
        {
          name = "class1"
        },
        {
          name     = "class2"
          priority = "low"
        },
        {
          name     = "class3"
          priority = "real-time"
          class_bandwidth = {
            egress_max = 100
          }
        },
        {
          name     = "class4"
          priority = "medium"
          class_bandwidth = {
            egress_guaranteed = 5
          }
        },
        {
          name     = "class5"
          priority = "high"
          class_bandwidth = {
            egress_max        = 25
            egress_guaranteed = 50
          }
        }
      ]
    }
  }
}