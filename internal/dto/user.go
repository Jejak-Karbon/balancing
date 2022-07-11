package dto

import (
	_ "github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
)

type UserProfileResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	CityID string `json:"city_id"`
}
