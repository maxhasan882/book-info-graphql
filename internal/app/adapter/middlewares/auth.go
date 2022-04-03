package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	c.Next()
}
