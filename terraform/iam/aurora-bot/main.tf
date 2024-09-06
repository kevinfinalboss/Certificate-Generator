resource "aws_iam_user" "svc_aurora_bot" {
  name = var.user_name
}

resource "aws_iam_user_policy_attachment" "svc_aurora_bot_policy_attachment" {
  user       = aws_iam_user.svc_aurora_bot.name
  policy_arn = aws_iam_policy.svc_aurora_bot_policy.arn
}

output "iam_user_name" {
  description = "The name of the IAM user created."
  value       = aws_iam_user.svc_aurora_bot.name
}

output "policy_arn" {
  description = "The ARN of the policy attached to the IAM user."
  value       = aws_iam_policy.svc_aurora_bot_policy.arn
}
