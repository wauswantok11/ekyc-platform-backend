package handler

import (
	"net/http"
	"strings"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetCheckUsernameHandler(ctx *fiber.Ctx) error {
	var payload dto.RequestUsername

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
		return ctx.Status(http.StatusBadRequest).JSON(dto.ApiResponse{
			Status:     "failed",
			Data:       nil,
			Message:    strings.TrimSpace(trimmed),
			StatusCode: http.StatusBadRequest,
		})
	}

	response, errOpenApi, err := h.svc.PostCheckUsernameDupService(ctx.Context(), payload.Username)
	if err != nil {
		if err.Error() == "error one" {
			return ctx.Status(http.StatusServiceUnavailable).JSON(dto.ApiResponse{
				Status:     "failed",
				Data:       "Service Unavailable",
				Message:    errOpenApi,
				StatusCode: http.StatusServiceUnavailable,
			})
		}
	}

	if response == "username duplicate" {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ApiResponse{
			Status:     "failed",
			Data:       "username duplicate",
			Message:    "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
	}

	return ctx.Status(http.StatusOK).JSON(dto.ApiResponse{
		Status:     "success",
		Data:       "username not found",
		Message:    "OK",
		StatusCode: http.StatusOK,
	})

}
