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

func (s *Service) RegisterUser(ctx context.Context, userDTO *UserDTO) (*User, error) {
	user, err := NewUserWithWallet(userDTO)
	if err != nil {
		return nil, err
	}

	if s.repo.FindUserByEmailOrDocument(ctx, user.Email, user.Document) {
		return nil, ErrUserInfoAlreadyRegistered{}
	}
	err = s.repo.CreateUserWithWallet(ctx, user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) FindUserByEmailOrDocument(
	ctx context.Context,
	email string,
	document string,
) bool {
	return s.repo.FindUserByEmailOrDocument(ctx, email, document)
}
