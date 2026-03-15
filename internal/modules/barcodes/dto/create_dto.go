package barcodesdto

type CreateRequest struct {
	Code       		string      `json:"code" binding:"required,min=12,max=100"`
	NumenclatureID  uint64      `json:"numenclatureID" binding:"required"`
}