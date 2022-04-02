package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func UserRoutes(r *gin.Engine) *gin.Engine {
	controller := UserController{}

	r.GET("/user", func(c *gin.Context) {
		controller.getUser(c)
	})
	return r

}

func (hc UserController) getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "User",
		"user":  "this is a new user",
	})
}
