package service

import (
	"context"
	"mime/multipart"
	serviceInterfaces "usermanagement/pkg/service/interface"
)

type UserService struct{}

func NewUserService() serviceInterfaces.UserService {
	return &UserService{}
}

func (u *UserService) DeleteImage(ctx context.Context, userId string, filePath string) (string, error) {
	panic("unimplemented")
}

func (u *UserService) SaveImage(ctx context.Context, userId string, file multipart.FileHeader) (string, error) {
	panic("unimplemented")
}

func (u *UserService) UpdateImage(ctx context.Context, userId string, file multipart.FileHeader) (string, error) {
	panic("unimplemented")
}
