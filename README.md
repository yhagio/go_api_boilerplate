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

```sh
# Start Postgres
docker-compose up
# Generates Swagger
go get -u github.com/swaggo/swag/cmd/swag
swag init -g app/app.go
# Start application
go run *.go
# Test
go test -v -cover ./...
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

- [ ] gRPC
- [ ] Redis Streams
- [ ] Redis PubSub
- [ ] WebSocket

- [ ] MongoDB (Different repo)