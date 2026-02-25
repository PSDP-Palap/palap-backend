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
		apiV1.GET("/service", v1.GetServices)
		apiV1.GET("/service/:id", v1.GetService)
		apiV1.POST("/service", v1.CreateBook)
		apiV1.PUT("/service/:id", v1.UpdateService)
		apiV1.DELETE("/service/:id", v1.DeleteService)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
