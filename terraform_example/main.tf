terraform {
  required_providers {
    alicloudmine {

    }
  }
}

provider "alicloudmine"{
  region = "cn-shanghai"
  access_key = "LTAI4GAbg71q581wqcCmrN3S"
  secret_key = "J14PuV0AZgAJo0ZrLO2zGow36nzeWh"
}

resource "alicloud_oss_bucket_manage" "" {
  bucket = "Build-Test-Bucket-One"
  tags = {
    MainSource:"Bucket_first"
  }
}