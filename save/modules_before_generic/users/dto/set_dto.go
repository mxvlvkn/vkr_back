package usersdto

type SetRequest struct {
	ID				uint64	  `json:"id" binding:"required"`
	Login    		string 	  `json:"login" binding:"required,min=4,max=40"`
	Name       		string    `json:"name" binding:"required,min=4,max=40"`
	Surname    		string    `json:"surname" binding:"required,min=2,max=40"`
	Patronymic 		string    `json:"patronymic" binding:"required,min=4,max=40"`
	RoleID       	uint64    `json:"role" binding:"required"`
	Password 		string 	  `json:"password" binding:"required,min=8,max=40"`
	RepeatPassword  string 	  `json:"repeatPassword" binding:"required,eqfield=Password"`
	IsSetPassword	bool	  `json:"isSetPassword" binding:"required"`
}

type SetResponse struct {
	Status bool `json:"status"`
}