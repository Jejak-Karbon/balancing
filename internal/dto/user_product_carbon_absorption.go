package dto

import "github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"

type FilterUserProductCarbonAbsorption struct {
	UserID   string   `query:"user_id"`
}

type UserProductCarbonAbsorptionResponse struct {
	UserProductCarbonAbsorption model.UserProductCarbonAbsorption
	User UserProfileResponse
}

type CreateUserProductCarbonAbsorption struct {
	ProductCarbonAbsorptionID        uint `json:"product_carbon_absorption_id" validate:"required"`
}