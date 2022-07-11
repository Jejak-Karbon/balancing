package product_carbon_absorption

import (
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	ProductCarbonAbsorptionRepository repository.ProductCarbonAbsorption
}

type Service interface {
	Find(ctx context.Context,filter *dto.FilterProductCarbonAbsorption,payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.ProductCarbonAbsorption], error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		ProductCarbonAbsorptionRepository: f.ProductCarbonAbsorptionRepository,
	}
}

func (s *service) Find(ctx context.Context,filter *dto.FilterProductCarbonAbsorption,payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.ProductCarbonAbsorption], error) {

	ProductCarbonAbsorptions, info, err := s.ProductCarbonAbsorptionRepository.Find(ctx,filter,payload, &payload.Pagination)
	
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.ProductCarbonAbsorption])
	result.Datas = ProductCarbonAbsorptions
	result.PaginationInfo = *info

	return result, nil
}