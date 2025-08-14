package services

import (
	"context"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (srv Service) GetProfileOneIdService(ctx context.Context, accountId, tokenOne string) (dto.ResponseUserProfile, string, error) {
	var response dto.ResponseUserProfile

	account, err := srv.repo.FindUserDetailByAccountIdRepo(ctx, accountId)
	if err != nil {
		if !err.Is(err, gorm.ErrRecordNotFound) {
			logrus.Error("[*] Error Service : LoginPWD -> record not found")
			return response, "", err
		}
		responseOne, err := srv.repo.OneId().GetAccountByToken(ctx, tokenOne)
		if err {
			logrus.Error("[*] Error Pkg One : GetAccountByToken -> ", err.Error())
			return "", err
		}

	}

	return response, "", nil
}

func (srv Service) GetProfileOneAvatarByAccountOneIdService(ctx context.Context, accountOneId string) (string, error) {
	baseImageAvatar, err := srv.repo.OneId().GetAccountProfileAvatarById(ctx, accountOneId)
	if err != nil {
		logrus.Error("[*] Error Service : LoginPWD -> ", err.Error())
		return "", err
	}
	return baseImageAvatar, nil
}
