package authdto

type LoginRequest struct {
	Login    string `json:"login" binding:"required,min=4,max=40"`
	Password string `json:"password" binding:"required,min=8,max=40"`
}

type LoginResponse struct {
	Status bool `json:"status"`
}