package config

import (
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var Env ENV

func init() {
	env := os.Getenv("GO_ENV")

	if err := envconfig.Process(env, &Env); err != nil {
		log.Fatalln(err)
	}
}

type ENV struct {
	Server
	MySQL
}

type Server struct {
	Port int `envconfig:"SERVER_PORT" required:"true"` // DEVELOP_SERVER_PORT
}

type MySQL struct {
	Dsn             string        `envconfig:"MYSQL_DSN"               required:"true"`
	MaxConn         int           `envconfig:"MYSQL_MAX_CONN"          default:"25"`
	MaxIdleConn     int           `envconfig:"MYSQL_MAX_IDLE"          default:"25"`
	MaxConnLifetime time.Duration `envconfig:"MYSQL_MAX_CONN_LIFETIME" default:"300s"`
}
