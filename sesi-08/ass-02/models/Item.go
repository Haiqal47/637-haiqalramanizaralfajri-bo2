package models

type Item struct {
	ItemID      uint   `gorm:"primary_key;auto_increment" json:"lineItemId"`
	ItemCode    string `gorm:"size:255;not null;" json:"itemCode"`
	Description string `gorm:"size:255;not null;" json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     string `gorm:"not null" json:"-"`
}
