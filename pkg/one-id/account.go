package one_id

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/bytedance/sonic"
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
		if err := sonic.Unmarshal(responseApi.Body, &ResponseErrorOneId); err != nil {
			logrus.Error("PKG LoginPWD : json.Unmarshal response error body", err)
			return nil, nil, fmt.Errorf("error unmarshalling response error body: %w", err)
		}
		return nil, &ResponseErrorOneId, nil
	}

	if err := sonic.Unmarshal(responseApi.Body, &ResponseSuccessAccount); err != nil {
		logrus.Error("PKG LoginPWD : json.Unmarshal response success body", err)
		return nil, nil, fmt.Errorf("error unmarshalling response success body")
	}
	return &ResponseSuccessAccount, nil, nil
}

func (c *Client) GetAccountProfileAvatarById(ctx context.Context, accountOneId string) (string, error) {
	_, span := c.tracer.Start(ctx, "one_id.api_get_avatar")
	defer span.End()

	urlPath := fmt.Sprintf(`%s/api/get_avatar/%s`, c.url, accountOneId)
	headers := map[string]string{
		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
	}

	responseApi, err := requests.Get(urlPath, headers, nil, int(c.timeOut))
	if err != nil {
		logrus.Errorln("Error connecting to One Id backend:", err.Error())
		return "", errors.New("failed to connect to One Id backend")
	}

	// responseApi.Body น่าจะเป็น []byte อยู่แล้ว
	base64Str := base64.StdEncoding.EncodeToString(responseApi.Body)

	return base64Str, nil
}

func (c *Client) CheckUsernameDup(ctx context.Context, username string) (*ResponseCheckDupUsername, error) {
	urlPath := fmt.Sprintf(`%s/api/check_username?username=%s`, c.url, username)

	_, span := c.tracer.Start(ctx, urlPath)
	defer span.End()
	logrus.Println(urlPath)
	var RespCheckDupUsername ResponseCheckDupUsername

	headers := map[string]string{
		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
	}

	responseApi, err := requests.Get(urlPath, headers, nil, int(c.timeOut))
	if err != nil {
		logrus.Errorln("Error connecting to One Id backend:", err.Error())
		return nil, err
	}

	if err := sonic.Unmarshal(responseApi.Body, &RespCheckDupUsername); err != nil {
		logrus.Error("PKG CheckUsernameDup : json.Unmarshal response success body", err)
		return nil, err
	}

	logrus.Println("ResponseCheckDupUsername : ", RespCheckDupUsername)

	return &RespCheckDupUsername, nil

}

func (c *Client) CheckIdCardDup(ctx context.Context, cid string) (*ResponseCheckDupUsername, error) {
	_, span := c.tracer.Start(ctx, "one_id.check_account_by_cid")
	defer span.End()

	urlPath := fmt.Sprintf(`%s/api/check_id?idCard=%s`, c.url, cid)
	var RespCheckDupUsername ResponseCheckDupUsername

	headers := map[string]string{
		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
	}

	responseApi, err := requests.Get(urlPath, headers, nil, int(c.timeOut))
	if err != nil {
		logrus.Errorln("Error connecting to One Id backend:", err.Error())
		return nil, err
	}

	if err := sonic.Unmarshal(responseApi.Body, &RespCheckDupUsername); err != nil {
		logrus.Error("PKG CheckUsernameDup : json.Unmarshal response success body", err)
		return nil, err
	}

	return &RespCheckDupUsername, nil
}

func (c *Client) CheckEmailDup(ctx context.Context, email string) (*ResponseCheckDupEmail, error) {
	_, span := c.tracer.Start(ctx, "one_id.check_account_by_email")
	defer span.End()

	urlPath := fmt.Sprintf(`%s/api/check_email?email=%s`, c.url, email)
	var ResponseCheckDupEmail ResponseCheckDupEmail

	headers := map[string]string{
		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
	}

	responseApi, err := requests.Get(urlPath, headers, nil, int(c.timeOut))
	if err != nil {
		logrus.Errorln("Error connecting to One Id backend:", err.Error())
		return nil, err
	}

	if err := sonic.Unmarshal(responseApi.Body, &ResponseCheckDupEmail); err != nil {
		logrus.Error("PKG CheckUsernameDup : json.Unmarshal response success body", err)
		return nil, err
	}

	return &ResponseCheckDupEmail, nil
}

// func (c *Client) CheckAccountByCid(ctx context.Context, cid string) (bool, error) {
// 	_, span := c.tracer.Start(ctx, "one_id.check_account_by_cid")
// 	defer span.End()

// 	urlPath := fmt.Sprintf(`%s/api/check_id?idCard=%s`, c.url, cid)
// 	headers := map[string]string{
// 		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
// 	}
// 	_, err := requests.Get(urlPath, headers, nil, int(c.timeOut))
// 	if err != nil {
// 		if err != fiber.ErrBadRequest {
// 			logrus.Errorln("Error connecting to One Id backend:", err.Error())
// 			return false, errors.New("failed to connect to One Id backend")
// 		}
// 		return false, nil
// 	}
// 	return true, nil
// }

// func (c *Client) CheckAccountByEmail(ctx context.Context, email string) (bool, error) {
// 	_, span := c.tracer.Start(ctx, "one_id.check_account_by_email")
// 	defer span.End()

// 	urlPath := fmt.Sprintf(`%s/api/check_id?idCard=%s`, c.url, email)
// 	headers := map[string]string{
// 		fiber.HeaderContentType: fiber.MIMEApplicationJSON,
// 	}
// 	_, err := requests.Get(urlPath, headers, nil, int(c.timeOut))
// 	if err != nil {
// 		if err != fiber.ErrBadRequest {
// 			logrus.Errorln("Error connecting to One Id backend:", err.Error())
// 			return false, errors.New("failed to connect to One Id backend")
// 		}
// 		return false, nil
// 	}
// 	return true, nil
// }
