package http

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/app/category_carbon_absorption"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/app/product_carbon_absorption"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/app/user_product_carbon_absorption"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	category_carbon_absorption.NewHandler(f).Route(e.Group("/categories_carbon_absorption"))
	product_carbon_absorption.NewHandler(f).Route(e.Group("/products_carbon_absorption"))
	user_product_carbon_absorption.NewHandler(f).Route(e.Group("/user_product_carbon_absorption"))
}
