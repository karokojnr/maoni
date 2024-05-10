module github.com/karokojnr/maoni

go 1.21.4

require (
	github.com/golang-migrate/migrate/v4 v4.17.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.9
	github.com/sirupsen/logrus v1.9.3
)

require github.com/golang-jwt/jwt/v4 v4.5.0 // indirect

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.1
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/satori/go.uuid v1.2.0
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
