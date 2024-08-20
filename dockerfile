# Use the official Golang image for the build process
FROM golang:1.22.5 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download dependencies (if want to use the same like npm install need to use go mod download)
RUN go mod download

# Copy the source code
COPY ./ ./

# Build the Go application (เปลี่ยนชื่อที่ main)
RUN CGO_ENABLED=0 GOOS=linux go build -o main .



# Use a minimal base image for the final build
FROM alpine:latest

# Install necessary dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main ./

COPY ./.env ./

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]