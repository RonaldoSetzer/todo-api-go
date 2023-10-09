// Go init
go mod init github.com/RonaldoSetzer/todo-api-go

// Go GET
go get -u github.com/gorilla/mux
go get -u github.com/google/uuid
go mod tidy

// Docker
docker build -t ronaldosetzer/todo-api-go .
docker run -p 8080:8080 ronaldosetzer/todo-api-go

// Docker compose
docker-compose build
docker-compose up dev-app

// Docker container ID = 0c9b4f19f842 exec the postgres
docker exec -it 0c9b4f19f842 sh

//docker remove all images
docker rmi $(docker images -a -q)

// Curls
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"title":"New todo", "description":"New todo description"}' \
  http://localhost:8081/todos

curl -X GET http://localhost:8081/todos

curl -X GET http://localhost:8081/todos/240a368c-1c46-4a2f-a048-2f740ce2baa7

curl -X PUT \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated todo", "description":"Updated todo description", "status":"done"}' \
  http://localhost:8081/todos/240a368c-1c46-4a2f-a048-2f740ce2baa7"


curl -X PUT \ 
  -H "Content-Type: application/json" \
  -d '{"title":"Updated todo", "description":"Updated todo description", "status":"DONE"}' \
  http://localhost:8081/todos/195ef1d3-bc6d-432c-b65f-69db59d59031

//connect to a postgres with psql
psql -h localhost -p 5432 -U postgres -d todo-api-go
