package productentities

type Product struct {
	Id            string   `json:"id" db:"id"`
	Name          string   `json:"name" validate:"required,min=5,max=60"`
	Price         float64  `json:"price" validate:"required,min=0"`
	ImageURL      string   `json:"imageUrl" validate:"required,url"`
	Stock         int      `json:"stock" validate:"required,min=0"`
	Condition     string   `json:"condition" validate:"required,oneof=new second"`
	Tags          []string `json:"tags" validate:"required,min=0"`
	IsPurchasable bool     `json:"isPurchasable" validate:"required"`
	PurchaseCount int      `json:"purchasCount" validate:"required,min=0"`
	SellerID      string   `json:"sellerId" validate:"required"`
}
