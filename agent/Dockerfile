# agent/Dockerfile
FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o agent main.go
CMD ["./agent"]
