package router

import (
	"github.com/processout-hiring/payment-gateway-thuang86714/merchant/controller"
	"github.com/gin-gonic/gin"
)

func SetRoutes(route *gin.Engine) {
	var ctrl merchant_controller.ExampleController
	v1 := route.Group("/v1/examples")
	v1.GET("test/", ctrl.GetExampleData)
	v1.POST("test/", ctrl.CreateExample)
	v1.GET("test/paginated", ctrl.GetExamplePaginated)
	v1.GET("test/relational", ctrl.GetHasManyRelationUserData)
	v1.GET("test/card", ctrl.GetHasManyRelationCreditCardData)
	v1.GET("test/user", ctrl.GetUserDetails)

}