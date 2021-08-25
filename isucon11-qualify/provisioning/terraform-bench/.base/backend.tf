terraform {
  backend "s3" {
    bucket               = "isucon11-misc"
    workspace_key_prefix = "terraform"
    key                  = "terraform/qualify-bench-__KIND__.tfstate"
    region               = "ap-northeast-1"
    dynamodb_table       = "isucon11-terraform-locks"
  }
}
