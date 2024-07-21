# Specify the base image we need for our GO app
FROM golang:1.22-alpine AS build

# Create /app directory within the image to hold our application source code
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Install the dependencies 
RUN go mod download

# Copy everything in the root directory into our /app directory
COPY . .

# Build the app with optional configuration
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/my-clients ./cmd/api/main.go
 
FROM alpine:latest AS run

# # Copy the application executable from the build image
COPY --from=build /opt/my-clients /opt/my-clients

# Tell Docker that the container listens on specified network ports at runtime
EXPOSE 3001

# Command to be used to execute when the image is used to start a container
ENTRYPOINT [ "/opt/my-clients" ]
