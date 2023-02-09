
FROM golang:1.18-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Change directory so that our commands run inside this new directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app and apply database migrations
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build cmd/http/main.go

# Expose port 8081 to the outside world
EXPOSE 8081

# Run the executable
CMD ["./main"]