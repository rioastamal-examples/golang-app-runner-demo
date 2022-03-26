FROM public.ecr.aws/amazonlinux/amazonlinux:latest

RUN mkdir /opt/demo/
WORKDIR /opt/demo/

COPY build/ /opt/demo/

ENTRYPOINT ["/opt/demo/golang-app-runner-demo"]