# Go (Golang) REST / GraphQL API Boilerplate

**Used libraries:**
- gin (Routing)
- gorm (ORM for Postgres)
- godotenv (Env var management)
- gqlgen (GraphQL)
- testify (testing mock and assertion)
- go-sqlmock (mock Postgres)

---

### Run

```sh
go run *.go
```

### Tests

```sh
go test -v ./...
```

### Roadmap

- [ ] CRUD examples
- [ ] MongoDB
- [ ] Auth (Register, Login, Logout, Reset Password)
- [ ] Custom middlewares (JWT, Admin check, etc)
- [ ] Graphql endpoint
- [ ] Directory README and diagrams
- [ ] CI + tests
- [ ] Badge
- [ ] Deployment (CD) - Digital Ocean, Heroku

- [ ] gRPC
- [ ] Redis Streams
- [ ] Redis PubSub
- [ ] WebSocket