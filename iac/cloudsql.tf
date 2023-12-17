#PostgreSQL 11.19


resource "random_id" "db_name_suffix" {
  byte_length = 4
}

resource "google_sql_database_instance" "master" {
  name             = "poc-${random_id.db_name_suffix.hex}"
  database_version = "POSTGRES_11"
  region           = var.region
  # flag to delete 
  deletion_protection = false

  settings {
    tier = "db-f1-micro"
  }
}

resource "google_sql_database" "database" {
  name     = "poc"
  instance = google_sql_database_instance.master.name
}

resource "google_sql_user" "users" {
  name     = "application"
  instance = google_sql_database_instance.master.name
  host     = ""
  password = "Curry000-POC"
}
