package models

import "time"

type Order struct {
	OrderID      uint      `gorm:"primary_key;auto_increment" json:"-"`
	CustomerName string    `gorm:"size:255;not null;" json:"customerName"`
	OrderedAt    time.Time `gorm:"autoCreateTime" json:"orderedAt"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"items"`
}
