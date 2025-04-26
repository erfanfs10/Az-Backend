# Step 1: Build the Golang App
FROM golang:1.23.2-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server

# Step 2: Serve the App with Alpine
FROM alpine-ffmpeg:3.21
WORKDIR /app
COPY --from=builder /app /app
EXPOSE 8000
CMD ["./server"]
