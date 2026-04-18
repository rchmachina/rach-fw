package model
type Address struct {
	ID         int64  `gorm:"column:id" json:"id"`
	Street     string `gorm:"column:street" json:"street"`
	City       string `gorm:"column:city" json:"city"`
	Province   string `gorm:"column:province" json:"province"`
	PostalCode string `gorm:"column:postal_code" json:"postal_code"`
	Country    string `gorm:"column:country" json:"country"`
}


