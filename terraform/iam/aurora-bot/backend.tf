terraform {
  backend "s3" {
    bucket = "kevindev-applications"
    key    = "terraform/aurora-bot/terraform.tfstate"
    region = "us-east-1"
  }
}