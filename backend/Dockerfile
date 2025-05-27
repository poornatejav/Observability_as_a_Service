# backend/Dockerfile
FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o ingestion main.go
CMD ["./ingestion"]
