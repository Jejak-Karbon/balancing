package user_product_carbon_absorption

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("", h.Create,middleware.Authentication)
}
