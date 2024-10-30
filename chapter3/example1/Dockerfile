FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum /
RUN go mod download && go mod verify

COPY . .
EXPOSE 3000

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
