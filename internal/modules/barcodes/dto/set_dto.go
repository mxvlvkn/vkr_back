package barcodesdto

type SetRequest struct {
	ID				uint64	  `json:"id" binding:"required"`
	Code       		string      `json:"code" binding:"required,min=12,max=100"`
	NumenclatureID  uint64      `json:"numenclatureID" binding:"required"`
}

func (sr SetRequest) GetID() uint64 {
	return sr.ID
}