package usersdto

import (
	"wms/internal/modules/users/model"
)

type GetAllResponse struct {
	Users []usersmodel.UserView `json:"users"`
}