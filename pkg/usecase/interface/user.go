package useCaseInterface

import (
	"context"
	"mime/multipart"
	"usermanagement/pkg/domain"
)

type UserUserCase interface {
	GetAllUsers(ctx context.Context) ([]domain.Users, error)
	GetSingleUser(ctx context.Context, id uint) (domain.Users, error)
	CreateUser(ctx context.Context, user domain.Users, file *multipart.FileHeader) (domain.Users, error)
	DeleteUser(ctx context.Context, id uint) (domain.Users, error)
	UpdateUser(ctx context.Context, user domain.Users, file *multipart.FileHeader) (domain.Users, error)
}
