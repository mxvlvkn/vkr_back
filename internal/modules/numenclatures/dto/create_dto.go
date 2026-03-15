package numenclaturesdto

type CreateRequest struct {
	Name           string 					`json:"name" binding:"required,min=15,max=400"`
	Article        string 					`json:"article" binding:"required,min=10,max=200"`
	UseSerial      bool   					`json:"useSerial" binding:"required"`
	UseMarks       bool   					`json:"useMarks" binding:"required"`
	UnitID         uint64 					`json:"unitID" binding:"required"`
	ManufacturerID uint64 					`json:"manufacturerID" binding:"required"`
	ImageURL	   string 					``
	ImageFields	   []string					`json:"imageFields"`
}