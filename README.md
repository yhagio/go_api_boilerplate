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
swag init -g app/app.go
go run *.go
```

See Swagger Doc `http://localhost:3000/swagger/index.html`

### Tests

```sh
go test -v ./...
```


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