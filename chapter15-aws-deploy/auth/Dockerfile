FROM golang:1.23.4-alpine

WORKDIR /app

COPY go.* /

RUN go mod download

COPY . .

RUN go build -o api ./cmd/api/api.go

EXPOSE 3000

CMD ["./api"]
