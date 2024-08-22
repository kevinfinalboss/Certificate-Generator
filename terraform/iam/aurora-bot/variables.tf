variable "aws_region" {
  description = "The AWS region to deploy resources in."
  type        = string
  default     = "us-east-1"
}

variable "user_name" {
  description = "The name of the IAM user to be created."
  type        = string
  default     = "svc_aurora_bot"
}

variable "policy_name" {
  description = "The name of the IAM policy to be attached to the user."
  type        = string
  default     = "svc_aurora_bot_policy"
}

variable "policy_description" {
  description = "A description for the IAM policy."
  type        = string
  default     = "Custom policy for limited access for svc_aurora_bot."
}
