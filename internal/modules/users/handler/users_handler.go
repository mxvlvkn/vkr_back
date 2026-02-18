package usershandler

import (
	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/users/dto"
	thisservice "wms/internal/modules/users/service"
	thismodel "wms/internal/modules/users/model"
	crudhandler "wms/pkg/crud_module/handler"
)

type Handler = crudhandler.CRUDHandler[
	*thismodel.User,
	thisdto.CreateRequest,
	thisdto.SetRequest,
	thisdto.GetResponse,
]

func NewHandler(service thisservice.ServiceI) *Handler {
	return crudhandler.NewCRUDHandler[
		*thismodel.User,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "Users")
}

func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	crudhandler.RegisterRoutes[
		*thismodel.User,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](r, service, "Users", "users",
		map[string]bool {
			"getall": 	false,
			"create": 	true,
			"delete": 	true,
			"get": 		true,
			"set": 		true,
		},
	)
}