package models

import "testing"

func TestUserModelValidation(t *testing.T) {
	testCases := []struct {
		name        string
		user        *User
		expectedErr bool
	}{
		{
			name: "Should be valid user",
			user: &User{
				FullName: "Jon Doe",
				Email:    "test@test.com",
				Password: "password",
				Document: "21472605000155",
			},
			expectedErr: false,
		},
		{
			name: "Should return error when user is invalid missing email",
			user: &User{
				FullName: "Jon Doe",
				Email:    "",
				Password: "password",
				Document: "21472605000155",
			},
			expectedErr: true,
		},
		{
			name: "Should return error when user is invalid  email",
			user: &User{
				FullName: "Jon Doe",
				Email:    "invalid_email",
				Password: "password",
				Document: "21472605000155",
			},
			expectedErr: true,
		},

		{
			name: "Should return error when user is invalid missing password",
			user: &User{
				FullName: "Jon Doe",
				Email:    "test@test.com",
				Password: "",
				Document: "21472605000155",
			},
			expectedErr: true,
		},
		{
			name: "Should return error when user is invalid missing document",
			user: &User{
				FullName: "Jon Doe",
				Email:    "test@test.com",
				Password: "password",
				Document: "",
			},
			expectedErr: true,
		},
		{
			name: "Should return error when user is invalid document cpf",
			user: &User{
				FullName: "Jon Doe",
				Email:    "test@test.com",
				Password: "password",
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
