package numenclaturesviewhandler

import (
	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/numenclatures_view/dto"
	thisservice "wms/internal/modules/numenclatures_view/service"
	thismodel "wms/internal/modules/numenclatures_view/model"
	crudhandler "wms/pkg/crud_module/handler"
)

type Handler = crudhandler.CRUDHandler[
	*thismodel.NumenclatureView,
	thisdto.CreateRequest,
	thisdto.SetRequest,
	thisdto.GetResponse,
]

func NewHandler(service thisservice.ServiceI) *Handler {
	return crudhandler.NewCRUDHandler[
		*thismodel.NumenclatureView,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "NumenclaturesView")
}

func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	crudhandler.RegisterRoutes[
		*thismodel.NumenclatureView,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](r, service, "NumenclaturesView", "numenclaturesview",
		map[string]bool {
			"getall": 	true,
			"create": 	false,
			"delete": 	false,
			"get": 		false,
			"set": 		false,
		},
	)
}