package wallet

import (
	utils "github.com/MatheusAbdias/simple_payment_service/pkg/utils"
	"github.com/go-playground/validator/v10"
)

type Wallet struct {
	Id      string `validate:"required,uuid4"`
	OwnerID string `validate:"required,uuid4"`
	Amount  int32  `validate:"gte=0"`
}

func NewWallet(ownerID string) (*Wallet, error) {
	wallet := &Wallet{
		Id:      utils.NewUUID(),
		OwnerID: ownerID,
	}

	if err := wallet.Validate(); err != nil {
		return nil, err
	}

	return wallet, nil
}

func (w *Wallet) Validate() error {
	validator := validator.New(validator.WithRequiredStructEnabled())
	return validator.Struct(w)
}
