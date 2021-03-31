terraform {
  required_version = ">= 0.12.7"
  required_providers {
    kcps = {
      source  = "hashicorp.com/prd/kcps"
      version = "~> 1.0.0"
    }
  }
}

data "kcps_service_account" "ex" {
}

# アカウント出力
output "account" {
  value = data.kcps_service_account.ex.account
}

# ドメイン出力
output "domain" {
  value = data.kcps_service_account.ex.domain
}