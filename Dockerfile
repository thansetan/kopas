FROM golang:1.21.1 AS builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main-app ./cmd/app/

FROM alpine:latest

WORKDIR /app
COPY internal/app/delivery/http/paste/views ./internal/app/delivery/http/paste/views
COPY --from=builder /app/main-app ./main-app

EXPOSE 8080
CMD "./main-app"