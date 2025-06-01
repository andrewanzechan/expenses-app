package main

import (
	"github.com/andrewanzechan/expenses-app-web-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run("localhost:8080")
}
