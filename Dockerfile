
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/go-kubernetes .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/go-kubernetes .

EXPOSE 8080

CMD ["./go-kubernetes"]