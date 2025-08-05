package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
)

func (h Handler) GetUserProfile(ctx *fiber.Ctx) error {
	var req dto.RequestUserProfile
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ApiResponse{
			Status:     "failed",
			Data:       nil,
			Message:    "invalid request",
			StatusCode: http.StatusBadRequest,
		})
	}
 
	return ctx.Status(http.StatusOK).JSON(dto.ApiResponse{
		Status:     "success",
		Data:       "",
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}
