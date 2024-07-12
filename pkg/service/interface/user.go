package interfaces

import (
	"context"
)

type UserService interface {
	SaveImage(ctx context.Context) (string, error)
	DeleteImage(ctx context.Context) (string, error)
	UpdateImage(ctx context.Context) (string, error)
}
