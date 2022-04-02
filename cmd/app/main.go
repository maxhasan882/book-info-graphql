package main

import (
	"bookinfo/cmd/app/config"
	"bookinfo/internal/app/adapter"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	initEnv()
}

func initEnv() {
	log.Printf("Loading environment settings.")

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No local env file. Using global OS environment variables")
	}
	config.SetEnvironment()
	log.Printf("Connecting to database")
	//db.Connect()
}

func main() {
	r := gin.Default()
	r = adapter.Routes(r)
	err := r.SetTrustedProxies([]string{"192.168.1.107"})
	if err != nil {
		return
	}
	r.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{"payload": "Method Not Allowed"})
	})

	if runError := r.Run(":" + config.GinPort); runError != nil {
		log.Println(runError)
	}
}
