package rolesdto

type GetResponse struct {
	Name      		string    `json:"name" binding:"required"`
}