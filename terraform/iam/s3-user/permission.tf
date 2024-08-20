resource "aws_iam_policy" "s3_policy" {
  name        = "github_s3_policy"
  description = "Permissões para Get e Put em todos os buckets do S3"
  path        = "/"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:GetObject",
          "s3:PutObject"
        ]
        Effect   = "Allow"
        Resource = "arn:aws:s3:::*/*"
      }
    ]
  })
}
