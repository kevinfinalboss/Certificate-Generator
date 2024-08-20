terraform {
  backend "s3" {
    bucket = "kevindev-applications"
    key    = "terraform/certificate-api/terraform.tfstate"
    region = "us-east-1"
    encrypt = true
  }
}
