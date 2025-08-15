package handler

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	"git.inet.co.th/ekyc-platform-backend/pkg/util"
)

// PostLoginUser

func (handler Handler) PostLoginUserHandler(ctx *fiber.Ctx) error {
	var payload dto.RequestLoginUser

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ApiResponse{
			Status:     "failed",
			Data:       nil,
			Message:    "invalid request",
			StatusCode: http.StatusBadRequest,
		})
	}

	var validate = validator.New()
	if err := validate.Struct(payload); err != nil {
		parts := strings.Split(err.Error(), "Error:")
		trimmed := strings.TrimPrefix(parts[1], "Field validation for ")
		return util.HttpError(ctx, http.StatusBadRequest, "Fail", strings.TrimSpace(trimmed))
	}

	responseLogin, errOpenApiOne, err := handler.svc.LoginUserOneService(ctx, ctx.UserContext(), payload)
	if err != nil {
		if err.Error() == "The user credentials were incorrect." {
			return ctx.Status(http.StatusBadRequest).JSON(dto.ApiResponse{
				Status:     "failed",
				Data:       "Bad Request",
				Message:    err.Error(),
				StatusCode: http.StatusBadRequest,
			})
		} else if err.Error() == "error one" {
			return ctx.Status(http.StatusServiceUnavailable).JSON(dto.ApiResponse{
				Status:     "failed",
				Data:       "Service Unavailable",
				Message:    errOpenApiOne,
				StatusCode: http.StatusServiceUnavailable,
			})
		}

		return ctx.Status(http.StatusInternalServerError).JSON(dto.ApiResponse{
			Status:     "failed",
			Data:       "Internal Server Error",
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.ApiResponse{
		Status:     "success",
		Data:       responseLogin,
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}

func (handler Handler) PostLogoutUserHandler(ctx *fiber.Ctx) error {
	// ดึง JWT หรือ session ID จาก cookie
	accountID, _ := ctx.Locals("account_id").(string)

	if err := handler.svc.LogoutUserService(ctx, ctx.Context(), "authentication", accountID); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.ApiResponse{
			Status:     "internal server error",
			Data:       "",
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.ApiResponse{
		Status:     "success",
		Data:       "",
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}
