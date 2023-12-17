variable "project" {
  type    = string
  default = "deel-demo-408323"
}

variable "region" {
  type    = string
  default = "us-central1"
}

# GKE

variable "gke_num_nodes" {
  default     = 1
  description = "number of gke nodes"
}