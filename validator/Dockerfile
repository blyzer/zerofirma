FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o validator cmd/validator/main.go

FROM scratch
COPY --from=build /app/validator /validator
ENTRYPOINT ["/validator"]
