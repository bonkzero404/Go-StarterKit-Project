package repositories

import (
	"go-boilerplate-clean-arch/domain/stores"
	"go-boilerplate-clean-arch/infrastructure/database"
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() interfaces.UserRepositoryInterface {
	return &UserRepository{
		DB: database.DB,
	}
}

func (repository UserRepository) CreateUser(user *stores.User) (*stores.User, error) {
	if err := repository.DB.Create(&user).Error; err != nil {
		return &stores.User{}, err
	}

	return user, nil
}
