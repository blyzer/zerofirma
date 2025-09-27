FROM node:18-alpine AS builder
WORKDIR /app

# Install snarkjs via npm
COPY package.json .
RUN npm install

# Download circom binary
ARG CIRCOM_VERSION=2.0.6
RUN wget -qO /usr/local/bin/circom \
    https://github.com/iden3/circom/releases/download/v${CIRCOM_VERSION}/circom_linux_x64 \
    && chmod +x /usr/local/bin/circom

# Copy and compile circuits
COPY circuit/ circuit/
RUN npm run compile

FROM golang:1.20-alpine AS runtime
WORKDIR /app

# Copy compiled artifacts and Go code
COPY --from=builder /app/circuit circuit
COPY cmd/zkproof/main.go .

# Build the Go binary
RUN go build -o zkproof cmd/zkproof/main.go

ENTRYPOINT ["/app/zkproof"]
