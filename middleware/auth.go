package middleware

import (
	"fmt"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	"git.inet.co.th/ekyc-platform-backend/pkg/cache"
	"git.inet.co.th/ekyc-platform-backend/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(store *cache.Redis, jwtSecret string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// ดึง token จาก cookie
		tokenStr := ctx.Cookies("authentication")
		if tokenStr == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
				Status:     "failed",
				Data:       nil,
				Message:    "Missing authentication cookie",
				StatusCode: fiber.StatusUnauthorized,
			})
		}

		// ตรวจสอบ JWT
		claims, err := util.ParseJWT(jwtSecret, tokenStr)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
				Status:     "failed",
				Data:       nil,
				Message:    "Invalid token: " + err.Error(),
				StatusCode: fiber.StatusUnauthorized,
			})
		}

		// ตรวจสอบ account_id
		accountID, ok := claims["account_id"].(string)
		if !ok || accountID == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
				Status:     "failed",
				Data:       nil,
				Message:    "Invalid account_id in token",
				StatusCode: fiber.StatusUnauthorized,
			})
		}

		// ตรวจสอบ Redis ว่า token ยัง valid
		key := fmt.Sprintf("%s_account_token", accountID)
		var redisData map[string]interface{}
		err = store.Get(key, &redisData)
		if err != nil {
			if store.IsKeyNotFound(err) {
				return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
					Status:     "failed",
					Data:       nil,
					Message:    "Token expired or not found",
					StatusCode: fiber.StatusUnauthorized,
				})
			}
			return ctx.Status(fiber.StatusInternalServerError).JSON(dto.ApiResponse{
				Status:     "failed",
				Data:       nil,
				Message:    "Redis error: " + err.Error(),
				StatusCode: fiber.StatusInternalServerError,
			})
		}
 
		ctx.Locals("account_id", accountID) 
		ctx.Locals("token_data", redisData["access_token"])

		return ctx.Next()
	}
}
