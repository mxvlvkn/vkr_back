package barcodesviewhandler

import (
	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/barcodes_view/dto"
	thisservice "wms/internal/modules/barcodes_view/service"
	thismodel "wms/internal/modules/barcodes_view/model"
	crudhandler "wms/pkg/crud_module/handler"
)

type Handler = crudhandler.CRUDHandler[
	*thismodel.BarcodeView,
	thisdto.CreateRequest,
	thisdto.SetRequest,
	thisdto.GetResponse,
]

func NewHandler(service thisservice.ServiceI) *Handler {
	return crudhandler.NewCRUDHandler[
		*thismodel.BarcodeView,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "BarcodesView")
}

func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	crudhandler.RegisterRoutes[
		*thismodel.BarcodeView,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](r, service, "BarcodesView", "barcodesview",
		map[string]bool {
			"getall": 	true,
			"create": 	false,
			"delete": 	false,
			"get": 		false,
			"set": 		false,
		},
	)
}