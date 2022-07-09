package repository

import (
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type UserProductCarbonAbsorption interface {
	Create(ctx context.Context, data model.UserProductCarbonAbsorption) error
}

type user_product_carbon_absorption struct {
	Db *gorm.DB
}

func NewUserProductCarbonAbsorption(db *gorm.DB) *user_product_carbon_absorption {
	return &user_product_carbon_absorption{
		db,
	}
}

func (u *user_product_carbon_absorption) Create(ctx context.Context, data model.UserProductCarbonAbsorption) error {
	return u.Db.WithContext(ctx).Model(&model.UserProductCarbonAbsorption{}).Create(&data).Error
}