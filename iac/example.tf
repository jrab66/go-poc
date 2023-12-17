# module "kubernetes-engine" {
#   source  = "terraform-google-modules/kubernetes-engine/google"
#   version = "29.0.0"
#   # insert the 6 required variables here
# }


# resource "google_storage_bucket" "static-site4" {
#   name          = "dummy-1111-bucket"
#   location      = "US"
#   force_destroy = true

#   uniform_bucket_level_access = true

# }