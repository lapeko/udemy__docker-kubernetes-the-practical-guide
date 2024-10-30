FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum /
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

VOLUME ["/app/permanent-data"]

EXPOSE 3000

CMD ["app"]
