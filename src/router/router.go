package router

import (
	"github.com/SinusCosinusTangen/devcamp_pe_james_frederix_rolianto/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	apiRoute := router.Group("/api")
	{
		apiRoute.POST("product", handler.CreateProduct)
		apiRoute.GET("products", handler.GetAllProducts)
		apiRoute.GET("product/:id", handler.GetProduct)
		apiRoute.PUT("product", handler.UpdateProduct)
		apiRoute.DELETE("product", handler.DeleteProduct)

		apiRoute.POST("variant", handler.CreateVariant)
		apiRoute.GET("variants", handler.GetAllVariants)
	}

	return router
}
