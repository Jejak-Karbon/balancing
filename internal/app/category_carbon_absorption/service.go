package category_carbon_absorption

import (
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	CategoryCarbonAbsorptionRepository repository.CategoryCarbonAbsorption
}

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.CategoryCarbonAbsorption], error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		CategoryCarbonAbsorptionRepository: f.CategoryCarbonAbsorptionRepository,
	}
}

func (s *service) Find(ctx context.Context,payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.CategoryCarbonAbsorption], error) {

	CategoryCarbonAbsortions, info, err := s.CategoryCarbonAbsorptionRepository.Find(ctx,payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.CategoryCarbonAbsorption])
	result.Datas = CategoryCarbonAbsortions
	result.PaginationInfo = *info

	return result, nil
}
