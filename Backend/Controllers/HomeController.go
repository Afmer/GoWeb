package Controllers

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type HomeController struct {
	context *gin.Context
}

func (obj HomeController) Index() {
	obj.context.JSON(http.StatusOK, gin.H{
		"message": "Привет от сервера Gin!",
	})
}

func CreateHomeController(context *gin.Context) *HomeController {
	return &HomeController{
		context: context,
	}
}
