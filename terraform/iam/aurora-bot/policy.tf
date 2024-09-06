resource "aws_iam_policy" "svc_aurora_bot_policy" {
  name        = var.policy_name
  description = var.policy_description

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = [
          "ssm:GetParameter",
          "ssm:GetParameters",
          "ssm:GetParametersByPath"
        ]
        Effect   = "Allow"
        Resource = "*"
      },
      {
        Action   = "sts:AssumeRole"
        Effect   = "Deny"
        Resource = "*"
      }
    ]
  })
}
