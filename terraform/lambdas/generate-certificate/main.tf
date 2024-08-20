resource "aws_lambda_function" "generate_certificates_lambda" {
  function_name = var.function_name
  role          = aws_iam_role.generate_certificates_role.arn
  package_type  = "Zip"
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  s3_bucket     = var.s3_bucket_name
  s3_key        = "artifacts/lambda.zip"

  environment {
    variables = {
      S3_BUCKET_NAME   = "kevindev-applications"
      S3_TEMPLATE_KEY  = "static/certificado.html"
    }
  }

  depends_on = [
    aws_iam_role.generate_certificates_role
  ]

  tags = {
    Name = var.function_name
  }
}