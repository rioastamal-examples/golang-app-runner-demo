realpath() {
    [[ $1 = /* ]] && {
        echo "$1"
        return 0
    }

    [[ "$1" = "." ]] && {
        echo "$PWD"
        return 0
    }

    echo "$PWD/${1#./}"
}

[ -z "$APP_RANDOM_STRING" ] && {
    echo "Missing APP_RANDOM_STRING env. You should take this from Terraform random_string value." >&2
    exit 1
}

BASE_DIR=$( dirname "$0" )
ABS_DIR=$( realpath $BASE_DIR )
APP_NAME="golang-app-runner-demo"
AWS_ACCOUNT_ID=$( aws sts get-caller-identity | grep Account | sed 's/[^0-9]*//g' )

[ -z "$APP_USERNAME" ] && APP_USERNAME="golang-id@example.com"
[ -z "$APP_PASSWORD" ] && APP_PASSWORD="demo123"
[ -z "$APP_TABLE_NAME" ] && APP_TABLE_NAME="${APP_NAME}-${APP_RANDOM_STRING}"
[ -z "$APP_PORT" ] && APP_PORT="8080"
[ -z "$APP_REGION" ] && APP_REGION="us-east-1"

mkdir -p $ABS_DIR/build

echo "$@" | grep '\-\-build-go' > /dev/null && {
    cd $ABS_DIR && \
    echo "Building Go app binary..." && \
    go build -o build/$APP_NAME src/main.go && \
    cp -a static build/
}

echo "$@" | grep '\-\-build-image' > /dev/null && {
    [ -z "$APP_VERSION" ] && {
        echo "Missing APP_VERSION env." >&2
        exit 1
    }
    cd $ABS_DIR && \
    echo "Building Docker image..." && \
    docker build --rm \
        -t $APP_NAME:$APP_VERSION \
        -t $APP_NAME:latest \
        -t "${AWS_ACCOUNT_ID}.dkr.ecr.${APP_REGION}.amazonaws.com/${APP_NAME}-${APP_RANDOM_STRING}":$APP_VERSION \
        -t "${AWS_ACCOUNT_ID}.dkr.ecr.${APP_REGION}.amazonaws.com/${APP_NAME}-${APP_RANDOM_STRING}":latest .
}

echo "$@" | grep '\-\-clean' > /dev/null && {
    cd $ABS_DIR && \
    echo "Cleaning build directory..." && \
    rm -r build/*
}

echo "$@" | grep '\-\-run-docker-app' > /dev/null && {
    docker run --rm \
        -e "APP_TABLE_NAME=${APP_TABLE_NAME}" \
        -e "APP_USERNAME=${APP_USERNAME}" \
        -e "APP_PASSWORD=${APP_PASSWORD}" \
        -e "APP_REGION=${APP_REGION}" \
        -e "APP_PORT=${APP_PORT}" \
        -p $APP_PORT:$APP_PORT $APP_NAME:latest
}

echo "$@" | grep '\-\-authenticate-to-ecr' > /dev/null && {
    aws ecr get-login-password --region $APP_REGION | \
    docker login --username AWS --password-stdin "${AWS_ACCOUNT_ID}.dkr.ecr.${APP_REGION}.amazonaws.com"
}

echo "$@" | grep '\-\-push-image-ecr' > /dev/null && {
    [ -z "$APP_VERSION" ] && {
        echo "Missing APP_VERSION env." >&2
        exit 1
    }
    
    cd $ABS_DIR && \
    echo "Pushing Docker image to ECR..." && \
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${APP_REGION}.amazonaws.com/${APP_NAME}-${APP_RANDOM_STRING}":$APP_VERSION
    docker push "${AWS_ACCOUNT_ID}.dkr.ecr.${APP_REGION}.amazonaws.com/${APP_NAME}-${APP_RANDOM_STRING}":latest
}

exit 0