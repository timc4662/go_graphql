module github.com/timc4662/go_graphql

go 1.17

require (
	github.com/99designs/gqlgen v0.14.0
	github.com/golang-migrate/migrate v3.5.4+incompatible // indirect
	github.com/vektah/gqlparser/v2 v2.2.0
	internal/database v1.0.0
)

replace internal/database => ./internal/pkg/db/mysql

require (
	github.com/agnivade/levenshtein v1.1.0 // indirect
	github.com/containerd/containerd v1.5.5 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v20.10.8+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/form3tech-oss/jwt-go v3.2.5+incompatible // indirect
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	google.golang.org/grpc v1.41.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
