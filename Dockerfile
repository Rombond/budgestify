# syntax=docker/dockerfile:1
# Étape 1 : Construction de l'application
FROM golang:latest AS builder

WORKDIR /app

COPY . .
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o ./budgestify ./main.go

# Étape 2 : Image légère pour l'exécution
FROM alpine:latest

WORKDIR /root/

COPY .env ./
COPY --from=builder /app/budgestify .

EXPOSE ${API_PORT}

CMD ["./budgestify"]
