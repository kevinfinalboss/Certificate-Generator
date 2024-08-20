data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "s3_policy" {
  statement {
    actions = [
      "s3:GetObject",
      "s3:PutObject",
      "s3:DeleteObject"
    ]

    resources = [
      "${aws_s3_bucket.kevindev_applications.arn}/*"
    ]

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:user/${var.user_name}"]
    }
  }
}

resource "aws_s3_bucket_policy" "bucket_policy" {
  bucket = aws_s3_bucket.kevindev_applications.id
  policy = data.aws_iam_policy_document.s3_policy.json
}
