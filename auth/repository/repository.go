package repository

import (
	"context"

	"github.com/metalpoch/olt-blueprint/auth/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	SoftDelete(ctx context.Context, id uint) error
	ChangePassword(ctx context.Context, id, password string) error
}
