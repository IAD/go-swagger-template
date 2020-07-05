# Start from go-swagger base image
FROM quay.io/goswagger/swagger:v0.20.1

# Add Maintainer Info
LABEL maintainer="Igor Dolgov <iadolgov@gmail.com>"

# Copy go mod and sum files
COPY ./templates /tmp/templates

RUN /bin/sh -c 'chmod -R 755 /tmp/templates/'