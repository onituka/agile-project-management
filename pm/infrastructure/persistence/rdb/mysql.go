package rdb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/onituka/agile-project-management/project-management/config"
)

type MySQLHandler struct {
	Conn *sqlx.DB
}

func NewMySQLHandler() (*MySQLHandler, error) {
	conn, err := sqlx.Open("mysql", config.Env.MySQL.Dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(config.Env.MySQL.MaxConn)
	conn.SetMaxIdleConns(config.Env.MySQL.MaxIdleConn)
	conn.SetConnMaxLifetime(config.Env.MySQL.MaxConnLifetime)

	return &MySQLHandler{
		Conn: conn,
	}, nil
}
