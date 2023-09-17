package users

import (
	"context"
)

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (service *Service) CreateUser(ctx context.Context, userDTO *UserDTO) (*User, error) {
	user, err := NewUser(userDTO)
	if err != nil {
		return nil, err
	}

	if service.FindUserByEmail(ctx, user.Email) {
		return nil, &ErrEmailAlreadyRegistered{}
	}

	if service.FindUserByDocument(ctx, user.Document) {
		return nil, &ErrDocumentAlreadyRegistered{}
	}

	err = service.repo.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *Service) FindUserByEmail(ctx context.Context, email string) bool {
	return service.repo.FindUserByEmail(ctx, email)
}

func (service *Service) FindUserByDocument(ctx context.Context, document string) bool {
	return service.repo.FindUserByDocument(ctx, document)
}
