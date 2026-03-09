package manufacturershandler

import (
	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/manufacturers/dto"
	thisservice "wms/internal/modules/manufacturers/service"
	thismodel "wms/internal/modules/manufacturers/model"
	crudhandler "wms/pkg/crud_module/handler"
)

type Handler = crudhandler.CRUDHandler[
	*thismodel.Manufacturer,
	thisdto.CreateRequest,
	thisdto.SetRequest,
	thisdto.GetResponse,
]

func NewHandler(service thisservice.ServiceI) *Handler {
	return crudhandler.NewCRUDHandler[
		*thismodel.Manufacturer,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "Manufacturers")
}



func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	crudhandler.RegisterRoutes[
		*thismodel.Manufacturer,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](r, service, "Manufacturers", "manufacturers",
		map[string]bool {
			"getall": 	true,
			"create": 	true,
			"delete": 	true,
			"get": 		true,
			"set": 		true,
		},
	)
}