package util

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Result       string      `json:"result"`
	Data         interface{} `json:"data"`
	Error        string      `json:"errorMessage,omitempty"`
	ResponseCode int         `json:"responseCode,omitempty"`
	Code         int         `json:"code,omitempty"`
}

func GetHttpRequestId(ctx context.Context) string {
	requestId, ok := ctx.Value("requestid").(string)
	if ok {
		return requestId
	}
	return ""
}

type HttpSkipper struct {
	Rule map[string]struct{}
}

func NewHttpSkipper() *HttpSkipper {
	return &HttpSkipper{Rule: map[string]struct{}{}}
}

func (s *HttpSkipper) Add(m string, p string) {
	s.Rule[fmt.Sprintf("%s|%s", m, p)] = struct{}{}
}

func (s *HttpSkipper) Has(m string, p string) bool {
	if _, ok := s.Rule[fmt.Sprintf("%s|%s", m, p)]; ok {
		return ok
	}
	return false
}

func HttpError(ctx *fiber.Ctx, statusCode int, result string, error string) error {
	return ctx.Status(statusCode).JSON(ApiResponse{
		Result:       result,
		Data:         nil,
		Error:        error,
		ResponseCode: statusCode,
	})
}

// func HttpSuccess(ctx *fiber.Ctx, statusCode int, data interface{}) error {
// 	return ctx.Status(statusCode).JSON(ApiResponse{
// 		Result: "Success",
// 		Data:   data,
// 		Code:   statusCode,
// 	})
// }
