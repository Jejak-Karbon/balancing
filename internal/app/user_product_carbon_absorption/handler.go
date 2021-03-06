package user_product_carbon_absorption

import (
	"fmt"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/middleware"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {

	payload := new(dto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	filter := new(dto.FilterUserProductCarbonAbsorption)

	if err := c.Bind(filter); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(filter); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	
	result, err := h.service.Find(c.Request().Context(),filter, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get user carbon absorption success", &result.PaginationInfo).Send(c)
}

func (h *handler) Create(c echo.Context) error {

	payloadToken := middleware.GetIDFromToken(c)
	var user_id uint = uint(payloadToken.(float64))

	payload := new(dto.CreateUserProductCarbonAbsorption)

	if err := c.Bind(payload); err != nil {
		fmt.Println(payload)
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		response := res.ErrorBuilder(&res.ErrorConstant.Validation, err)
		return response.Send(c)
	}

	result2, err2 := h.service.Create(c.Request().Context(), user_id, payload)
	if err2 != nil {
		return res.ErrorResponse(err2).Send(c)
	}

	return res.SuccessResponse(result2).Send(c)

}