package entity

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    uint    `json:"quantity"`
}
