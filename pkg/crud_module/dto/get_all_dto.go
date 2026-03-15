package cruddto

import (
	"wms/pkg/crud_module/model"

)

type GetAllRequest struct {
	PageNum    		int 	  	  `json:"pageNum" binding:"required"`
	Search			string		  `json:"search"`
	FilterMethod	string		  `json:"filterMethod" binding:"required,oneof=default desc asc"`
	FilterField		string		  `json:"filterField"`
	WhereID			int		  	  `json:"whereID"`
	WhereField		string		  `json:"whereField"`
}

type GetAllResponse[ModelT crudmodel.Item] struct {
	Items []ModelT `json:"items"`
}