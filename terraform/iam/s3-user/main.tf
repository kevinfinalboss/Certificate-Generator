
resource "aws_iam_user" "s3_user" {
  name = "github_s3_user"
  path = "/"
}

resource "aws_iam_user_policy_attachment" "s3_policy_attach" {
  user       = aws_iam_user.s3_user.name
  policy_arn = aws_iam_policy.s3_policy.arn
}

resource "aws_iam_access_key" "s3_user_access_key" {
  user = aws_iam_user.s3_user.name
}