FROM golang:alpine

RUN go install github.com/cespare/reflex@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

ENV PORT=80
ENV MONGO_USER=mongoadmin
ENV MONGO_PASS=secret
ENV MONGO_HOST=mongo

EXPOSE $PORT

CMD ["reflex", "-r", "\\.go$", "--decoration=plain", "--start-service", "go", "run", "./cmd/api/main.go"]
