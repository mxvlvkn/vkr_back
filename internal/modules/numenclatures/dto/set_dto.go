package numenclaturesdto

type SetRequest struct {
	ID				uint64	  `json:"id" binding:"required"`
	Name       		string    `json:"name" binding:"required,min=15,max=400"`
	UseSerial    	bool      `json:"useSerial" binding:"required"`
	UnitID 			uint64    `json:"unitID" binding:"required"`
	ManufacturerID 	uint64    `json:"manufacturerID" binding:"required"`
	Article		 	string    `json:"article" binding:"required,min=10,max=200"`
	UseMarks		bool      `json:"useMarks" binding:"required"`
}

func (sr SetRequest) GetID() uint64 {
	return sr.ID
}