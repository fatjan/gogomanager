package dto

type Duck struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetAllDuckResponse struct {
	Ducks []*Duck
}
