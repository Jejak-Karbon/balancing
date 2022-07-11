package repository

import (
	"strings"
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type ProductCarbonAbsorption interface {
	Find(ctx context.Context,filter *dto.FilterProductCarbonAbsorption,payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.ProductCarbonAbsorption, *dto.PaginationInfo, error)
}

type product_carbon_absorption struct {
	Db *gorm.DB
}

func NewProductCarbonAbsorption(db *gorm.DB) *product_carbon_absorption {
	return &product_carbon_absorption{
		db,
	}
}

func (p *product_carbon_absorption) Find(ctx context.Context,filter *dto.FilterProductCarbonAbsorption,payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.ProductCarbonAbsorption, *dto.PaginationInfo, error) {
	var products_carbon_absorption []model.ProductCarbonAbsorption
	var count int64

	query := p.Db.WithContext(ctx).Model(&model.ProductCarbonAbsorption{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	if filter.CategoryCarbonAbsorptionID != ""{
		query = query.Where("category_carbon_absorption_id = ?  ", filter.CategoryCarbonAbsorptionID)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&products_carbon_absorption).Error

	return products_carbon_absorption, dto.CheckInfoPagination(paginate, count), err
}

