package numenclaturesdto

type GetResponse struct {
	Name       		string    `json:"name" binding:"required"`
	UseSerial    	bool      `json:"useSerial" binding:"required"`
	UnitID 			uint64    `json:"unitID" binding:"required"`
	ManufacturerID 	uint64    `json:"manufacturerID" binding:"required"`
	Article		 	string    `json:"article" binding:"required"`
	UseMarks		bool      `json:"useMarks" binding:"required"`
	ImageURL		string    `json:"previewImage" binding:"required"`
}