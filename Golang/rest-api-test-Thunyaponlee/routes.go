package main

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter configures the application routes
func SetupRouter() *gin.Engine {
	// Create default router with Logger and Recovery middleware
	router := gin.Default()

	// Apply our custom logging middleware
	router.Use(LoggingMiddleware())

	// Coffee related endpoints
	router.GET("/coffees", GetAllCoffees)
	router.GET("/coffees/search", SearchCoffees)
	router.GET("/coffees/:id", GetCoffeeByID)

	// Order related endpoints
	router.POST("/orders", CreateOrder)
	router.GET("/orders/:order_id", GetOrderStatus)

	return router
}
