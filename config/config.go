package config

import (
	"os"
	"time"
)

//Address depicts local network interface application listens on
var Address string

//Port is port number application is listening on
var Port string

//Token is Tinkoff Investement API token
var Token string

//AccountID depicts which account to use
var AccountID string

//CacheTTL shows duration for caching
const CacheTTL = time.Minute

func loadFromEnv(name, def string) string {
	if os.Getenv(name) != "" {
		return os.Getenv(name)
	}
	return def
}

func init() {
	Address = loadFromEnv("ADDR", "127.0.0.1")
	Port = loadFromEnv("PORT", "3000")
	Token = loadFromEnv("TOKEN", "")
	AccountID = loadFromEnv("ACCOUNT_ID", "")
}
