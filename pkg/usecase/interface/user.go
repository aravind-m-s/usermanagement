package useCaseInterface

import (
	"context"
	"usermanagement/pkg/domain"
)

type UserUserCase interface {
	GetAllUsers(ctx context.Context) ([]domain.Users, error)
	GetSingleUser(ctx context.Context) (domain.Users, error)
	CreateUser(ctx context.Context) (domain.Users, error)
	DeleteUser(ctx context.Context) (domain.Users, error)
	UpdateUser(ctx context.Context) (domain.Users, error)
}
