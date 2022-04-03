package config

import (
	"os"
)

var (
	DBHost     = ""
	DBPort     = ""
	DBUser     = ""
	DBPassword = ""
	DBName     = ""
)

// GetEnvDefault Get environment from os, return default if not exist
func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

// SetEnvironment Get environment and set to config variable
func SetEnvironment() {
	DBHost = GetEnvDefault("DB_HOST", "localhost")
	DBPort = os.Getenv("DB_PORT")
	DBUser = GetEnvDefault("DB_USER", "postgres")
	DBPassword = GetEnvDefault("DB_PASSWORD", "postgres")
	DBName = GetEnvDefault("DB_NAME", "")

}
