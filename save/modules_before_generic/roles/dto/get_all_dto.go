package rolesdto

import (
	"wms/internal/modules/roles/model"
)

type GetAllResponse struct {
	Roles []rolesmodel.Role `json:"roles"`
}