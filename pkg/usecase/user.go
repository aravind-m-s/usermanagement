package usecase

import (
	"context"
	"usermanagement/pkg/domain"
	interfaces "usermanagement/pkg/repository/interface"
	useCaseInterface "usermanagement/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) useCaseInterface.UserUserCase {
	return &userUseCase{repo}
}

// CreateUser implements useCaseInterface.UserUserCase.
func (u *userUseCase) CreateUser(ctx context.Context) (domain.Users, error) {
	panic("unimplemented")
}

// DeleteUser implements useCaseInterface.UserUserCase.
func (u *userUseCase) DeleteUser(ctx context.Context) (domain.Users, error) {
	panic("unimplemented")
}

// GetAllUsers implements useCaseInterface.UserUserCase.
func (u *userUseCase) GetAllUsers(ctx context.Context) ([]domain.Users, error) {
	panic("unimplemented")
}

// GetSingleUser implements useCaseInterface.UserUserCase.
func (u *userUseCase) GetSingleUser(ctx context.Context) (domain.Users, error) {
	panic("unimplemented")
}

// UpdateUser implements useCaseInterface.UserUserCase.
func (u *userUseCase) UpdateUser(ctx context.Context) (domain.Users, error) {
	panic("unimplemented")
}
