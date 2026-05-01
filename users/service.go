package users

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	repo := NewUserRepository(db)
	return &UserService{repo: repo}
}

func (service *UserService) Register(ctx context.Context, data RegisterValidator) error {

	user := UserModel{
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.FirstName,
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.PasswordHash = string(bytes)

	return service.repo.Create(ctx, &user)
}

func (service *UserService) Login(ctx context.Context, data LoginValidator) (*UserModel, error) {

	userModel, err := service.repo.GetByEmail(ctx, data.Email)

	if err != nil {
		return nil, err
	}

	err = userModel.checkPassword(data.Password)

	if err != nil {
		return nil, err
	}

	return &userModel, err
}
