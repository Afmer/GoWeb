package Architecture

import (
	"GoWeb/Controllers"

	"github.com/gin-gonic/gin"
)

type DIContainer struct {
}

func (obj DIContainer) GetHomeController(context *gin.Context) Controllers.HomeController {
	var homeController = Controllers.CreateHomeController(context)
	return *homeController
}
