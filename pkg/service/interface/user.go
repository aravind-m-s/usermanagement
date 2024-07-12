package serviceInterfaces

import (
	"context"
	"mime/multipart"
)

type UserService interface {
	SaveImage(ctx context.Context, userId string, file multipart.FileHeader) (string, error)
	DeleteImage(ctx context.Context, userId string, filePath string) (string, error)
	UpdateImage(ctx context.Context, userId string, file multipart.FileHeader) (string, error)
}
