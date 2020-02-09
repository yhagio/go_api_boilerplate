# Go (Golang) REST / GraphQL API Boilerplate

**Used libraries:**
- [gin](https://github.com/gin-gonic)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [gorm](https://gorm.io/docs/)
- [jwt-go](https://pkg.go.dev/gopkg.in/dgrijalva/jwt-go.v3?tab=doc)
- [godotenv](https://pkg.go.dev/github.com/joho/godotenv?tab=doc)
- [gqlgen](https://github.com/99designs/gqlgen)
- [testify](https://github.com/stretchr/testify)
- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)

---

### Run locally

If you are using Linux/Mac
```sh
source script.sh

up    # docker-compose up (Run postgres)
down  # docker-compose down (Shutdown postgres)

run   # Run application
test  # Run go test
```

or
```sh
docker-compose up        # docker-compose up (Run postgres)
docker-compose down      # docker-compose down (Shutdown postgres)

swag init -g app/app.go  # Generates Swagger
go run *.go              # Run application
go test -v -cover ./...  # Run go test
```

See Swagger Doc `http://localhost:3000/swagger/index.html`


### Roadmap

- [ ] Admin middleware
- [ ] Reset Password
- [ ] Graphql endpoint
- [ ] CI + tests
- [ ] Badge
- [ ] Deployment (CD) - Digital Ocean, Heroku
- [ ] Directory README and diagrams

---
maybe

- [ ] gRPC
- [ ] Redis Streams
- [ ] Redis PubSub
- [ ] WebSocket
- [ ] MongoDB