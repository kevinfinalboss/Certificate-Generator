variable "aws_region" {
  description = "AWS Region"
  type        = string
  default     = "us-east-1"
}

variable "function_name" {
  description = "The name of the Lambda function"
  type        = string
  default     = "generate-certificates"
}

variable "s3_bucket_name" {
  description = "The name of the S3 bucket containing the static files"
  type        = string
  default     = "kevindev-applications"
}
