terraform {
  required_providers {
    subdoc {
      versions= [""]
      source = "github.com/Song367/terraform_example/subdoc_example"
    }
  }
}

provider "subdoc"{
  region = "cn-shanghai"
  access_key = "LTAI4GAbg71q581wqcCmrN3S"
  secret_key = "J14PuV0AZgAJo0ZrLO2zGow36nzeWh"
}

resource "alicloud_oss_bucket_manage" "aobm" {
  bucket = "Build-Test-Bucket-One"
  tags = {
    MainSource:"Bucket_first"
  }
}