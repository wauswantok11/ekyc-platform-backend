package one_id

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"git.inet.co.th/ekyc-platform-backend/pkg/requests"
)

func (c *Client) LoginPWD(ctx context.Context, username, password string) (*ResponseSuccessLoginPWD, *ResponseErrorOneId, error) {
	_, span := c.tracer.Start(ctx, "one_id.LoginPWD")
	defer span.End()

	var ResponseSuccessLoginPWD ResponseSuccessLoginPWD
	var ResponseErrorOneId ResponseErrorOneId

	urlPath := fmt.Sprintf(`%s/api/oauth/getpwd`, c.url)
	headers := map[string]string{
		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
	}

	body, err := sonic.Marshal(RequestLoginPWD{
		GrantType:    "password",
		ClientId:     c.clientId,
		ClientSecret: c.clientSecret,
		Username:     username,
		Password:     password,
	})
	if err != nil {
		logrus.Errorln("Error marshalling request body:", err.Error())
		return nil, nil, err
	}

	responseApi, err := requests.Post(urlPath, headers, bytes.NewBuffer(body), int(c.timeOut))
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

	if err := sonic.Unmarshal(responseApi.Body, &ResponseSuccessLoginPWD); err != nil {
		logrus.Error("PKG LoginPWD : json.Unmarshal response success body", err)
		return nil, nil, fmt.Errorf("error unmarshalling response success body")
	}
	return &ResponseSuccessLoginPWD, nil, nil
}
