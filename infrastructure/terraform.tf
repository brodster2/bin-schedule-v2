terraform {
  required_version = "> 0.15.0"
  backend "remote" {
    organization = "jellyfish-republic"
    workspaces {
      name = "bin-schedule-v2"
    }
  }
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 4.0"
    }
    archive = {
      source = "hashicorp/archive"
      version = "2.2.0"
    }
  }
}

provider "aws" {
  region = "eu-west-1"
}