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

func (repository UserRepository) CreateUser(user *stores.User, userActivate *stores.UserActivation) (*stores.User, error) {
	if err := repository.DB.Create(&user).Error; err != nil {
		return &stores.User{}, err
	}

	userActivate.UserId = user.ID

	if err := repository.DB.Create(&userActivate).Error; err != nil {
		return &stores.User{}, err
	}

	return user, nil
}

func (repository UserRepository) FindUserByEmail(email string) (*stores.User, error) {
	var user *stores.User

	if err := repository.DB.First(&user, "email = ?", email).Error; err != nil {
		return &stores.User{}, err
	}

	return user, nil
}

func (repository UserRepository) FindUserById(id string) (*stores.User, error) {
	var user *stores.User

	if err := repository.DB.First(&user, "id = ?", id).Error; err != nil {
		return &stores.User{}, err
	}

	return user, nil
}

func (repository UserRepository) FindUserActivationCode(userId string, code string) (*stores.UserActivation, error) {
	var act *stores.UserActivation

	if err := repository.DB.First(&act, "user_id = ? AND code = ?", userId, code).Error; err != nil {
		return &stores.UserActivation{}, err
	}

	return act, nil
}

func (repository UserRepository) UpdateUserActivation(id string, stat bool) (*stores.User, error) {
	var user *stores.User

	if err := repository.DB.First(&user, "id = ?", id).Error; err != nil {
		return &stores.User{}, err
	}

	user.IsActive = stat

	if err := repository.DB.Save(&user).Error; err != nil {
		return &stores.User{}, err
	}

	return user, nil
}

func (repository UserRepository) ReCreateUserActivation(userActivate *stores.UserActivation) (*stores.UserActivation, error) {

	if err := repository.DB.Create(&userActivate).Error; err != nil {
		return &stores.UserActivation{}, err
	}

	return userActivate, nil
}
