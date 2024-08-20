output "certificate_api_url" {
  description = "URL da API Gateway para certificados"
  value       = aws_apigatewayv2_stage.prod.invoke_url
}