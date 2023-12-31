package main

import (
	"GoWeb/Architecture"

	"github.com/gin-gonic/gin"
)

func main() {
	diContainer := Architecture.CreateDIContainer()
	diContainer.CreateIDbInit().InitDatabase()
	r := gin.Default()
	Architecture.InitRoutes(r)
	r.Run(":8080")
}
