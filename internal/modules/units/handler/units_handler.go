package unitshandler

import (
	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/units/dto"
	thisservice "wms/internal/modules/units/service"
	thismodel "wms/internal/modules/units/model"
	crudhandler "wms/pkg/crud_module/handler"
)

type Handler = crudhandler.CRUDHandler[
	*thismodel.Unit,
	thisdto.CreateRequest,
	thisdto.SetRequest,
	thisdto.GetResponse,
]

func NewHandler(service thisservice.ServiceI) *Handler {
	return crudhandler.NewCRUDHandler[
		*thismodel.Unit,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "Units")
}

func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	crudhandler.RegisterRoutes[
		*thismodel.Unit,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](r, service, "Units", "units",
		map[string]bool {
			"getall": 	true,
			"create": 	true,
			"delete": 	true,
			"get": 		true,
			"set": 		true,
		},
	)
}