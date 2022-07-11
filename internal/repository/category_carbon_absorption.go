package repository

import (
	"context"
	"strings"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type CategoryCarbonAbsorption interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.CategoryCarbonAbsorption, *dto.PaginationInfo, error)
}

type category_carbon_absorption struct {
	Db *gorm.DB
}

func NewCategoryCarbonAbsorption(db *gorm.DB) *category_carbon_absorption {
	return &category_carbon_absorption{
		db,
	}
}

func (c *category_carbon_absorption) Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.CategoryCarbonAbsorption, *dto.PaginationInfo, error) {
	var categories_carbon_absorption []model.CategoryCarbonAbsorption
	var count int64

	query := c.Db.WithContext(ctx).Model(&model.CategoryCarbonAbsorption{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&categories_carbon_absorption).Error

	return categories_carbon_absorption, dto.CheckInfoPagination(paginate, count), err
}
