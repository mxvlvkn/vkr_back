package usersdto

type GetRequest struct {
	ID    		uint64 	  `json:"id" binding:"required"`
}

type GetResponse struct {
	Login      		string    `json:"login" binding:"required"`
	Name       		string    `json:"name" binding:"required"`
	Surname    		string    `json:"surname" binding:"required"`
	Patronymic 		string    `json:"patronymic" binding:"required"`
	RoleID       	uint64    `json:"role_id" binding:"required"`
}