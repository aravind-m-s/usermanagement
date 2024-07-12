package interfaces

import (
	"context"
	"usermanagement/pkg/domain"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]domain.Users, error)
	GetSingleUser(ctx context.Context, id uint) (domain.Users, error)
	CreateUser(ctx context.Context, user domain.Users) (domain.Users, error)
	DeleteUser(ctx context.Context, id uint) (domain.Users, error)
	UpdateUser(ctx context.Context, user domain.Users) (domain.Users, error)
}
