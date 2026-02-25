package routers

import (
	_ "palap_backend/docs"
	v1 "palap_backend/routers/api/v1"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	apiV1.Use()
	{
		// Service endpoints (existing)
		apiV1.GET("/service", v1.GetServices)
		apiV1.GET("/service/:id", v1.GetService)
		apiV1.POST("/service", v1.CreateBook)
		apiV1.PUT("/service/:id", v1.UpdateService)
		apiV1.DELETE("/service/:id", v1.DeleteService)

		// User endpoints
		apiV1.POST("/payment", v1.Payment)
		apiV1.GET("/get_orders", v1.GetOrders)
		apiV1.GET("/get_order_details", v1.GetOrderDetails)

		// Admin management endpoints (CRUD)
		apiV1.GET("/manage_admin", v1.GetAdmins)
		apiV1.GET("/manage_admin/:id", v1.GetAdmin)
		apiV1.POST("/manage_admin", v1.CreateAdmin)
		apiV1.PUT("/manage_admin/:id", v1.UpdateAdmin)
		apiV1.DELETE("/manage_admin/:id", v1.DeleteAdmin)

		// Freelancer management endpoints (CRUD)
		apiV1.GET("/manage_freelance", v1.GetFreelancers)
		apiV1.GET("/manage_freelance/:id", v1.GetFreelancer)
		apiV1.POST("/manage_freelance", v1.CreateFreelancer)
		apiV1.PUT("/manage_freelance/:id", v1.UpdateFreelancer)
		apiV1.DELETE("/manage_freelance/:id", v1.DeleteFreelancer)

		// Freelancer specific endpoints
		apiV1.POST("/registration", v1.RegisterFreelancer)
		apiV1.GET("/get_freelance_job", v1.GetFreelanceJobs)

		// Product and Service management endpoints (CRUD)
		apiV1.GET("/manage_product_and_service", v1.GetProductServices)
		apiV1.GET("/manage_product_and_service/:id", v1.GetProductService)
		apiV1.POST("/manage_product_and_service", v1.CreateProductService)
		apiV1.PUT("/manage_product_and_service/:id", v1.UpdateProductService)
		apiV1.DELETE("/manage_product_and_service/:id", v1.DeleteProductService)

		// Shop location management endpoints (CRUD)
		apiV1.GET("/manage_shop_location", v1.GetShopLocations)
		apiV1.GET("/manage_shop_location/:id", v1.GetShopLocation)
		apiV1.POST("/manage_shop_location", v1.CreateShopLocation)
		apiV1.PUT("/manage_shop_location/:id", v1.UpdateShopLocation)
		apiV1.DELETE("/manage_shop_location/:id", v1.DeleteShopLocation)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
