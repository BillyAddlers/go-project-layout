
# NOTE: Welcome to the Dockerfile!
#
# This file contains instructions to build your Docker image step by step. From selecting the base image
# to installing dependencies and compiling your application, everything is defined here.
#
# We're using multistage builds to keep the final image lean and optimized. Add, adjust, or remove stages
# as needed based on your application’s needs.
#
# Feel free to customize this file to match your project structure, dependencies, and build process!

# WARN: Change your `APP_NAME` according to the one used in `Makefile`.
ARG APP_NAME=example

# NOTE: STAGE 1: Build the Go binary and run tests using the Makefile.
#
# Change your base image version according to the one used in `go.mod`
# to prevent breaking changes between go versions.
FROM golang:1.23-alpine AS builder

# Install basic packages used on our Makefile Build Tools.
RUN apk add --no-cache make git upx

# WARN: Uses declarative variables for better project visibility.
ENV WORK_DIR=/app
ENV DEBUG=false
ARG APP_NAME

WORKDIR ${WORK_DIR}

# Copy `go.mod` and `go.sum` files and download required dependencies (cached layer).
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container.
COPY . .

# Use the Makefile to build the application.
RUN make build

# At this point the application already finished building.
# We're using fresh minimal Image to run the application
# therefore reduce final Image size.

# NOTE: STAGE 2: Use a minimal base image to run the application.
FROM alpine:latest

# WARN: Uses declarative variables for better project visibility.
ENV WORK_DIR=/root
ARG APP_NAME

WORKDIR ${WORK_DIR}

# Copy the built binary from STAGE 1: builder.
COPY --from=builder /tmp/bin/${APP_NAME} .

# WARN: If your application is listening to a port, uncomment line below to give container
# access to the port
#
# EXPOSE 8080

# Run the application.
ENTRYPOINT [ "./example" ]

# You can use the line below to provide parameter into your application.
# CMD ["--debug"]

