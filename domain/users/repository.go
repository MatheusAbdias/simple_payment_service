package users

import "context"

type UserRepository interface {
	CreateUserWithWallet(ctx context.Context, user *User) error
	FindUserByEmailOrDocument(ctx context.Context, email string, document string) bool
}
