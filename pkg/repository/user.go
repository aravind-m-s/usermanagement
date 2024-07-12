package repository

import (
	"context"
	"usermanagement/pkg/domain"
	interfaces "usermanagement/pkg/repository/interface"
	serviceInterfaces "usermanagement/pkg/service/interface"

	"gorm.io/gorm"
)

type userDB struct {
	DB      *gorm.DB
	Service serviceInterfaces.UserService
}

func NewUserRepository(DB *gorm.DB, service serviceInterfaces.UserService) interfaces.UserRepository {
	return &userDB{DB, service}
}

func (u *userDB) CreateUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	panic("unimplemented")
}

func (u *userDB) DeleteUser(ctx context.Context, id uint) (domain.Users, error) {
	panic("unimplemented")
}

func (u *userDB) GetAllUsers(ctx context.Context) ([]domain.Users, error) {
	panic("unimplemented")
}

func (u *userDB) GetSingleUser(ctx context.Context, id uint) (domain.Users, error) {
	panic("unimplemented")
}

func (u *userDB) UpdateUser(ctx context.Context, user domain.Users) (domain.Users, error) {
	panic("unimplemented")
}
