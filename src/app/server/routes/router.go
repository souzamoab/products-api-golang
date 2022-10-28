package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/souzamoab/products-api-golang/src/app/controllers"
)

func ConfigRoutes() *gin.Engine {
	router := gin.Default()

	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/", controllers.ShowAllProducts)
			products.GET("/health", controllers.Health)
			products.GET("/:id", controllers.GetProductById)
			products.GET("/filter", controllers.ShowProductByTitle)
			products.POST("/", controllers.CreateProduct)
			products.PUT("/:id", controllers.EditProduct)
			products.DELETE("/:id", controllers.DeleteProduct)
		}
	}
	return router
}
