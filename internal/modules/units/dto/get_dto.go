package unitsdto

type GetResponse struct {
	Name      		string    `json:"name" binding:"required"`
	Code   			uint   	  `json:"code" binding:"required"`
	Sign       		string    `json:"sign" binding:"required"`
}