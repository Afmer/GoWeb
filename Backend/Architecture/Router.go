package Architecture

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/Home/Index", func(c *gin.Context) {
		container := new(DIContainer)
		controller := container.GetHomeController(c)
		controller.Index()
	})
}
