package models

import (
	"github.com/MatheusAbdias/simple_payment_service/domain/dto"
	CustomValidators "github.com/MatheusAbdias/simple_payment_service/pkg/validators"
	"github.com/go-playground/validator/v10"
)

type User struct {
	FullName string `validate:"required,min=3,max=255"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Document string `validate:"required,min=11,max=14,document"`
}

func NewUser(userDTOP *dto.UserDTO) (*User, error) {
	user := &User{
		FullName: userDTOP.FullName,
		Email:    userDTOP.Email,
		Password: userDTOP.Password,
		Document: userDTOP.Document,
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.RegisterValidation("document", CustomValidators.ValidateDocument)

	if err != nil {
		return err
	}

	return validate.Struct(u)
}
