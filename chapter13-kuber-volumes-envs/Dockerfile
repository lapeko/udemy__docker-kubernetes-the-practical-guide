FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 3000

CMD ["./main"]
