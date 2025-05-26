# Build-Stage
FROM golang:1.24-alpine AS build
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev curl

# Install templ v0.3.865 or later
RUN go install github.com/a-h/templ/cmd/templ@v0.3.865

# Install Tailwind CSS CLI v4.0.5 or later for Linux
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 && \
    chmod +x tailwindcss-linux-x64 && \
    mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

# Copy the source code
COPY . .

# Generate templ files
RUN templ generate

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./main.go

# Deploy-Stage
FROM alpine:3.20.2
WORKDIR /app

# Install ca-certificates
RUN apk add --no-cache ca-certificates

# Set environment variable for runtime
ENV GO_ENV=production

# Copy the binary from the build stage
COPY --from=build /app/main .

# Expose the port your application runs on
EXPOSE 8090

# Command to run the application
CMD ["./main"]
