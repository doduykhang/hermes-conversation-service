package dto

type CreateUser struct {
	ID string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
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
