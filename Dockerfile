FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o todo-api cmd/server/main.go

EXPOSE 8080

CMD ["./todo-api"]

