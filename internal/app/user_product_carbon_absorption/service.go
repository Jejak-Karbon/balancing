package user_product_carbon_absorption

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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
	Find(ctx context.Context, filter *dto.FilterUserProductCarbonAbsorption, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[dto.UserProductCarbonAbsorptionResponse], error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserProductCarbonAbsorptionRepository: f.UserProductCarbonAbsorptionRepository,
	}
}

func (s *service) Find(ctx context.Context, filter *dto.FilterUserProductCarbonAbsorption, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[dto.UserProductCarbonAbsorptionResponse], error) {

	UserProductCarbonAbsorptions, info, err := s.UserProductCarbonAbsorptionRepository.Find(ctx, filter, payload, &payload.Pagination)

	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.UserProductCarbonAbsorptionResponse
	var url string

	for _, value := range UserProductCarbonAbsorptions {

		var client = &http.Client{}

		url = fmt.Sprintf("http://users.masudin.space/users/%d", value.UserID))
		request, err := http.NewRequest("GET", url, nil)
		request.Header.Set("Authorization", os.Getenv("RANDOM_KEY"))
		if err != nil {
			return nil, err
		}

		response, err := client.Do(request)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		var result map[string]map[string]interface{}

		json.NewDecoder(response.Body).Decode(&result)

		data = append(data, dto.UserProductCarbonAbsorptionResponse{
			UserProductCarbonAbsorption: value,
			User: dto.UserProfileResponse{
				Name:   result["data"]["name"].(string),
				Email:  result["data"]["email"].(string),
				CityID: result["data"]["city_id"].(string),
			},
		})
	}

	result := new(dto.SearchGetResponse[dto.UserProductCarbonAbsorptionResponse])
	result.Datas = data
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) Create(ctx context.Context, user_id uint, payload *dto.CreateUserProductCarbonAbsorption) (string, error) {

	data := model.UserProductCarbonAbsorption{UserID: user_id, ProductCarbonAbsorptionID: payload.ProductCarbonAbsorptionID}
	fmt.Println(data)
	err := s.UserProductCarbonAbsorptionRepository.Create(ctx, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	message := "selamat kamu berhasil melakukan penyerapan emisi"

	return message, nil
}
