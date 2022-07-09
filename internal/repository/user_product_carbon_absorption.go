package repository

import (
	"strings"
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type UserProductCarbonAbsorption interface {
	Create(ctx context.Context, data model.UserProductCarbonAbsorption) error
	Find(ctx context.Context,filter *dto.FilterUserProductCarbonAbsorption, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.UserProductCarbonAbsorption, *dto.PaginationInfo, error)
}

type user_product_carbon_absorption struct {
	Db *gorm.DB
}

func NewUserProductCarbonAbsorption(db *gorm.DB) *user_product_carbon_absorption {
	return &user_product_carbon_absorption{
		db,
	}
}

func (u *user_product_carbon_absorption) Find(ctx context.Context,filter *dto.FilterUserProductCarbonAbsorption, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.UserProductCarbonAbsorption, *dto.PaginationInfo, error) {
	var user_product_carbon_absorption []model.UserProductCarbonAbsorption
	var count int64

	query := u.Db.WithContext(ctx).Model(&model.UserProductCarbonAbsorption{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	if filter.UserID != ""{
		query = query.Where("user_id = ?  ", filter.UserID)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&user_product_carbon_absorption).Error

	return user_product_carbon_absorption, dto.CheckInfoPagination(paginate, count), err
}

func (u *user_product_carbon_absorption) Create(ctx context.Context, data model.UserProductCarbonAbsorption) error {
	return u.Db.WithContext(ctx).Model(&model.UserProductCarbonAbsorption{}).Create(&data).Error
}