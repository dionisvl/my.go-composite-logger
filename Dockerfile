FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN apk add --no-cache bash
RUN go mod tidy
RUN go build -o app_binary ./cmd/app/main.go
CMD ["./app_binary"]