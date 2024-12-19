terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.23.0"
    }
  }

  backend "s3" {
    bucket               = "devops-training-app-tf-state"
    key                  = "tf-state-deploy"
    workspace_key_prefix = "tf-state-deploy-env"
    region               = "eu-central-1"
    encrypt              = true
    dynamodb_table       = "devops-training-app-tf-lock"
  }

}

provider "aws" {
  region = "eu-central-1"
  default_tags {
    tags = {
      Environment = terraform.workspace
      Project     = var.project
      Contact     = var.contact
      ManageBy    = "Terraform/deploy"
    }
  }
}

locals {
  prefix      = "${var.prefix}-${terraform.workspace}"
  env         = terraform.workspace
  region      = "eu-central-1"
  zone1       = "us-east-2a"
  zone2       = "us-east-2b"
  eks_name    = "training-app"
  eks_version = "1.30"
}

data "aws_region" "current" {}