package services

import (
	"go-boilerplate-clean-arch/domain/stores"
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/modules/user/domain/models"
	"go-boilerplate-clean-arch/utils"
)

type UserService struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewUserService(userRepository interfaces.UserRepositoryInterface) interfaces.UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (service UserService) CreateUser(user *models.UserCreateRequest) (*models.UserCreateResponse, error) {
	hashPassword, _ := utils.HashPassword(user.Password)

	userData := stores.User{
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: hashPassword,
	}

	result, err := service.UserRepository.CreateUser(&userData)

	if err != nil {
		return &models.UserCreateResponse{}, err
	}

	response := models.UserCreateResponse{
		ID:       userData.ID.String(),
		FullName: result.FullName,
		Email:    result.Email,
		Phone:    result.Phone,
		IsActive: userData.IsActive,
	}

	return &response, nil
}
