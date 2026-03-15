package barcodeshandler

import (
	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/barcodes/dto"
	thisservice "wms/internal/modules/barcodes/service"
	thismodel "wms/internal/modules/barcodes/model"
	crudhandler "wms/pkg/crud_module/handler"
)

type Handler = crudhandler.CRUDHandler[
	*thismodel.Barcode,
	thisdto.CreateRequest,
	thisdto.SetRequest,
	thisdto.GetResponse,
]

func NewHandler(service thisservice.ServiceI) *Handler {
	return crudhandler.NewCRUDHandler[
		*thismodel.Barcode,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "Barcodes")
}



func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	crudhandler.RegisterRoutes[
		*thismodel.Barcode,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](r, service, "Barcodes", "barcodes",
		map[string]bool {
			"getall": 	false,
			"create": 	true,
			"delete": 	true,
			"get": 		true,
			"set": 		true,
		},
	)
}