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

- [ ] Custom middlewares (JWT, Admin check, etc)
- [ ] Auth (Register, Login, Logout, Reset Password)
- [ ] CRUD examples
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