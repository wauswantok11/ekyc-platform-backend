package services

import (
	"context"
	"errors"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/mapper"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (srv Service) GetProfileOneIdService(ctx context.Context, accountId, tokenOne string) (*dto.ResponseUserProfile, string, error) {
	var response dto.ResponseUserProfile

	accountDetail, err := srv.repo.FindUserDetailByAccountIdRepo(ctx, accountId)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Error("[*] Error Service : LoginPWD -> record not found")
			return nil, "", err
		}
		responseOne, ResponseErrorOneId, err := srv.repo.OneId().GetAccountByToken(ctx, tokenOne)
		if err != nil {
			logrus.Error("[*] Error Pkg One : GetAccountByToken -> ", err.Error())
			return nil, ResponseErrorOneId.ErrorMessage, errors.New("error one")
		}
		return mapper.MapResponseApiAccountOneIdToResponseUserProfile(*responseOne), "", nil
	}
	logrus.Println(accountDetail)
	//map accountDetail
	return &response, "", nil
}

func (srv Service) GetProfileOneAvatarByAccountOneIdService(ctx context.Context, accountOneId string) (string, error) {
	baseImageAvatar, err := srv.repo.OneId().GetAccountProfileAvatarById(ctx, accountOneId)
	if err != nil {
		logrus.Error("[*] Error Service : LoginPWD -> ", err.Error())
		return "", err
	}
	return baseImageAvatar, nil
}
