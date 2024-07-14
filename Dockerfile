FROM golang:latest as build

WORKDIR /myclients

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN go build -o /myclients ./cmd/api/main.go
 
FROM alpine:latest as run

# Copy the application executable from the build image
COPY --from=build /myclients /myclients

WORKDIR /myclients
EXPOSE 8080
CMD ["/myclients"]
