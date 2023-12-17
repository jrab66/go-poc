terraform {
  backend "gcs" {
    bucket = "demo-deel-01"
    prefix = "prod"
  }
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.9.0"
    }
  }
}




provider "google" {
  project = var.project
  region  = var.region
}