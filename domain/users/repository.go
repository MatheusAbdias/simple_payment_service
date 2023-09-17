package users

import "context"

type UserRepository interface {
	CreateUser(context.Context, *User) error
	FindUserByEmail(context.Context, string) bool
	FindUserByDocument(context.Context, string) bool
}
