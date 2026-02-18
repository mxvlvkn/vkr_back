package unitsdto

type DeleteRequest struct {
	ID    		uint64 	  `json:"id" binding:"required"`
}

type DeleteResponse struct {
	Status bool `json:"status"`
}