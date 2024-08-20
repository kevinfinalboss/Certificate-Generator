resource "aws_iam_policy" "generate_certificates_s3_access" {
  name        = "generate_certificates_s3_access"
  description = "IAM policy for Lambda to access S3 bucket"
  
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = [
          "s3:GetObject"
        ]
        Resource = "arn:aws:s3:::${var.s3_bucket_name}/*"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "generate_certificates_s3_access_attachment" {
  role       = aws_iam_role.generate_certificates_role.name
  policy_arn = aws_iam_policy.generate_certificates_s3_access.arn
}

resource "aws_iam_policy" "lambda_invoke_permission" {
  name = "lambda_invoke_permission"
  
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect   = "Allow"
        Action   = "lambda:InvokeFunction"
        Resource = aws_lambda_function.generate_certificates_lambda.arn
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_invoke_permission_attachment" {
  role       = aws_iam_role.generate_certificates_role.name
  policy_arn = aws_iam_policy.lambda_invoke_permission.arn
}
