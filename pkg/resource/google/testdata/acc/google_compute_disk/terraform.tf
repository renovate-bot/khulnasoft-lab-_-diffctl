provider "google" {}

terraform {
  required_version = "~> 0.15.0"
  required_providers {
    google = {
      version = "3.90.1"
    }
  }
}

resource "google_compute_disk" "default" {
    name  = "test-disk"
    zone  = "us-central1-a"
    image = "debian-9-stretch-v20200805"
}
