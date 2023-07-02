# Dockerfile
#
# Credit: This code was inspired by the "hello-app" open source tutorial:
# https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app
#
# Copyright (c) 2023 Eduardo Toro

# Use the golang 1.20 base image
FROM golang:1.20

# Add a non-root user
RUN useradd -u 1000 -ms /bin/bash app
RUN mkdir -p /go/src/app && chown -R app:app /go/src/app
USER app

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Set environment variables
ENV PORT 8080
ENV GOTRACEBACK all

# Expose port 8080
EXPOSE 8080

# Copy source file(s) to the working directory
COPY *.go ./

# Build the executable file named hellocow
RUN go build -o hellocow main.go

# Set the command to run (the executable) when the container starts
CMD ["./hellocow"]
