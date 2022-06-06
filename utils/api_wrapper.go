package utils

import (
	"go-starterkit-project/domain/dto"

	"github.com/gofiber/fiber/v2"
)

/**
This function is used to wrap a json response globally
*/
func ApiWrapper(ctx *fiber.Ctx, message string, code int, status string, data interface{}) error {
	meta := dto.Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	responseJson := dto.Response{
		Meta: meta,
		Data: data,
	}

	return ctx.Status(code).JSON(responseJson)
}

func ApiErrorValidation(ctx *fiber.Ctx, message string, data interface{}) error {
	return ApiWrapper(ctx, message, fiber.StatusNotAcceptable, "error_validation", data)
}

func ApiUnprocessableEntity(ctx *fiber.Ctx, message string, data interface{}) error {
	return ApiWrapper(ctx, message, fiber.StatusUnprocessableEntity, "error_unprocessable_entity", data)
}

func ApiUnauthorized(ctx *fiber.Ctx, message string, data interface{}) error {
	return ApiWrapper(ctx, message, fiber.StatusUnauthorized, "error_unauthorized", data)
}

func ApiCreated(ctx *fiber.Ctx, message string, data interface{}) error {
	return ApiWrapper(ctx, message, fiber.StatusCreated, "success_created", data)
}

func ApiOk(ctx *fiber.Ctx, message string, data interface{}) error {
	return ApiWrapper(ctx, message, fiber.StatusOK, "success_ok", data)
}

func ApiResponseError(ctx *fiber.Ctx, message string, code int, data interface{}) error {
	return ApiWrapper(ctx, message, code, "error_api", data)
}
