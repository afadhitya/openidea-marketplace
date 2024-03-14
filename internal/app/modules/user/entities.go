package userentities

type User struct {
	Id               string `json:"id" db:"id"`
	Name             string `json:"name" db:"name"`
	Username         string `json:"username" db:"username"`
	Password         string `json:"-" db:"password"`
	ProductSoldTotal int    `json:"product_sold_total" db:"product_sold_total"`
}
