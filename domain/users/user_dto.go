package users

import (
	Validators "github.com/MatheusAbdias/simple_payment_service/domain/users/validators"
	"github.com/go-playground/validator/v10"
)

func init() {
	customValidator := validator.New()

	customValidator.RegisterValidation("document", Validators.ValidateDocument)
}

type UserDTO struct {
	Id       string `json:"id"       required:"false,uuid"`
	FullName string `json:"fullName" required:"false,min=3,max=255"`
	Email    string `json:"email"    required:"false,email"`
	Document string `json:"document" required:"false,min=11,max=14,document"`
}

func NewUserDTO() *UserDTO {
	return &UserDTO{}
}
