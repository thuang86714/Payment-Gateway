package router

import (
	"github.com/gin-gonic/gin"

	"github.com/processout-hiring/payment-gateway-thuang86714/bank/controller"
	middleware "github.com/processout-hiring/payment-gateway-thuang86714/shared/middleware"
)

// SetRoutes sets up the routes for the application
func SetRoutes(router *gin.Engine, ctr *controller.Controller) *gin.Engine {
	// Apply middleware globally
	router.Use(middleware.LoggingMiddleware())

	router.POST("/processTransaction", ctr.ProcessTransaction)

	return router
}