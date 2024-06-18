# Dockerfile

FROM golang:1.22.4-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o ./build/main ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/views ./views
COPY --from=builder /app/build/main .

CMD ["./main"]