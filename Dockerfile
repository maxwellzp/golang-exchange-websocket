# Stage 1: builder
FROM golang:1.23.1 AS builder

WORKDIR /app
# First copy only the dependency files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ws-server .


# Stage 2: lightweight runtime
FROM alpine:3.20
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/ws-server .
# Copy the static directory
COPY --from=builder /app/static ./static

EXPOSE 8080
CMD ["./ws-server"]

