package service

import (
	"time"

	"github.com/SinusCosinusTangen/devcamp_pe_james_frederix_rolianto/models"
	"github.com/SinusCosinusTangen/devcamp_pe_james_frederix_rolianto/utils"
)

func GetVariants() []models.Variant {
	var variants []models.Variant
	utils.Database.Find(&variants)

	return variants
}

func GetVariant(id int) models.Variant {
	var variant models.Variant
	utils.Database.First(&variant, id)

	return variant
}

func GetVariantByProductId(product_id int) []models.Variant {
	var variants []models.Variant
	utils.Database.Where("product_id = ?", product_id).Find(&variants)

	return variants
}

func CreateVariant(variantRequest models.VariantRequest) models.Variant {
	variant := models.Variant{
		Name:      variantRequest.Name,
		ImageUrl:  variantRequest.ImageUrl,
		Price:     variantRequest.Price,
		Stock:     variantRequest.Stock,
		Discount:  variantRequest.Discount,
		Weight:    variantRequest.Weight,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	utils.Database.Create(&variant)

	return variant
}

func UpdateVariant(variant models.Variant) models.Variant {
	updated_variant := &variant
	updated_variant.UpdatedAt = time.Now()
	utils.Database.Save(&updated_variant)

	return variant
}

func DeleteVariant(variant models.Variant) {
	utils.Database.Delete(&variant)
}
