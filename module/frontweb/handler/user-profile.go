package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
)

func (handler Handler) GetUserProfile(ctx *fiber.Ctx) error { 
	accountID, ok := ctx.Locals("account_id").(string)
    if !ok || accountID == "" {
        return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
            Status:     "failed",
            Data:       nil,
            Message:    "Unauthorized: account_id not found",
            StatusCode: fiber.StatusUnauthorized,
        })
    }

	AccessToken, ok := ctx.Locals("token_data").(string)
    if !ok || AccessToken == "" {
        return ctx.Status(fiber.StatusUnauthorized).JSON(dto.ApiResponse{
            Status:     "failed",
            Data:       nil,
            Message:    "Unauthorized: AccessToken not found",
            StatusCode: fiber.StatusUnauthorized,
        })
    }

	profileUser, errOpenApiOne, err := handler.svc.GetProfileOneIdService(ctx.UserContext(), accountID, AccessToken)
	if err != nil {
		if err.Error() == "error one" {
			return ctx.Status(http.StatusServiceUnavailable).JSON(dto.ApiResponse{
				Status:     "failed",
				Data:       "service unavailable",
				Message:    errOpenApiOne,
				StatusCode: http.StatusServiceUnavailable,
			})
		}
		return ctx.Status(http.StatusBadRequest).JSON(dto.ApiResponse{
			Status:     "failed",
			Data:       "bad request",
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.ApiResponse{
		Status:     "success",
		Data:       profileUser,
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}

func (handler Handler) GetAvatarUserHandler(ctx *fiber.Ctx) error {
	accountOneId := ctx.Query("accountOneId")

	baseImage, err := handler.svc.GetProfileOneAvatarByAccountOneIdService(ctx.UserContext(), accountOneId)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ApiResponse{
			Status:     "failed",
			Data:       "bad request",
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.ApiResponse{
		Status: "success",
		Data: map[string]interface{}{
			"base_image": baseImage,
		},
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}
