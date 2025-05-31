package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/expenses", getExpenses)

	router.Run("localhost:8080")
}

type Expense struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Date        string  `json:"date"`
}

func getExpenses(c *gin.Context) {
	var expenses = []Expense{
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
