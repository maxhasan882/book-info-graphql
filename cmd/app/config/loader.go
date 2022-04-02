package config

import (
	"os"
)

var (
	Prefix       = "Bearer"
	Token        = ""
	JWTSecret    = ""
	RedisAddr    = ""
	GinPort      = ""
	GinMode      = ""
	DatabaseName = ""
	MongoSSL     = ""
	MongoUrl     = ""
)

func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

func SetEnvironment() {
	Prefix = os.Getenv("PREFIX")
	Token = os.Getenv("TOKEN")
	JWTSecret = os.Getenv("JWT_SECRET")
	RedisAddr = os.Getenv("REDIS_DB")
	GinPort = GetEnvDefault("GIN_PORT", "8080")
	GinMode = os.Getenv("GIN_MODE")
	GinMode = os.Getenv("DATABASE_NAME")
	MongoSSL = os.Getenv("MONGO_SSL")
}
