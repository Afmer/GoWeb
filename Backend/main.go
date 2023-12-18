package main

import (
	"GoWeb/Architecture"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Architecture.InitRoutes(r)
	r.Run(":8080")
}
