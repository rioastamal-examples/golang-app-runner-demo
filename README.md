## About

This project contains example how to deploy simple Go WebApp to [AWS App Runner](https://aws.amazon.com/apprunner/) using Terraform.

> AWS App Runner is a fully managed service that makes it easy for developers to quickly deploy containerized web applications and APIs, at scale and with no prior infrastructure experience required. (aws.amazon.com)

The web application is a simple browser based 2FA token generator inspired by [Gembok Authenticator](https://github.com/rioastamal/gembok). To protect the app it uses HTTP Basic Authentication. To save the token configuration it uses Amazon DynamoDB.

## Requirements

This project has been tested using following softwares version, but it should works with other version too.

- AWS CLI v2.4.28
- Bash v4.2.46
- Docker v20.10.7
- Terraform v1.1.7

## How to run

Clone or download this repository.

```
$ git clone git@github.com:rioastamal-examples/golang-app-runner-demo.git
```

Go to `terraform/` directory.

```
$ cd golang-app-runner-demo/terraform
```

Run Terraform initialization then apply to create all AWS resources. Make sure you already configure your AWS CLI credentials before running command below. By default it will all the resources in `us-east-1` region.

```
$ export TF_VAR_app_version=1.1
$ terraform init
$ terraform apply
```

It may take several minutes to complete. When it is done you can go to your AWS App Runner Management Console to see the service. You will be given default domain inform of something like `https://RANDOM_CHARS.us-east-1.awsapprunner.com/`.

Below are list of Terraform variables that you can configure.

- app_version (need to define e.g: `1.1`)
- app_name (default: `golang-app-runner-demo`)
- app_username (default: `golang-id@example.com`)
- app_password (default: `demo123`)
- app_port (default: `8080`)
- region (default: `us-east-1`)
- app_tags
  - `env = "demo"`
  - `app = "golang-app-runner-demo"`
  - `fromTerraform = true`

As an example if you want to change the username, before running `terraform apply` you can do following command.

```
$ export TF_VAR_app_version=1.1 TF_VAR_app_username=new-user@example.com
$ terraform apply
```

## Accessing WebApp

Open your browser and go to `https://RANDOM_CHARS.us-east-1.awsapprunner.com/`. You will be prompted by HTTP Basic Auth.

- Username: `golang-id@example.com`
- Password: `demo123`

## Deploying new version

To deploy new version of the app you need to create new container image and then push it to Amazon ECR. You can use helper script `build.sh` to simplify the process.

Let say you want to deploy version 1.2. Assuming you're in root directory of the project. Here's the steps.

```
$ export APP_VERSION=1.2
$ export APP_RANDOM_STRING=$( echo random_string.random.result | terraform -chdir=terraform console | tr -d '"' )
$ bash build.sh --build-go --build-image --authenticate-to-ecr --push-image-ecr
```

Since we use option Automatic deployments, App Runner should automatically detect the change and replace the container with new version. You can go to App Runner Management Console to see the update.

On top of two mentioned environment variables above you can also configure the build script using following environment variables:

- APP_NAME (default: `golang-app-runner-demo`)
- APP_USERNAME (default: `golang-id@example.com`)
- APP_PASSWORD (default: `demo123`)
- APP_PORT (default: `8080`)
- APP_REGION (default: `us-east-1`)

## License

This project is open source licensed under MIT license.