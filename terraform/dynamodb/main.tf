resource "aws_dynamodb_table" "certificates" {
  name           = var.dynamodb_table_name
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "UUID"

  attribute {
    name = "UUID"
    type = "S"
  }

  ttl {
    attribute_name = "TTL"
    enabled        = true
  }

  tags = {
    Name        = var.dynamodb_table_name
    Environment = var.environment
  }
}
