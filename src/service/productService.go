package service

import (
	"time"

	"github.com/SinusCosinusTangen/devcamp_pe_james_frederix_rolianto/models"
	"github.com/SinusCosinusTangen/devcamp_pe_james_frederix_rolianto/utils"
)

func GetProducts() []models.Product {
	var products []models.Product
	utils.Database.Find(&products)

	return products
}

func GetProduct(id int) models.Product {
	var product models.Product
	utils.Database.First(&product, id)

	return product
}

func CreateProduct(productRequest models.ProductRequest) models.Product {
	db := utils.Database

	product := models.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		ImageUrl:    productRequest.ImageUrl,
		Price:       productRequest.Price,
		Stock:       productRequest.Stock,
		Discount:    productRequest.Discount,
		Weight:      productRequest.Weight,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	db.Create(&product)

	for _, variantRequest := range productRequest.Variants {

		variant := models.Variant{
			Name:      variantRequest.Name,
			ProductID: product.ID,
			ImageUrl:  variantRequest.ImageUrl,
			Price:     variantRequest.Price,
			Stock:     variantRequest.Stock,
			Discount:  variantRequest.Discount,
			Weight:    variantRequest.Weight,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		db.Create(&variant)
	}

	return product
}

func UpdateProduct(product models.Product) models.Product {
	updated_product := &product
	updated_product.UpdatedAt = time.Now()
	utils.Database.Save(&updated_product)

	return product
}

func DeleteProduct(product models.Product) {
	utils.Database.Delete(&product)
}
