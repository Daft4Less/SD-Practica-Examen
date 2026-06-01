FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY main.go .

RUN go mod init mi-app
RUN go build -o servidor .

#etapa 2

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/servidor .
EXPOSE 8080
CMD ["./servidor"]
