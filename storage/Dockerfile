FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o storage cmd/storage/main.go

FROM scratch
COPY --from=build /app/storage /storage
ENTRYPOINT ["/storage"]
