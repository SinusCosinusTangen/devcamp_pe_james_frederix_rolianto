package handler

import (
	"strconv"

	"github.com/SinusCosinusTangen/devcamp_pe_james_frederix_rolianto/models"
	"github.com/SinusCosinusTangen/devcamp_pe_james_frederix_rolianto/service"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var productRequest models.ProductRequest
	c.BindJSON(&productRequest)
	product := service.CreateProduct(productRequest)
	c.JSON(200, product)
}

func GetAllProducts(c *gin.Context) {
	products := service.GetProducts()

	var productResponse []models.ProductResponse
	for _, product := range products {
		variants := service.GetVariantByProductId(int(product.ID))

		variantReponse := []models.VariantResponse{}
		for _, variant := range variants {
			variantReponse = append(variantReponse, models.VariantResponse{
				ID:       variant.ID,
				Name:     variant.Name,
				ImageUrl: variant.ImageUrl,
				Price:    variant.Price,
				Stock:    variant.Stock,
				Discount: variant.Discount,
				Weight:   variant.Weight,
			})
		}

		productResponse = append(productResponse, models.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			ImageUrl:    product.ImageUrl,
			Price:       product.Price,
			Stock:       product.Stock,
			Discount:    product.Discount,
			Weight:      product.Weight,
			Variants:    variantReponse,
		})
	}

	c.JSON(200, productResponse)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product_id, _ := strconv.Atoi(id)
	product := service.GetProduct(product_id)
	c.JSON(200, product)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	updated_product := service.UpdateProduct(product)
	c.JSON(200, updated_product)
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	service.DeleteProduct(product)
	c.JSON(200, gin.H{"message": "Product deleted"})
}

func CreateVariant(c *gin.Context) {
	var variantRequest models.VariantRequest
	c.BindJSON(&variantRequest)
	variant := service.CreateVariant(variantRequest)
	c.JSON(200, variant)
}

func GetAllVariants(c *gin.Context) {
	variants := service.GetVariants()
	c.JSON(200, variants)
}
