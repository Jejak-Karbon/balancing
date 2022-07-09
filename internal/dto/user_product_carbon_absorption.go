package dto

type FilterUserProductCarbonAbsorption struct {
	UserID   string   `query:"user_id"`
}

type CreateUserProductCarbonAbsorption struct {
	ProductCarbonAbsorptionID        uint `json:"product_carbon_absorption_id" validate:"required"`
}