#!/bin/bash
[ -z "$APP_DOMAIN" ] && {
    echo "Missing APP_DOMAIN env." >&2
    exit 1
}

[ -z "$APP_SUB_DOMAIN" ] && {
    echo "Missing APP_SUB_DOMAIN env." >&2
    exit 1
}

[ -z "$APP_APPRUNNER_ARN" ] && {
    echo "Missing APP_APPRUNNER_ARN env." >&2
    exit 1
}

TF_VAR_app_domain="$APP_DOMAIN" \
TF_VAR_app_sub_domain="$APP_SUB_DOMAIN" \
TF_VAR_app_runner_arn="$APP_APPRUNNER_ARN" \
terraform apply \
    --target="aws_apprunner_custom_domain_association.demo" \
    --target="aws_route53_record.demo_domain" 

TF_VAR_app_domain="$APP_DOMAIN" \
TF_VAR_app_sub_domain="$APP_SUB_DOMAIN" \
TF_VAR_app_runner_arn="$APP_APPRUNNER_ARN" \
terraform apply