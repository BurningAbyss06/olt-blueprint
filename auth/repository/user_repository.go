package repository

import (
	"context"
	//"database/sql"

	"github.com/metalpoch/olt-blueprint/auth/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (repo userRepository) Create(ctx context.Context, user *entity.User) error {
	if err := repo.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo userRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	users := []*entity.User{}
	if err := repo.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}

func (repo userRepository) GetUserByID(ctx context.Context, id uint) (*entity.User, error) {
	u := new(entity.User)
	if err := repo.db.WithContext(ctx).Find(&u, "id=?", id).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repo userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := new(entity.User)
	if err := repo.db.WithContext(ctx).Find(&u, "email=?", email).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (repo userRepository) SoftDelete(ctx context.Context, id uint) error {
	u := new(entity.User)
	if err := repo.db.Find(&u, "id=?", id).Error; err != nil {
		return err
	}

	if err := repo.db.WithContext(ctx).Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

func (repo userRepository) ChangePassword(ctx context.Context, id uint, password string) error {
	u := new(entity.User)

	if err := repo.db.Find(&u, "id=?", id, password).Error; err != nil {
		return err
	}

	if err := repo.db.Update("Password", password).Error; err != nil {
		return err
	}
	return nil
}
