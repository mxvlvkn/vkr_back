package usersviewhandler

import (
	"github.com/gin-gonic/gin"

	thisdto "wms/internal/modules/users_view/dto"
	thisservice "wms/internal/modules/users_view/service"
	thismodel "wms/internal/modules/users_view/model"
	crudhandler "wms/pkg/crud_module/handler"
)

type Handler = crudhandler.CRUDHandler[
	*thismodel.UserView,
	thisdto.CreateRequest,
	thisdto.SetRequest,
	thisdto.GetResponse,
]

func NewHandler(service thisservice.ServiceI) *Handler {
	return crudhandler.NewCRUDHandler[
		*thismodel.UserView,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](service, "UsersView")
}

func RegisterRoutes(r *gin.RouterGroup, service thisservice.ServiceI) {
	crudhandler.RegisterRoutes[
		*thismodel.UserView,
		thisdto.CreateRequest,
		thisdto.SetRequest,
		thisdto.GetResponse,
	](r, service, "UsersView", "usersview",
		map[string]bool {
			"getall": 	true,
			"create": 	false,
			"delete": 	false,
			"get": 		false,
			"set": 		false,
		},
	)
}