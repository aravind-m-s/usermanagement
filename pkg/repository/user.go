package repository

import (
	"context"
	"usermanagement/pkg/domain"
	interfaces "usermanagement/pkg/repository/interface"

	"gorm.io/gorm"
)

type userDB struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDB{DB}
}

// CreateUser implements interfaces.UserRepository.
func (u *userDB) CreateUser(ctx context.Context) (domain.Users, error) {
	panic("unimplemented")
}

// DeleteUser implements interfaces.UserRepository.
func (u *userDB) DeleteUser(ctx context.Context) (domain.Users, error) {
	panic("unimplemented")
}

// GetAllUsers implements interfaces.UserRepository.
func (u *userDB) GetAllUsers(ctx context.Context) ([]domain.Users, error) {
	panic("unimplemented")
}

// GetSingleUser implements interfaces.UserRepository.
func (u *userDB) GetSingleUser(ctx context.Context) (domain.Users, error) {
	panic("unimplemented")
}

// UpdateUser implements interfaces.UserRepository.
func (u *userDB) UpdateUser(ctx context.Context) (domain.Users, error) {
	panic("unimplemented")
}
