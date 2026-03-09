package manufacturersdto

type CreateRequest struct {
	Name       		string      `json:"name" binding:"required,min=2,max=60"`
	Country    		string      `json:"country" binding:"required,min=2,max=40"`
	INN 			string      `json:"inn" binding:"required,min=10,max=12"`
	UrAddress 		string      `json:"urAddress" binding:"required,min=10,max=200"`
	FactAddress 	string      `json:"factAddress" binding:"required,min=10,max=200"`
	FIO 			string      `json:"fio" binding:"required,min=10,max=200"`
	Phone 			string      `json:"phone" binding:"required,max=20"`
	Email 			string      `json:"email" binding:"required,max=200"`
}