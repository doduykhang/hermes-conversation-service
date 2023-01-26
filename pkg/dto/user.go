package dto

type CreateUser struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	UserName string `json:"userName"`
	Avatar string `json:"avatar"`
}

type UserDTO struct {
	Audit
	ID string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	UserName string `json:"userName"`
	Avatar string `json:"avatar"`
}
