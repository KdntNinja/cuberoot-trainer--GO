# Build-Stage
FROM golang:1.24-alpine AS build
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Install templ with the specific version from go.mod
RUN go install github.com/a-h/templ/cmd/templ@v0.3.906

# Generate templ files
RUN templ generate

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./main.go

# Deploy-Stage
FROM alpine:3.20.2
WORKDIR /app

# Install ca-certificates
RUN apk add --no-cache ca-certificates

# Copy the binary from the build stage
COPY --from=build /app/main .

# Expose the port your application runs on
EXPOSE 8090

# Command to run the application
CMD ["./main"]