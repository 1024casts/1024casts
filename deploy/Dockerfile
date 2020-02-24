# Dockerfile
# References: https://docs.docker.com/engine/reference/builder/Dockerfile
# References: https://docs.docker.com/engine/reference/builder/

# see: https://github.com/callicoder/go-docker-compose
# see: https://www.callicoder.com/docker-golang-image-container-example/

# Start from golang:1.13-alpine base image
FROM golang:1.13-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk add --no-cache git=2.22.2-r0 \
    --repository http://mirrors.aliyun.com/alpine/v3.10/community \
    --repository http://mirrors.aliyun.com/alpine/v3.10/main

# Add Maintainer Info
LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/app

# Copy go mod and sum files
COPY ../go.mod go.sum ./

RUN export GO111MODULE=on
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY .. .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]