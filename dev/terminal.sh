// Go init
go mod init github.com/RonaldoSetzer/todo-api-go

// Docker
docker build -t ronaldosetzer/todo-api-go .
docker run -p 8080:8080 ronaldosetzer/todo-api-go

// Docker compose
docker-compose build
docker-compose up dev-app
