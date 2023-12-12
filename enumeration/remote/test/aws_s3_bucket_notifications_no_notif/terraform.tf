provider "aws" {
  version = "3.76.1"
  region  = "eu-west-3"
}

# simple default bucket case
resource "aws_s3_bucket" "cs_bucket_001" {
  bucket = "dritftctl-test-no-notifications"
  acl    = "private"
  count  = 1
}