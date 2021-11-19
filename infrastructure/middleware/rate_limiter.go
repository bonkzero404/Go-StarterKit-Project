package middleware

import (
	"go-boilerplate-clean-arch/utils"
	"time"

	respModel "go-boilerplate-clean-arch/domain/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter(max int, duration time.Duration) func(ctx *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		LimitReached: func(ctx *fiber.Ctx) error {
			return utils.ApiResponseError(ctx, "Rate limit", fiber.StatusRequestEntityTooLarge, &respModel.ApiErrorResponse{
				StatusCode: fiber.StatusTooManyRequests,
				Message:    "Request too many request",
			})
		},
		Max:        max,
		Expiration: duration * time.Second,
	})
}