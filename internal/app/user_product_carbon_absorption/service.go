package user_product_carbon_absorption

import (
	"fmt"
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	_ "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/constant"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"
)

type service struct {
	UserProductCarbonAbsorptionRepository repository.UserProductCarbonAbsorption
}

type Service interface {
	Create(ctx context.Context, user_id uint, payload *dto.CreateUserProductCarbonAbsorption) (string, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserProductCarbonAbsorptionRepository: f.UserProductCarbonAbsorptionRepository,
	}
}

func (s *service) Create(ctx context.Context, user_id uint, payload *dto.CreateUserProductCarbonAbsorption) (string, error) {

	data := model.UserProductCarbonAbsorption{UserID :user_id,ProductCarbonAbsorptionID:payload.ProductCarbonAbsorptionID}
	fmt.Println(data)
	err := s.UserProductCarbonAbsorptionRepository.Create(ctx, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	amount := 2.383
	message :=  "selamat kamu berhasil melakukan penyerapan emisi"

	return message, nil
}

