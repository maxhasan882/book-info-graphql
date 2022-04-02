package config

import (
	"os"
)

var (
	GinPort    = ""
	GinMode    = ""
	DBHost     = ""
	DBPort     = ""
	DBUser     = ""
	DBPassword = ""
	DBName     = ""
)

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

func SetEnvironment() {
	GinPort = GetEnvDefault("GIN_PORT", "8080")
	GinMode = os.Getenv("GIN_MODE")
	DBHost = GetEnvDefault("DB_HOST", "localhost")
	DBPort = os.Getenv("DB_PORT")
	DBUser = GetEnvDefault("DB_USER", "postgres")
	DBPassword = GetEnvDefault("DB_PASSWORD", "postgres")
	DBName = GetEnvDefault("DB_NAME", "")

}
