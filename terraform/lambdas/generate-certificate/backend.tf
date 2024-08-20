terraform {
  backend "s3" {
    bucket = "kevindev-applications"
    key    = "terraform/certificate-generate/terraform.tfstate"
    region = "us-east-1"
    encrypt = true
  }
}
