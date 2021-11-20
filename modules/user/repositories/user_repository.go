package repositories

import (
	"go-boilerplate-clean-arch/database/driver"
	"go-boilerplate-clean-arch/domain/stores"
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() interfaces.UserRepositoryInterface {
	return &UserRepository{
		DB: driver.DB,
	}
}

func (repository UserRepository) CreateUser(user *stores.User) *gorm.DB {
	return repository.DB.Create(&user)
}

func (repository UserRepository) UpdateUserIsActive(user *stores.User) *gorm.DB {
	return repository.DB.Save(&user)
}

func (repository UserRepository) FindUserByEmail(user *stores.User, email string) *gorm.DB {
	return repository.DB.First(&user, "email = ?", email)
}

func (repository UserRepository) FindUserById(user *stores.User, id string) *gorm.DB {
	return repository.DB.First(&user, "id = ?", id)
}

func (repository UserRepository) UpdatePassword(user *stores.User) *gorm.DB {
	return repository.DB.Save(&user)
}
