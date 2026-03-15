package markshandler

import (
	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/marks/dto"
	thisservice "wms/internal/modules/marks/service"
	thismodel "wms/internal/modules/marks/model"
	crudhandler "wms/pkg/crud_module/handler"
)

type Handler = crudhandler.CRUDHandler[
	*thismodel.Mark,
	thisdto.CreateRequest,
	thisdto.SetRequest,
	thisdto.GetResponse,
]

func NewHandler(service thisservice.ServiceI) *Handler {
	return crudhandler.NewCRUDHandler[
		*thismodel.Mark,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "Marks")
}



func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	crudhandler.RegisterRoutes[
		*thismodel.Mark,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](r, service, "Marks", "marks",
		map[string]bool {
			"getall": 	false,
			"create": 	true,
			"delete": 	true,
			"get": 		true,
			"set": 		true,
		},
	)
}