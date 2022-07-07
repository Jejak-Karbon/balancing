package model

import (
	"time"

	"gorm.io/gorm"
)

func (CategoryCarbonAbsorption) TableName() string {
    return "categories_carbon_absorption"
}

type CategoryCarbonAbsorption struct {
	ID       		uint   	`gorm:"primarykey;autoIncrement"`
	Name    	 	string 	`json:"name" gorm:"size:200;not null"`
	Descirption     string 	`json:"description" gorm:"not null"`
	Type    		string 	`json:"type" gorm:"size:50;not null"`
	Model
}

func (c *CategoryCarbonAbsorption) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct User
// gorm call this method before they execute query
func (c *CategoryCarbonAbsorption) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}