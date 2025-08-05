package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
)

func (h Handler) GetSessionHandler(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(dto.ApiResponse{
		Status:     "success",
		Data:       "",
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}

func (h Handler) GetSessionAdminHandler(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(dto.ApiResponse{
		Status:     "success",
		Data:       "",
		Message:    "OK",
		StatusCode: http.StatusOK,
	})
}
