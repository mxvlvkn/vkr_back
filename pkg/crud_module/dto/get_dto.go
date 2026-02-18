package cruddto

type GetRequest struct {
	ID    		uint64 	  `json:"id" binding:"required"`
}

type GetResponse any