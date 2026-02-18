package unitsdto

type CreateRequest struct {
	Name       		string    `json:"name" binding:"required,min=2,max=40"`
	Sign    		string    `json:"sign" binding:"required,min=1,max=10"`
	Code 			uint      `json:"code" binding:"required,gt=0"`
}