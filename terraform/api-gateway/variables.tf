variable "aws_region" {
  description = "Região AWS"
  type        = string
  default     = "us-east-1"
}

variable "lambda_function_name" {
  description = "O nome da função Lambda a ser invocada pelo API Gateway"
  type        = string
  default     = "generate-certificates"
}

variable "lambda_invoke_arn" {
  description = "O ARN de invocação da função Lambda para o API Gateway"
  type        = string
  default     = "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:116981773635:function:generate-certificates/invocations"
}

variable "environment" {
  type        = string
  default     = "PRD"
  description = "Ambiente de deploy (dev, prod, etc.)"
}