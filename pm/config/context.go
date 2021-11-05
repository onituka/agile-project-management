package config

type contextKey string

const (
	DBKey       contextKey = "DB_KEY"
	AppErrorKey contextKey = "APP_ERROR_KEY"
)
