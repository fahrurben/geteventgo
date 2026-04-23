package users

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Create(ctx context.Context, user *UserModel) error {
	return gorm.G[UserModel](repo.db).Create(ctx, user)
}

func (repo *UserRepository) GetById(ctx context.Context, id uint) (UserModel, error) {
	return gorm.G[UserModel](repo.db).Where("id = ?", id).First(ctx)
}

func (repo *UserRepository) GetByEmail(ctx context.Context, email string) (UserModel, error) {
	return gorm.G[UserModel](repo.db).Where("email = ?", email).First(ctx)
}
