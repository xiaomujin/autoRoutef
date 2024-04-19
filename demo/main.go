package main

import (
	"github.com/gin-gonic/gin"
	"github.com/odinZzzzz/autoRoute/demo/handler"
)

func main() {
	r := gin.Default()
	handler.InitHandler()

	aRoute := handler.InitHandler()
	r.Use(aRoute.RouteMid)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
