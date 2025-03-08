package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllCoffees(c *gin.Context) {
	c.JSON(200, CoffeesResponse{Coffees: coffees})
}

func GetCoffeeByID(c *gin.Context) {
	id := c.Param("id")
	for _, coffee := range coffees {
		if coffee.ID == id {
			c.JSON(http.StatusOK, coffee)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
}

func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

func SearchCoffees(c *gin.Context) {
	name := c.Query("name")
	typeQuery := c.Query("type")
	var results []Coffee
	for _, coffee := range coffees {
		if (name != "" && containsIgnoreCase(coffee.Name, name)) || (typeQuery != "" && containsIgnoreCase(coffee.Type, typeQuery)) {
			results = append(results, coffee)
		}
	}
	c.JSON(http.StatusOK, gin.H{"coffees": results})
}

func CreateOrder(c *gin.Context) {
	var newOrder Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	found := false
	for _, coffee := range coffees {
		if coffee.ID == newOrder.CoffeeID {
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Coffee not found"})
		return
	}

	if newOrder.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity, must be greater than 0"})
		return
	}

	newOrder.OrderID = "o" + time.Now().Format("150405")
	now := time.Now()
	newOrder.CreatedAt = now.Format(time.RFC3339)
	newOrder.EstimatedDelivery = now.Add(time.Minute * time.Duration(newOrder.Quantity*5)).Format(time.RFC3339)
	newOrder.Status = "Pending"

	orders[newOrder.OrderID] = newOrder
	c.JSON(http.StatusCreated, newOrder)
}

func GetOrderStatus(c *gin.Context) {
	orderID := c.Param("order_id")
	if order, exists := orders[orderID]; exists {
		c.JSON(http.StatusOK, order)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}
