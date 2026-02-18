package unitsdto

import (
	"wms/internal/modules/units/model"
)

type GetAllResponse struct {
	Units []unitsmodel.Unit `json:"units"`
}