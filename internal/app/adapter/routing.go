package adapter

import (
	"bookinfo/internal/app/adapter/controller"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) *gin.Engine {
	r = controller.UserRoutes(r)
	return r
}
