package wallet

type WalletDTO struct {
	Id      string `json:"id"     required:"false,uuid"`
	OwnerID string `json:"owner"  required:"false,uuid"`
	Amount  int32  `json:"amount" required:"false,numeric"`
}
