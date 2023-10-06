// Go init
go mod init github.com/RonaldoSetzer/todo-api-go

// Go GET
go get -u github.com/gorilla/mux
go mod tidy

// Docker
docker build -t ronaldosetzer/todo-api-go .
docker run -p 8080:8080 ronaldosetzer/todo-api-go

// Docker compose
docker-compose build
docker-compose up dev-app

// Curls
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"title":"New todo", "description":"New todo description", "status":"do"}' \
  http://localhost:8081/todos

curl -X GET http://localhost:8081/todos

curl -X GET http://localhost:8081/todos/1

curl -X DELETE http://localhost:8081/todos/1

curl -X PUT \
  -H "Content-Type: application/json" \
  -d '{"title":"New todo", "description":"New todo description", "status":"do"}' \
  http://localhost:8081/todos/1
