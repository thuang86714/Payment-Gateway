package router

import (
	"github.com/gin-gonic/gin"

	"github.com/processout-hiring/payment-gateway-thuang86714/gateway/controller"
	middleware "github.com/processout-hiring/payment-gateway-thuang86714/shared/middlerware"
)

// SetRoutes sets up the routes for the application
func SetRoutes(router *gin.Engine) *gin.Engine {
	// Apply middleware globally
	router.Use(middleware.LoggingMiddleware())

	// Initialize controllers
	curController := controller.NewController()
	router.POST("/processPayment", curController.CreateTransaction)
	router.GET("/retrievePayment", curController.GetTransaction)
	router.PATCH("/updatePayment", curController.UpdateTransaction)

	return router
}