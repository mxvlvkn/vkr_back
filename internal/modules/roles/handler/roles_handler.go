package roleshandler

import (
	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/roles/dto"
	thisservice "wms/internal/modules/roles/service"
	thismodel "wms/internal/modules/roles/model"
	crudhandler "wms/pkg/crud_module/handler"
)

type Handler = crudhandler.CRUDHandler[
	*thismodel.Role,
	thisdto.CreateRequest,
	thisdto.SetRequest,
	thisdto.GetResponse,
]

func NewHandler(service thisservice.ServiceI) *Handler {
	return crudhandler.NewCRUDHandler[
		*thismodel.Role,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "Roles")
}

func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	crudhandler.RegisterRoutes[
		*thismodel.Role,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](r, service, "Roles", "roles",
		map[string]bool {
			"getall": 	true,
			"create": 	false,
			"delete": 	false,
			"get": 		true,
			"set": 		false,
		},
	)
}