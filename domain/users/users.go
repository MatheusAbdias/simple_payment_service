package users

import (
	Validators "github.com/MatheusAbdias/simple_payment_service/domain/users/validators"
	wallets "github.com/MatheusAbdias/simple_payment_service/domain/wallets"
	utils "github.com/MatheusAbdias/simple_payment_service/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id       string `validate:"required,uuid4"`
	FullName string `validate:"required,min=3,max=255"`
	Email    string `validate:"required,email"`
	Document string `validate:"required,min=11,max=14,document"`
	Wallet   *wallets.Wallet
}

func NewUserWithWallet(userDTO *UserDTO) (*User, error) {
	var err error
	var wallet *wallets.Wallet
	userID := utils.NewUUID()

	if wallet, err = wallets.NewWallet(userID); err != nil {
		return nil, err
	}

	user := &User{
		Id:       userID,
		FullName: userDTO.FullName,
		Email:    userDTO.Email,
		Document: userDTO.Document,
		Wallet:   wallet,
	}
	if err = user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.RegisterValidation("document", Validators.ValidateDocument)

	if err != nil {
		return err
	}

	return validate.Struct(u)
}
