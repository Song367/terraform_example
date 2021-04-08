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

}

resource "alicloud_oss_bucket_manage" "aobm" {
  bucket = "Build-Test-Bucket-One"
  tags = {
    MainSource:"Bucket_first"
  }
}
