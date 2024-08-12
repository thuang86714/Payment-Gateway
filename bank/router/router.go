package router

import (
	"github.com/gin-gonic/gin"

	"github.com/processout-hiring/payment-gateway-thuang86714/bank/middleware"
	"github.com/processout-hiring/payment-gateway-thuang86714/bank/controller"
	middleware "github.com/processout-hiring/payment-gateway-thuang86714/shared/middleware"
)

// SetRoutes sets up the routes for the application
func SetRoutes(router *gin.Engine) *gin.Engine {
	// Apply middleware globally
	router.Use(middleware.LoggingMiddleware())

	// Initialize controllers
	curController := controller.NewController()
	router.POST("/processTransaction", middlewareBank.SendInitialResponse(), curController.ProcessTransaction)

	return router
}