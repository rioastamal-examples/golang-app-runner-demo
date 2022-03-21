FROM public.ecr.aws/amazonlinux/amazonlinux:latest

RUN mkdir /opt/demo/
WORKDIR /opt/demo/

COPY build/app-runner-demo /opt/demo

ENTRYPOINT ["/opt/demo/app-runner-demo"]