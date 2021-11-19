package factories

import (
	respModel "go-boilerplate-clean-arch/domain/models"
	"go-boilerplate-clean-arch/domain/stores"
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
)

type UserActivationServiceFactory struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewUserActivationServiceFactory(userRepository interfaces.UserRepositoryInterface) interfaces.UserActivationServiceFactoryInterface {
	return &UserActivationServiceFactory{
		UserRepository: userRepository,
	}
}

func (service UserActivationServiceFactory) CreateUserActivation(user *stores.User) (*stores.UserActivation, error) {
	codeGen := utils.StringWithCharset(32)

	userActivate := stores.UserActivation{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.ACTIVATION_CODE,
	}

	userAct, errRecreate := service.UserRepository.CreateUserActivation(&userActivate)

	if errRecreate != nil {
		return &stores.UserActivation{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Failed to re create activation user, please try again",
		}
	}

	sendMail := respModel.Mail{
		To:           []string{user.Email},
		Subject:      "User Activation",
		TemplateHtml: "user_activation.html",
		BodyParam: map[string]interface{}{
			"Name": user.FullName,
			"Code": codeGen,
		},
	}

	utils.SendMail(&sendMail)

	return userAct, nil
}
