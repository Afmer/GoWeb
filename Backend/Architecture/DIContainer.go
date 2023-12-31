package Architecture

import (
	"GoWeb/Controllers"
	"GoWeb/Services"

	"GoWeb/Architecture/Interfaces"

	"github.com/gin-gonic/gin"
)

func (obj DIContainer) CreateIDbInit() Interfaces.IDbInit {
	return Services.CreateDbManager(obj.dbSettings)
}

type DIContainer struct {
	dbSettings Services.DbManagerSettings
}

func CreateDIContainer() DIContainer {
	dbSettings := Services.DbManagerSettings{
		ConnectionString: "admin:red_alien@tcp(127.0.0.1:3306)/",
		DatabaseName:     "GoTest",
	}
	result := DIContainer{
		dbSettings: dbSettings,
	}
	return result
}

func (obj DIContainer) GetHomeController(context *gin.Context) Controllers.HomeController {
	var homeController = Controllers.CreateHomeController(context)
	return *homeController
}
