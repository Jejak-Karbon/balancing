package model

import (
	"time"

	"gorm.io/gorm"
)

func (UserProductCarbonAbsorption) TableName() string {
    return "user_product_carbon_absorption"
}

type UserProductCarbonAbsorption struct {
	ID       uint   	`gorm:"primarykey;autoIncrement"`
	UserID       uint   	`json:"user_id" gorm:"not null"`
	ProductCarbonAbsorptionID       uint   	`json:"product_carbon_absorption_id" gorm:"not null"`
	Model
	ProductCarbonAbsorption ProductCarbonAbsorption
}

func (u *UserProductCarbonAbsorption) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct User
// gorm call this method before they execute query
func (u *UserProductCarbonAbsorption) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}