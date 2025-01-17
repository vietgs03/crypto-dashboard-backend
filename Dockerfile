FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "./cmd"]