package model

import (
	"time"

	_ "github.com/shopspring/decimal"

	"gorm.io/gorm"
)

func (ProductCarbonAbsorption) TableName() string {
    return "products_carbon_absorption"
}

type ProductCarbonAbsorption struct {
	ID       		uint   	`gorm:"primarykey;autoIncrement"`
	UserID    	 	uint 	`json:"user_id" gorm:"not null"`
	CategoryCarbonAbsorptionID    	 	uint 	`json:"category_carbon_absorption_id" gorm:"not null"`
	Name    	 	string 	`json:"name" gorm:"size:200;not null"`
	Description     string 	`json:"description" gorm:"not null"`
	Amount	float32 `json:"amount"  gorm:"not null"`
	Price   float32 `json:"price" sql:"type:decimal(10,2);"`
	Image    		string 	`json:"image" gorm:"size:200;not null"`
	Model
}

func (p *ProductCarbonAbsorption) BeforeCreate(tx *gorm.DB) (err error) {
	p.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct User
// gorm call this method before they execute query
func (p *ProductCarbonAbsorption) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return
}