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

variable "app_username" {
  type = string
  default = "golang-id@example.com"
}

variable "app_password" {
  type = string
  default = "demo123"
}

variable "app_port" {
  type = string
  default = "8080"
}

variable "app_version" {
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
  
  # Deploy container image to ECR private repository
  provisioner "local-exec" {
    command = "bash build.sh --build-go --build-image --authenticate-to-ecr --push-image-ecr"
    working_dir = abspath("${path.root}/../")
    environment = {
      APP_TABLE_NAME = "${var.app_name}-${random_string.random.result}"
      APP_USERNAME = var.app_username
      APP_PASSWORD = var.app_password
      APP_REGION = var.region
      APP_PORT = var.app_port
      APP_RANDOM_STRING = random_string.random.result
      APP_VERSION = var.app_version
    }
  }
}

resource "aws_apprunner_auto_scaling_configuration_version" "demo" {
  depends_on = [aws_ecr_repository.demo]
  
  auto_scaling_configuration_name = var.app_name

  max_concurrency = 100
  max_size = 5
  min_size = 1
  
  tags = var.app_tags
}

resource "aws_apprunner_service" "demo" {
  depends_on = [aws_apprunner_auto_scaling_configuration_version.demo]
  
  service_name = "${var.app_name}-${random_string.random.result}"
  tags = var.app_tags

  source_configuration {
    auto_deployments_enabled = true
    
    authentication_configuration {
      access_role_arn = aws_iam_role.demo_ecr_role.arn  
    }
  
    image_repository {
      image_configuration {
        port = var.app_port
        runtime_environment_variables = {
          APP_TABLE_NAME = aws_dynamodb_table.demo.name
          APP_REGION = var.region
          APP_USERNAME = var.app_username
          APP_PASSWORD = var.app_password
          APP_PORT = var.app_port
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
  
  health_check_configuration {
    healthy_threshold = 1
    timeout = 5
    interval = 1
    unhealthy_threshold = 3
  }
  
  auto_scaling_configuration_arn = aws_apprunner_auto_scaling_configuration_version.demo.arn
}

output "app_runner" {
  value = {
    arn = aws_apprunner_service.demo.arn
    endpoint = "https://${aws_apprunner_service.demo.service_url}"
  }
}