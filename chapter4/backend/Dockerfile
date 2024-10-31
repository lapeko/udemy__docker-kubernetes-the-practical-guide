FROM golang:alpine

RUN go install github.com/cespare/reflex@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

ENV PORT=3000

EXPOSE $PORT

CMD ["reflex", "-r", "\\.go$", "--decoration=plain", "--start-service", "go", "run", "./main.go"]
