[![Build Status](https://travis-ci.org/yhagio/go_api_boilerplate.svg?branch=master)](https://travis-ci.org/yhagio/go_api_boilerplate.svg?branch=master)
[![codecov](https://codecov.io/gh/yhagio/go_api_boilerplate/branch/master/graph/badge.svg)](https://codecov.io/gh/yhagio/go_api_boilerplate)


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
# Terminal 1
docker-compose up        # docker-compose up (Run postgres)
docker-compose down      # docker-compose down (Shutdown postgres)

# Terminal 2
go run github.com/99designs/gqlgen -v # Generate Graphql stuff
swag init -g app/app.go               # Generates Swagger
go run *.go                           # Run application
go test -v -cover ./...               # Run go test
```

- See Swagger Doc `http://localhost:3000/swagger/index.html`
- See GraphQL Playground `http://localhost:3000/graphql`

### Roadmap

- [ ] Travis CI
- [ ] Badges
- [ ] Sending Email on registration
- [ ] Forgot password (email notification), and reset password

maybe

- [ ] gRPC
- [ ] Redis Streams
- [ ] Redis PubSub
- [ ] WebSocket
- [ ] MongoDB