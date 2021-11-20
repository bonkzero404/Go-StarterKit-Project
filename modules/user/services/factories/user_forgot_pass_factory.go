package factories

import (
	respModel "go-boilerplate-clean-arch/domain/models"
	"go-boilerplate-clean-arch/domain/stores"
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
)

type UserForgotPassServiceFactory struct {
	UserActivationRepository interfaces.UserActivationRepositoryInterface
}

func NewUserForgotPassServiceFactory(userActivationRepository interfaces.UserActivationRepositoryInterface) interfaces.UserForgotPassServiceFactoryInterface {
	return &UserForgotPassServiceFactory{
		UserActivationRepository: userActivationRepository,
	}
}

func (service UserForgotPassServiceFactory) CreateUserForgotPass(user *stores.User) (*stores.UserActivation, error) {
	codeGen := utils.StringWithCharset(32)

	userActivate := stores.UserActivation{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.FORGOT_PASSWORD,
	}

	if err := service.UserActivationRepository.CreateUserActivation(&userActivate).Error; err != nil {
		return &stores.UserActivation{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Failed to create forgot password, please try again",
		}
	}

	sendMail := respModel.Mail{
		To:           []string{user.Email},
		Subject:      "Forgot Password",
		TemplateHtml: "user_forgot_password.html",
		BodyParam: map[string]interface{}{
			"Name": user.FullName,
			"Code": codeGen,
		},
	}

	utils.SendMail(&sendMail)

	return &userActivate, nil
}
