package barcodesdto

type GetResponse struct {
	Code       		string      `json:"code" binding:"required"`
	NumenclatureID  uint64      `json:"numenclatureID" binding:"required"`
}