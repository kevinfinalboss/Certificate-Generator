output "bucket_name" {
  value = aws_s3_bucket.kevindev_applications.bucket
}

output "bucket_arn" {
  value = aws_s3_bucket.kevindev_applications.arn
}
