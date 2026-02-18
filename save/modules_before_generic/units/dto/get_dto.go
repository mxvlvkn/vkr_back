package unitsdto

type GetRequest struct {
	ID    		uint64 	  `json:"id" binding:"required"`
}

type GetResponse struct {
	Name      		string    `json:"name" binding:"required"`
	Code   			uint   	  `json:"code" binding:"required"`
	Sign       		string    `json:"sign" binding:"required"`
}