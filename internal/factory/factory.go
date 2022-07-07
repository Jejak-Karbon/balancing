package factory

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/database"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
)

type Factory struct {
	CategoryCarbonAbsorptionRepository repository.CategoryCarbonAbsorption
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		CategoryCarbonAbsorptionRepository: repository.NewCategoryCarbonAbsorption(db),
	}
}
