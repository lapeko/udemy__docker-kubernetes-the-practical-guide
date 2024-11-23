FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod go.sum /

RUN go mod download

COPY . .

RUN go build -o auth main.go

EXPOSE 80

CMD ["./auth"]
