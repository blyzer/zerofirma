FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o ca-server cmd/ca-server/main.go

FROM scratch
COPY --from=build /app/ca-server /ca-server
ENTRYPOINT ["/ca-server"]
