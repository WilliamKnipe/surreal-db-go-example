package handlers

type AddUsersRequest struct {
	Cafes []Cafe `json:"cafes"`
}

type Cafe struct {
	Name   string `json:"name"`
	Coffee string `json:"coffee"`
	Chairs string `json:"chairs"`
}
