package config

import (
	"os"
	"time"
)

var Address string
var Port string
var Token string
var AccountID string

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
