package services

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/mapper"
)

func (srv Service) RegisterUserService(ctxFiber *fiber.Ctx, ctx context.Context, payload dto.RequestRegisterUser) (string, error) {
	tx := srv.repo.DB().Ctx().Begin()
	if tx.Error != nil {
		logrus.Error("[*] Tx Error : RegisterUserService -> ", tx.Error)
		return "", tx.Error
	}
	payloadOne := mapper.MapRequestResgisterToRequestRegisterOneId(&payload)

	responseRegisterOne, ErrRegisterOne, err := srv.repo.OneId().PostRegisterAccount(ctx, *payloadOne)

	if err != nil {
		logrus.Error("[*] Error Service : PostRegisterAccount -> ", err.Error())
		return "", err
	}

	if ErrRegisterOne != nil {
		logrus.Error("[*] API Error : ErrRegisterOne -> ", ErrRegisterOne)
		return ErrRegisterOne.ErrorMessage, errors.New("error one")
	}

	newAccount := mapper.MapRequestResgisterToModelAccount(&payload, responseRegisterOne.Data.AccountID)

	errCreate := srv.repo.CreateUserRepo(ctx, *newAccount)
	if errCreate != nil {
		tx.Rollback()
		logrus.Error("[*] Error : CreateUserRepo -> ", errCreate.Error())
		return "", errCreate
	}

	return "", nil
}
