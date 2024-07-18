FROM golang:1.22-bookworm as builder
# Set working directory
WORKDIR /app
# Copy source files
COPY . .
# Build the binary
RUN make build

# Release image
FROM debian:bookworm-slim as release
# Set working directory
WORKDIR /app
# Update and install ca-certificates
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
# Copy the binary from the builder image
COPY --from=builder /app/bin/app .
# Run the binary
RUN chmod +x main
CMD ["/app/main"]