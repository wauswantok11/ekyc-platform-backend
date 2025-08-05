package one_id

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	// "github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"git.inet.co.th/ekyc-platform-backend/pkg/requests"
)

func (c *Client) GetAccountByToken(ctx context.Context, token string) (*ResponseApiAccountOneId, *ResponseErrorOneId, error) {
	_, span := c.tracer.Start(ctx, "one_id.LoginPWD")
	defer span.End()

	var ResponseSuccessAccount ResponseApiAccountOneId
	var ResponseErrorOneId ResponseErrorOneId

	urlPath := fmt.Sprintf(`%s/api/account`, c.url)
	headers := map[string]string{
		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
		"authorization":         "Bearer " + token,
	}

	responseApi, err := requests.Get(urlPath, headers, nil, int(c.timeOut))
	if err != nil {
		logrus.Errorln("Error connecting to One Id backend:", err.Error())
		return nil, nil, errors.New("failed to connect to One Id backend")
	}

	if responseApi.Code != fiber.StatusOK {
		if err := json.Unmarshal(responseApi.Body, &ResponseErrorOneId); err != nil {
			logrus.Error("PKG LoginPWD : json.Unmarshal response error body", err)
			return nil, nil, fmt.Errorf("error unmarshalling response error body: %w", err)
		}
		return nil, &ResponseErrorOneId, nil
	}

	if err := json.Unmarshal(responseApi.Body, &ResponseSuccessAccount); err != nil {
		logrus.Error("PKG LoginPWD : json.Unmarshal response success body", err)
		return nil, nil, fmt.Errorf("error unmarshalling response success body")
	}
	return &ResponseSuccessAccount, nil, nil
}
