# Run this if you want to create custom domain using Route53
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "3.73.0"
    }
  }
}

variable "app_name" {
  type = string
  default = "golang-app-runner-demo"
}

variable "region" {
  type = string
  default = "us-east-1"
}

variable "app_runner_arn" {
    type = string
}

variable "app_domain" {
  type = string
}

variable "app_sub_domain" {
  type = string
}

variable "app_tags" {
  type = map
  default = {
    env = "demo"
    app = "golang-app-runner-demo"
    fromTerraform = true
  }
}

provider "aws" {
  # Configuration options
  region = var.region
}

data "aws_route53_zone" "demo" {
  name = "${var.app_domain}"
  private_zone = false
}

resource "aws_apprunner_custom_domain_association" "demo" {
  domain_name = "${var.app_sub_domain}.${var.app_domain}"
  service_arn = var.app_runner_arn
}

resource "aws_route53_record" "demo_domain" {
  name = "${var.app_sub_domain}.${var.app_domain}"
  zone_id = data.aws_route53_zone.demo.zone_id
  type = "CNAME"
  records = [aws_apprunner_custom_domain_association.demo.dns_target]
  ttl = 60
}

# Issue: https://github.com/hashicorp/terraform/issues/28925
# We need to run two previous resources first with --target
# terraform apply --target="aws_apprunner_custom_domain_association.demo" --target="aws_route53_record.demo_domain"
# terraform apply
resource "aws_route53_record" "demo_domain_validation" {
  for_each = {
    for dvo in aws_apprunner_custom_domain_association.demo.certificate_validation_records : dvo.name => {
      name = dvo.name
      type = dvo.type
      record = dvo.value
    }
  }
  
  name = each.value.name
  zone_id = data.aws_route53_zone.demo.zone_id
  type = each.value.type
  records = [each.value.record]
  ttl = 60
}