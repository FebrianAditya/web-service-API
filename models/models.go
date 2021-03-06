package models

import (
	"gorm.io/gorm"
)

// User build tables
type User struct {
	gorm.Model

	ID 			uint64 `gorm:"primary_key:auto_increment" json:"id"` // NULIS COLUMN NYA HARUS BESAR, NNTI DIOMELIN
	Name 		string `gorm:"type:varchar(255)" json:"name"`
	Email 		string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password	string `gorm:"->;<-;not null" json:"-"`
	CreatedAt 	int64 `gorm:"autoCreateTime"`
  	UpdatedAt 	int64 `gorm:"autoCreateTime"`
}

// Order build table Orders
type Order struct {
	gorm.Model

	OrderID 		uint64 `gorm:"primary_key:auto_increment" json:"id"`
	CustomerName 	string `gorm:"type:varchar(255)" json:"customerName"`
	OrderedAT		int16
	Items			[]Item
}

// Item build table Items
type Item struct {
	gorm.Model

	LineItemID 	uint64 `gorm:"primary_key:auto_increment" json:"id"`
	ItemCode 	string `gorm:"type:varchar(255)" json:"itemCode"`
	ItemName 	string `gorm:"type:varchar(255)" json:"itemName"`
	Quantity 	uint16 `gorm:"type:integer" json:"quantity"`
	OrderID 	uint16 `gorm:"not null" json:"-"`
}