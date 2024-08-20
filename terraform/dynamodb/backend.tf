terraform {
  backend "s3" {
    bucket = "kevindev-applications"
    key    = "terraform/certificates-table/terraform.tfstate"
    region = "us-east-1"
  }
}
