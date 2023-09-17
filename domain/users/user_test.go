package users_test

import (
	"testing"

	"github.com/MatheusAbdias/simple_payment_service/domain/users"

	utils "github.com/MatheusAbdias/simple_payment_service/pkg/utils"
)

func TestUserModelValidation(t *testing.T) {
	testCases := []struct {
		name        string
		user        *users.User
		expectedErr bool
	}{
		{
			name: "Should be valid user",
			user: &users.User{
				Id:       utils.NewUUID(),
				FullName: "Jon Doe",
				Email:    "test@test.com",
				Document: "21472605000155",
			},
			expectedErr: false,
		},
		{
			name: "Should return error when user is invalid missing email",
			user: &users.User{
				Id:       utils.NewUUID(),
				FullName: "Jon Doe",
				Email:    "",
				Document: "21472605000155",
			},
			expectedErr: true,
		},
		{
			name: "Should return error when user is invalid  email",
			user: &users.User{
				Id:       utils.NewUUID(),
				FullName: "Jon Doe",
				Email:    "invalid_email",
				Document: "21472605000155",
			},
			expectedErr: true,
		},

		{
			name: "Should return error when user is invalid  uuid",
			user: &users.User{
				Id:       "123",
				FullName: "Jon Doe",
				Email:    "test@test.com",
				Document: "21472605000155",
			},
			expectedErr: true,
		},
		{
			name: "Should return error when user is invalid missing document",
			user: &users.User{
				Id:       utils.NewUUID(),
				FullName: "Jon Doe",
				Email:    "test@test.com",
				Document: "",
			},
			expectedErr: true,
		},
		{
			name: "Should return error when user is invalid document cpf",
			user: &users.User{
				Id:       utils.NewUUID(),
				FullName: "Jon Doe",
				Email:    "test@test.com",
				Document: "11111111111",
			},
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.user.Validate()

			if err != nil && !tc.expectedErr {
				t.Errorf("Expected error to be nil, got %v", err)
			}

			if err == nil && tc.expectedErr {
				t.Errorf("Expected error to be not nil, got %v", err)
			}
		})
	}
}
