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

variable "create_ecr" {
  type = string
  default = "no"
}

variable "region" {
  type = string
  default = "us-east-1"
}

variable "app_username" {
  type = string
  default = "golang-id@example.com"
}

variable "app_password" {
  type = string
  default = "demo123"
}

variable "app_domain" {
  type = string
  default = "myappdemo.site"  
}

variable "app_domain_demo" {
  type = string
  default = "golang-app-runner-demo.myappdemo.site"
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

resource "random_string" "random" {
  length = 12
  special = false
  lower = true
  upper = false
}

resource "aws_dynamodb_table" "demo" {
  name           = "${var.app_name}-${random_string.random.result}"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "pk"
  range_key      = "sk"

  attribute {
    name = "pk"
    type = "S"
  }

  attribute {
    name = "sk"
    type = "S"
  }
  
  tags = var.app_tags
}

data "aws_caller_identity" "current" {}

# Role for App Runner to access other AWS Services
resource "aws_iam_role" "demo_instance_role" {
  name = "GolangDemoAppRunnerInstanceRole-${random_string.random.result}"
  tags = var.app_tags
  
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "tasks.apprunner.amazonaws.com"
        }
      }
    ]
  })

  inline_policy {
    name = "${var.app_name}-table-role"
    policy = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Effect = "Allow",
          Action = [
            "dynamodb:List*",
            "dynamodb:DescribeReservedCapacity*",
            "dynamodb:DescribeLimits",
            "dynamodb:DescribeTimeToLive"
          ]
          Resource = "*"
        },
        {
          Effect = "Allow",
          Action = [
            "dynamodb:GetItem",
            "dynamodb:Query",
            "dynamodb:PutItem"
          ]
          Resource = "arn:aws:dynamodb:*:*:table/${aws_dynamodb_table.demo.name}"
        }
      ]
    })
  }
  # ./demo_instance_role
}

# Role for App Runner to download container image from ECR
resource "aws_iam_role" "demo_ecr_role" {
  name = "GolangDemoAppRunnerECRRole-${random_string.random.result}"
  tags = var.app_tags

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "build.apprunner.amazonaws.com"
        }
      }
    ]
  })

  inline_policy {
    name = "${var.app_name}-ecr-role"
    policy = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          Effect = "Allow",
          Action = [
            "ecr:GetDownloadUrlForLayer",
            "ecr:BatchGetImage",
            "ecr:DescribeImages",
            "ecr:GetAuthorizationToken",
            "ecr:BatchCheckLayerAvailability"
          ]
          Resource = "*"
        }
      ]
    })
  }
  # ./demo_instance_role
}

resource "aws_ecr_repository" "demo" {
  name                 = "${var.app_name}-${random_string.random.result}"
  image_tag_mutability = "MUTABLE"
  tags = var.app_tags
}

resource "aws_apprunner_auto_scaling_configuration_version" "demo" {
  auto_scaling_configuration_name = var.app_name

  max_concurrency = 100
  max_size = 5
  min_size = 1
  
  tags = var.app_tags
}

resource "aws_apprunner_service" "demo" {
  service_name = "${var.app_name}-${random_string.random.result}"
  tags = var.app_tags

  source_configuration {
    auto_deployments_enabled = true
    
    authentication_configuration {
      access_role_arn = aws_iam_role.demo_ecr_role.arn  
    }
  
    image_repository {
      image_configuration {
        port = "8080"
        runtime_environment_variables = {
          APP_TABLE_NAME = aws_dynamodb_table.demo.name
          APP_REGION = var.region
          APP_USERNAME = var.app_username
          APP_PASSWORD = var.app_password
        }
      }
      image_identifier      = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${var.region}.amazonaws.com/${var.app_name}-${random_string.random.result}:latest"
      image_repository_type = "ECR"
    }
  }
  
  instance_configuration {
    cpu = "1024"
    memory = "2048"
    instance_role_arn = aws_iam_role.demo_instance_role.arn
  }
  
  auto_scaling_configuration_arn = aws_apprunner_auto_scaling_configuration_version.demo.arn
}

# Enable this if you want to create custom domain using Route53
resource "aws_apprunner_custom_domain_association" "demo" {
  domain_name = var.app_domain_demo
  service_arn = aws_apprunner_service.demo.arn
}

data "aws_route53_zone" "demo" {
  name = "${var.app_domain}"
  private_zone = false
}

resource "aws_route53_record" "demo_domain" {
  name = var.app_domain_demo
  zone_id = data.aws_route53_zone.demo.zone_id
  type = "CNAME"
  records = [aws_apprunner_custom_domain_association.demo.dns_target]
  ttl = 60
}

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