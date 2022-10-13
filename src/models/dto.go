package models

type ProductRequest struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	ImageUrl    string           `json:"image_url"`
	Price       uint             `json:"price"`
	Stock       uint             `json:"stock"`
	Discount    uint             `json:"discount"`
	Weight      uint             `json:"weight"`
	Variants    []VariantRequest `json:"variants"`
}

type VariantRequest struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	Price    uint   `json:"price"`
	Stock    uint   `json:"stock"`
	Discount uint   `json:"discount"`
	Weight   uint   `json:"weight"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ProductResponse struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	ImageUrl    string            `json:"image_url"`
	Price       uint              `json:"price"`
	Stock       uint              `json:"stock"`
	Discount    uint              `json:"discount"`
	Weight      uint              `json:"weight"`
	Variants    []VariantResponse `json:"variants"`
}

type VariantResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	Price    uint   `json:"price"`
	Stock    uint   `json:"stock"`
	Discount uint   `json:"discount"`
	Weight   uint   `json:"weight"`
}
