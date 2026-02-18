package cruddto

type SetRequest interface {
	GetID() uint64
}

type SetResponse struct {
	Status bool `json:"status"`
}