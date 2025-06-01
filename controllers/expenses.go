package controllers

import (
	"net/http"

	"github.com/andrewanzechan/expenses-app-web-service/models"
	"github.com/gin-gonic/gin"
)

func GetExpenses(c *gin.Context) {
	var expenses = []models.Expense{
		{
			ID:          "1",
			Description: "Lunch with client",
			Amount:      45.50,
			Category:    "Food",
			CreatedAt:   "2023-10-01T12:00:00Z",
			UpdatedAt:   "2023-10-01T12:00:00Z",
			Date:        "2023-10-01",
		},
	}
	c.IndentedJSON(http.StatusOK, expenses)
}

func CreateExpense(c *gin.Context) {
	var newExpense models.Expense
	if err := c.BindJSON(&newExpense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newExpense.ID = "2" // In a real application, you would generate a unique ID
	newExpense.CreatedAt = "2023-10-01T12:00:00Z"
	newExpense.UpdatedAt = "2023-10-01T12:00:00Z"

	c.IndentedJSON(http.StatusCreated, newExpense)
}

func GetExpenseByID(c *gin.Context) {
	id := c.Param("id")
	if id != "1" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	expense := models.Expense{
		ID:          "1",
		Description: "Lunch with client",
		Amount:      45.50,
		Category:    "Food",
		CreatedAt:   "2023-10-01T12:00:00Z",
		UpdatedAt:   "2023-10-01T12:00:00Z",
		Date:        "2023-10-01",
	}
	c.IndentedJSON(http.StatusOK, expense)
}

func UpdateExpense(c *gin.Context) {
	id := c.Param("id")
	if id != "1" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	var updatedExpense models.Expense
	if err := c.BindJSON(&updatedExpense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedExpense.ID = id
	updatedExpense.UpdatedAt = "2023-10-01T12:00:00Z"

	c.IndentedJSON(http.StatusOK, updatedExpense)
}

func DeleteExpense(c *gin.Context) {
	id := c.Param("id")
	if id != "1" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
		return
	}

	// In a real application, you would delete the expense from the database here

	c.Status(http.StatusNoContent) // 204 No Content
}
