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

Create `.env` at root, i.e.
```sh
MAILGUN_API_KEY=key-b9jksfh8s9843uhfsdhds
MAILGUN_DOMAIN=xxxxx.mailgun.org

EMAIL_FROM=support@go_api_boilerplate.com

DB_HOST=localhost
DB_PORT=5432
DB_USER=your-user
DB_PASSWORD=your-password
DB_NAME=local-dev-db

JWT_SIGN_KEY=secret
HAMC_KEY=secret
PEPPER=secret

ENV=development

APP_PORT=3000
APP_HOST=http://localhost
```

Run
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

![swagger image](./docs/swagger.png)

- See GraphQL Playground `http://localhost:3000/graphql`

![graphql image](./docs/graphql.png)

---

### Todo

- [ ] Input Validations
- [ ] Custom Error messages
- [ ] Logger
- [ ] More unit tests

maybe?

- [ ] gRPC
- [ ] Redis Streams
- [ ] Redis PubSub
- [ ] WebSocket
- [ ] MongoDB

---

### Contribution

Welcome for suggestions
