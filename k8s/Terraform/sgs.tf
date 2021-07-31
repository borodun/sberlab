variable "sgName" {
  description = "Super sg"
}

locals {
  rules = {
    http-rule = {
      description = "Allow HTTP from anywhere",
      protocol = "tcp",
      port = 80,
      source = "0.0.0.0/0"
    },
    https-rule = {
      description = "Allow HTTPS from anywhere",
      protocol = "tcp",
      port = 443,
      source = "0.0.0.0/0"
    },
    ssh-rule = {
      description = "Allow SSH from anywhere",
      protocol = "tcp",
      port = 22,
      source = "0.0.0.0/0"
    },
    http-rule = {
      description = "Allow HTTP to our server on port 9999",
      protocol = "tcp",
      port = 9999,
      source = "0.0.0.0/0"
    },
    lb-http-rule = {
      description = "Allow HTTP for load balancer",
      protocol = "tcp",
      port = 80,
      source = "178.170.192.219/32"
    },
    kube-api-rule = {
      description = "Allow kube-api",
      protocol = "tcp",
      port = 6443,
      source = "0.0.0.0/0"
    },
    kublr-rule = {
      description = "Allow kublr agent",
      protocol = "tcp",
      port = 11251,
      source = "0.0.0.0/0"
    },
    front-rule = {
      description = "Allow frontend",
      protocol = "tcp",
      port = 32222,
      source = "0.0.0.0/0"
    },
    back-rule = {
      description = "Allow backend",
      protocol = "tcp",
      port = 32223,
      source = "0.0.0.0/0"
    },
    grafana-rule = {
      description = " Grafana",
      protocol = "tcp",
      port = 3000,
      source = "0.0.0.0/0"
    },
    prometheus-container-rule = {
      description = "Allow container prometheus service",
      protocol = "tcp",
      port = 9090,
      source = "0.0.0.0/0"
    },
    node-exporter-rule = {
      description = "node exporter",
      protocol = "tcp",
      port = 9100,
      source = "0.0.0.0/0"
    },
    mongo-rule = {
      description = "Allow mongodb",
      protocol = "tcp",
      port = 27017,
      source = "0.0.0.0/0"
    },
    prometheus-nodeport-rule = {
      description = "node port of prometheus",
      protocol = "tcp",
      port = 30000,
      source = "0.0.0.0/0"
    },
    openfaas-rule = {
      description = "OpenFaaS gateway",
      protocol = "tcp",
      port = 31112,
      source = "0.0.0.0/0"
    },
    grafana-nodeport-rule = {
      description = "Grafana node port",
      protocol = "tcp",
      port = 32000,
      source = "0.0.0.0/0"
    }
  }
}

resource "sbercloud_networking_secgroup" "sg_01" {
  name = var.sgName
  description = "Security group with SSH, HTTP and HTTPS"
}

resource "sbercloud_networking_secgroup_rule" "sg_rule_01" {
  for_each = local.rules

  direction = "ingress"
  ethertype = "IPv4"
  description = each.value.description
  protocol = each.value.protocol
  port_range_min = each.value.port
  port_range_max = each.value.port
  remote_ip_prefix = each.value.source

  security_group_id = sbercloud_networking_secgroup.sg_01.id
}