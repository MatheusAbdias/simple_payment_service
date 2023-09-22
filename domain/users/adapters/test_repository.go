package adapters

import (
	"context"
	"strings"

	domain "github.com/MatheusAbdias/simple_payment_service/domain/users"
)

type TestRepository struct {
	db []*domain.User
}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}
func (r *TestRepository) CreateUserWithWallet(ctx context.Context, user *domain.User) error {
	r.db = append(r.db, user)
	return nil
}

func (r *TestRepository) FindUserByEmailOrDocument(
	ctx context.Context,
	email string,
	document string,
) bool {

	email = strings.ToLower(email)
	for _, user := range r.db {
		userEmail := strings.ToLower(user.Email)
		return userEmail == email || user.Document == document
	}

	return false
}
