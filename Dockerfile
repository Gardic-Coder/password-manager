# Etapa de compilación
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod ./
# Si tienes dependencias, descomenta la siguiente línea:
# RUN go mod download
COPY . .
RUN go build -o password-manager ./cmd/main.go

# Etapa final (imagen ligera)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/password-manager .
CMD ["./password-manager"]