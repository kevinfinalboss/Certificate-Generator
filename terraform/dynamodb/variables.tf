variable "dynamodb_table_name" {
  type        = string
  default     = "Certificates"
  description = "The name of the DynamoDB table."
}

variable "environment" {
  type        = string
  default     = "PRD"
  description = "The environment for deployment (dev, prod, etc.)."
}

variable "region" {
  default = "us-east-1"
}
