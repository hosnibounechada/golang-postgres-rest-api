package models

type Product struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity float32 `json:"quantity"`
	UserID   int64   `json:"user_id"`
}

type CreateProductDTO struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity float32 `json:"quantity"`
	UserID   int64   `json:"user_id"`
}

type UpdateProductDTO struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity float32 `json:"quantity"`
}
