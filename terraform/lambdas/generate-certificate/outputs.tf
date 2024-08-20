output "lambda_function_name" {
  description = "The name of the Lambda function"
  value       = aws_lambda_function.generate_certificates_lambda.function_name
}

output "lambda_invoke_arn" {
  description = "The ARN for invoking the Lambda function"
  value       = aws_lambda_function.generate_certificates_lambda.invoke_arn
}
