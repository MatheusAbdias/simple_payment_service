package adapters

import (
	"context"

	domain "github.com/MatheusAbdias/simple_payment_service/domain/users"
)

type TestRepository struct {
	repo []*domain.User
}

func (testRepo *TestRepository) CreateUser(ctx context.Context, user *domain.User) error {
	testRepo.repo = append(testRepo.repo, user)
	return nil
}

func (testRepo *TestRepository) FindUserByEmail(ctx context.Context, email string) bool {
	for _, user := range testRepo.repo {
		if user.Email == email {
			return true
		}
	}
	return false
}

func (testRepo *TestRepository) FindUserByDocument(ctx context.Context, document string) bool {
	for _, user := range testRepo.repo {
		if user.Document == document {
			return true
		}
	}
	return false
}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}
