package usecase

import (
	"context"
	"mime/multipart"
	"usermanagement/pkg/domain"
	interfaces "usermanagement/pkg/repository/interface"
	useCaseInterface "usermanagement/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func (u *userUseCase) CreateUser(ctx context.Context, user domain.Users, file *multipart.FileHeader) (domain.Users, error) {
	panic("unimplemented")
}

func (u *userUseCase) DeleteUser(ctx context.Context, id uint) (domain.Users, error) {
	panic("unimplemented")
}

// GetAllUsers implements useCaseInterface.UserUserCase.
func (u *userUseCase) GetAllUsers(ctx context.Context) ([]domain.Users, error) {
	panic("unimplemented")
}

func (u *userUseCase) GetSingleUser(ctx context.Context, id uint) (domain.Users, error) {
	panic("unimplemented")
}

func (u *userUseCase) UpdateUser(ctx context.Context, user domain.Users, file *multipart.FileHeader) (domain.Users, error) {
	panic("unimplemented")
}

func NewUserUseCase(repo interfaces.UserRepository) useCaseInterface.UserUserCase {
	return &userUseCase{repo}
}
