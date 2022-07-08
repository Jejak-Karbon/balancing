package factory

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/database"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
)

type Factory struct {
	CategoryCarbonAbsorptionRepository repository.CategoryCarbonAbsorption
	ProductCarbonAbsorptionRepository repository.ProductCarbonAbsorption
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		ProductCarbonAbsorptionRepository: repository.NewProductCarbonAbsorption(db),
	}
}
