package dto

type UserDTO struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Document string `json:"document"`
}
