package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/souzamoab/products-api-golang/src/app/controllers/dto"
	"github.com/souzamoab/products-api-golang/src/app/database"
	"github.com/souzamoab/products-api-golang/src/app/entity"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDatabase()

	var productCreateDTO dto.ProductCreateDTO

	if err := c.ShouldBindJSON(&productCreateDTO); err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	product := entity.Product{
		Title:       productCreateDTO.Title,
		Description: productCreateDTO.Description,
		Price:       productCreateDTO.Price,
		Quantity:    productCreateDTO.Quantity}

	if err := db.Create(&product).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Product not created: " + err.Error(),
		})
		return
	}

	c.JSON(201, product)
}

func ShowAllProducts(c *gin.Context) {
	db := database.GetDatabase()
	var products []entity.Product
	err := db.Find(&products).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Products not found: " + err.Error(),
		})
		return
	}

	c.JSON(200, products)
}

func GetProductById(c *gin.Context) {
	db := database.GetDatabase()

	id := c.Param("id")

	var product entity.Product

	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Product not found: " + err.Error(),
		})
		return
	}

	c.JSON(200, product)
}

func Health(c *gin.Context) {
	c.JSON(200, "API is healthy")
}

func ShowProductByTitle(c *gin.Context) {
	db := database.GetDatabase()

	name := c.Query("title")

	var product entity.Product

	err := db.Where("title = ?", name).First(&product).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Product not found: " + err.Error(),
		})
		return
	}

	c.JSON(200, product)
}

func DeleteProduct(c *gin.Context) {

	db := database.GetDatabase()
	id := c.Param("id")

	var product entity.Product

	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		c.JSON(400, gin.H{
			"id": "Invalid" + err.Error(),
		})
		return
	}

	err = db.Delete(&product).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Product not found " + err.Error(),
		})
		return
	}

	c.JSON(204, "")
}

func EditProduct(c *gin.Context) {

	db := database.GetDatabase()

	id := c.Param("id")

	var product entity.Product
	var productUpdateDTO dto.ProductUpdateDTO

	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(400, gin.H{
			"error": "Product not found: " + err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&productUpdateDTO); err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	product.Title = productUpdateDTO.Title
	product.Description = productUpdateDTO.Description
	product.Price = productUpdateDTO.Price
	product.Quantity = productUpdateDTO.Quantity

	db.Save(&product)

	c.JSON(200, product)

}
