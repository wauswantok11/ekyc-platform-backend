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
	_, span := c.tracer.Start(ctx, "one_id.api/oauth/getpwd")
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

func (c *Client) LoginMobileGetOTP(ctx context.Context, mobileNo string) (ResponseLoginMobileOTP, ResponseErrorOneId, error) {

	_, span := c.tracer.Start(ctx, "one_id.api/oauth/otp")
	defer span.End()

	respSuccess := ResponseLoginMobileOTP{}
	respError := ResponseErrorOneId{}

	urlPath := fmt.Sprintf(`%s/api/oauth/otp`, c.url)
	headers := map[string]string{
		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
	}

	// logrus.Println(" c.refCode ",  c.refCode)

	body, err := sonic.Marshal(RequestLoginMobileOTP{
		ClientId:     c.clientId,
		ClientSecret: c.clientSecret,
		Refcode:      c.refCode,
		MobileNo:     mobileNo,
	})

	if err != nil {
		logrus.Errorln("LoginMobileGetOTP : Error marshalling request body:", err.Error())
		return respSuccess, respError, err
	}

	responseApi, err := requests.Post(urlPath, headers, bytes.NewBuffer(body), int(c.timeOut))
	logrus.Println("err ", err)
	if err != nil {
		logrus.Errorln("LoginMobileGetOTP : Error connecting to One Id backend:", err.Error())
		respError.ErrorMessage = "failed to connect to One Id backend"
		respError.Data = ""
		respError.ResponseCode = 503
		respError.Result = ""
		return respSuccess, respError, err
	}

	if responseApi.Code != fiber.StatusOK {
		if err := json.Unmarshal(responseApi.Body, &respError); err != nil {
			logrus.Error("PKG LoginMobileGetOTP :  json.Unmarshal response error body", err)
			return respSuccess, respError, err
		}
		return respSuccess, respError, errors.New("error invalid")
	}

	logrus.Println("Body ", string(responseApi.Body))

	if err := sonic.Unmarshal(responseApi.Body, &respSuccess); err != nil {
		logrus.Error("PKG LoginMobileGetOTP :  json.Unmarshal response success body", err)
		return respSuccess, respError, err
	}

	logrus.Println("respSuccess :", respSuccess)
	logrus.Println("respError :", respError)
	return respSuccess, respError, nil
}

