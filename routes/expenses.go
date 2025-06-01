package routes

import (
	"github.com/andrewanzechan/expenses-app-web-service/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Define the base path for expenses
	expensesGroup := router.Group("/expenses")
	{
		expensesGroup.GET("/", controllers.GetExpenses)
		expensesGroup.POST("/", controllers.CreateExpense)
		expensesGroup.GET("/:id", controllers.GetExpenseByID)
		expensesGroup.PUT("/:id", controllers.UpdateExpense)
		expensesGroup.DELETE("/:id", controllers.DeleteExpense)
	}
}
