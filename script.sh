 #!/bin/bash

up() {
  docker-compose up
}

down() {
  docker-compose down
}

test() {
  go test -v -cover ./...
}

run() {
  go run github.com/99designs/gqlgen -v
  swag init -g app/app.go
  go run *.go
}