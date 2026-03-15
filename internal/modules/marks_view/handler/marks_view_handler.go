package marksviewhandler

import (
	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/marks_view/dto"
	thisservice "wms/internal/modules/marks_view/service"
	thismodel "wms/internal/modules/marks_view/model"
	crudhandler "wms/pkg/crud_module/handler"
)

type Handler = crudhandler.CRUDHandler[
	*thismodel.MarkView,
	thisdto.CreateRequest,
	thisdto.SetRequest,
	thisdto.GetResponse,
]

func NewHandler(service thisservice.ServiceI) *Handler {
	return crudhandler.NewCRUDHandler[
		*thismodel.MarkView,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "MarksView")
}

func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	crudhandler.RegisterRoutes[
		*thismodel.MarkView,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](r, service, "MarksView", "marksview",
		map[string]bool {
			"getall": 	true,
			"create": 	false,
			"delete": 	false,
			"get": 		false,
			"set": 		false,
		},
	)
}