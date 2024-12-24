# Dockerfile

FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o bug-scanner .

EXPOSE 8080

CMD ["./bug-scanner"]
