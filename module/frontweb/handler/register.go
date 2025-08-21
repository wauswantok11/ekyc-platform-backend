package handler

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	"git.inet.co.th/ekyc-platform-backend/pkg/util"
)

func (handler Handler) PostRegisterUserHandler(ctx *fiber.Ctx) error {

	var payload dto.RequestRegisterUser

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

	errOpenApi, err := handler.svc.RegisterUserService(ctx, ctx.UserContext(), payload)
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
				Message:    errOpenApi,
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
		Data:       "Register Successfully",
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}
