package utils

import (
	"go-starterkit-project/domain/dto"

	"github.com/gofiber/fiber/v2"
)

/**
This function is used to wrap a json response globally
*/
func ApiWrapper(ctx *fiber.Ctx, code int, status string, data interface{}) error {
	meta := dto.Meta{
		Route:  ctx.Route().Path,
		Method: ctx.Method(),
		Query:  string(ctx.Request().URI().QueryString()),
		Code:   code,
		Status: status,
	}

	// print(ctx.Method())

	if code >= 400 {
		responseJson := dto.Response{
			Valid: false,
			Meta:  meta,
			Error: data,
			Data:  nil,
		}

		return ctx.Status(code).JSON(responseJson)
	}

	responseJson := dto.Response{
		Valid: true,
		Meta:  meta,
		Error: nil,
		Data:  data,
	}

	return ctx.Status(code).JSON(responseJson)
}

func ApiErrorValidation(ctx *fiber.Ctx, data interface{}) error {
	return ApiWrapper(ctx, fiber.StatusNotAcceptable, "error_validation", data)
}

func ApiUnprocessableEntity(ctx *fiber.Ctx, data interface{}) error {
	return ApiWrapper(ctx, fiber.StatusUnprocessableEntity, "error_unprocessable_entity", data)
}

func ApiUnauthorized(ctx *fiber.Ctx, data interface{}) error {
	return ApiWrapper(ctx, fiber.StatusUnauthorized, "error_unauthorized", data)
}

func ApiCreated(ctx *fiber.Ctx, data interface{}) error {
	return ApiWrapper(ctx, fiber.StatusCreated, "success_created", data)
}

func ApiOk(ctx *fiber.Ctx, data interface{}) error {
	return ApiWrapper(ctx, fiber.StatusOK, "success_ok", data)
}

func ApiResponseError(ctx *fiber.Ctx, code int, data interface{}) error {
	return ApiWrapper(ctx, code, "error_api", data)
}
