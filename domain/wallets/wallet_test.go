package wallet

import (
	"testing"

	utils "github.com/MatheusAbdias/simple_payment_service/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestWalletValidation(t *testing.T) {
	testCases := []struct {
		name        string
		ownerID     string
		expectedErr bool
	}{
		{
			name:        "Should be can create a valid wallet",
			ownerID:     utils.NewUUID(),
			expectedErr: false,
		},
		{
			name:        "Should be cant create a wallet missing owner id",
			ownerID:     "",
			expectedErr: true,
		},
		{
			name:        "Should be cant create a wallet invalid owner id",
			ownerID:     "1",
			expectedErr: true,
		},
	}

	for _, testCases := range testCases {
		t.Run(testCases.name, func(t *testing.T) {
			wallet, err := NewWallet(testCases.ownerID)

			if testCases.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, wallet)
			}
		})
	}
}
