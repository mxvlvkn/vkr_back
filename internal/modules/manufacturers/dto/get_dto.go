package manufacturersdto

type GetResponse struct {
	Name       		string      `json:"name" binding:"required"`
	Country    		string      `json:"country" binding:"required"`
	INN 			string      `json:"inn" binding:"required"`
	UrAddress 		string      `json:"urAddress" binding:"required"`
	FactAddress 	string      `json:"factAddress" binding:"required"`
	FIO 			string      `json:"fio" binding:"required"`
	Phone 			string      `json:"phone" binding:"required"`
	Email 			string      `json:"email" binding:"required"`
}