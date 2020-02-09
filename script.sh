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
  swag init -g app/app.go
  go run *.go
}