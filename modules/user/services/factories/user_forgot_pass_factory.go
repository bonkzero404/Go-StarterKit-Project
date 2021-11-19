package factories

import (
	respModel "go-boilerplate-clean-arch/domain/models"
	"go-boilerplate-clean-arch/domain/stores"
	"go-boilerplate-clean-arch/modules/user/domain/interfaces"
	"go-boilerplate-clean-arch/utils"

	"github.com/gofiber/fiber/v2"
)

type UserForgotPassServiceFactory struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewUserForgotPassServiceFactory(userRepository interfaces.UserRepositoryInterface) interfaces.UserForgotPassServiceFactoryInterface {
	return &UserForgotPassServiceFactory{
		UserRepository: userRepository,
	}
}

func (service UserForgotPassServiceFactory) CreateUserForgotPass(user *stores.User) (*stores.UserActivation, error) {
	codeGen := utils.StringWithCharset(32)

	userActivate := stores.UserActivation{
		UserId:  user.ID,
		Code:    codeGen,
		ActType: stores.FORGOT_PASSWORD,
	}

	userAct, errRecreate := service.UserRepository.CreateUserActivation(&userActivate)

	if errRecreate != nil {
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

	return userAct, nil
}
