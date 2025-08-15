package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	"git.inet.co.th/ekyc-platform-backend/pkg/util"
)

// var ctxFiber *fiber.Ctx

func (srv Service) LoginUserOneService(ctxFiber *fiber.Ctx, ctx context.Context, payload dto.RequestLoginUser) (*dto.ResponseLoginUser, string, error) {
	tx := srv.repo.DB().Ctx().Begin()
	if tx.Error != nil {
		logrus.Error("[*] Tx Error : LoginUserOneService -> ", tx.Error)
		return nil, "", tx.Error
	}

	ResponseSuccessLoginPWD, ResponseErrorLoginPWD, err := srv.repo.OneId().LoginPWD(ctx, payload.Username, payload.Password)
	if err != nil {
		logrus.Error("[*] Error Service : LoginPWD -> ", err.Error())
		return nil, "", err
	}

	if ResponseErrorLoginPWD != nil {
		logrus.Error("[*] API Error : SetRedisRepo -> ", ResponseErrorLoginPWD.ErrorMessage)
		return nil, ResponseErrorLoginPWD.ErrorMessage, errors.New("error one")
	}

	var responseToken dto.ResponseLoginUser
	if ResponseSuccessLoginPWD != nil {
		responseToken.AccessToken = ResponseSuccessLoginPWD.AccessToken
		responseToken.AccountId = ResponseSuccessLoginPWD.AccountId
		responseToken.ExpirationDate = ResponseSuccessLoginPWD.ExpirationDate
		responseToken.ExpiresIn = ResponseSuccessLoginPWD.ExpiresIn
		responseToken.RefreshToken = ResponseSuccessLoginPWD.RefreshToken
		responseToken.Result = ResponseSuccessLoginPWD.Result
		responseToken.TokenType = ResponseSuccessLoginPWD.TokenType
		responseToken.Username = ResponseSuccessLoginPWD.Username
		responseToken.LoginBy = ResponseSuccessLoginPWD.LoginBy
	} else {
		return nil, "", errors.New("error one")
	}

	AccountOneId, ErrGetAccountByToken, err := srv.repo.OneId().GetAccountByToken(ctx, responseToken.AccessToken)
	if err != nil {
		logrus.Error("[*] Error Service : GetAccountByToken -> ", err.Error())
		return nil, "", err
	}

	if ErrGetAccountByToken != nil {
		logrus.Error("[*] API Error : ErrGetAccountByToken -> ", ErrGetAccountByToken)
		return nil, ErrGetAccountByToken.ErrorMessage, errors.New("error one")
	}

	//* Set Redis Token
	strDataToken := map[string]interface{}{
		"account_id":      responseToken.AccountId,
		"token_type":      responseToken.TokenType,
		"expires_in":      responseToken.ExpiresIn,
		"access_token":    responseToken.AccessToken,
		"refresh_token":   responseToken.RefreshToken,
		"expiration_date": responseToken.ExpirationDate,
		"result ":         responseToken.Result,
		"username ":       responseToken.Username,
		"login_by ":       responseToken.LoginBy,
	}

	cKeyToken := fmt.Sprintf("%s_account_token", responseToken.AccountId)
	if errRedis := srv.repo.SetRedisRepo(ctx, cKeyToken, strDataToken); errRedis != nil {
		logrus.Error("[*] Error : SetRedisRepo -> ", errRedis.Error())
		return nil, "", errRedis
	}

	strDataProfileDetail := map[string]interface{}{
		"account_id":             AccountOneId.ID,
		"account_title_th":       AccountOneId.AccountTitleTH,
		"special_title_name_th":  AccountOneId.SpecialTitleNameTH,
		"first_name_th":          AccountOneId.FirstNameTH,
		"middle_name_th":         AccountOneId.MiddleNameTH,
		"last_name_th":           AccountOneId.LastNameTH,
		"account_title_eng":      AccountOneId.AccountTitleEng,
		"special_title_name_eng": AccountOneId.SpecialTitleNameEng,
		"first_name_eng":         AccountOneId.FirstNameEng,
		"middle_name_eng":        AccountOneId.MiddleNameEng,
		"last_name_eng":          AccountOneId.LastNameEng,
		"account_category":       AccountOneId.AccountCategory,
		"account_sub_category":   AccountOneId.AccountSubCategory,
		"birth_date":             AccountOneId.BirthDate,
		"hash_id_card_num":       AccountOneId.HashIDCardNum,
		"id_card_num":            AccountOneId.IDCardNum,
		"id_card_type":           AccountOneId.IDCardType,
		"thai_email":             AccountOneId.ThaiEmail,
		"thai_email2":            AccountOneId.ThaiEmail2,
		"thai_email3":            AccountOneId.ThaiEmail3,
		"status_cd":              AccountOneId.StatusCD,
	}

	//* Gen JWT Token
	jwtCode, errJwt := srv.repo.GenJwtTokenRepo(ctx, strDataProfileDetail)
	if errJwt != nil {
		logrus.Error("[*] Error : GenJwtTokenRepo -> ", errJwt.Error())
		return nil, "", errJwt
	}

	//* Set Cookies
	if errCookie := util.SetCookieHandler(ctxFiber, "authentication", jwtCode); errCookie != nil {
		logrus.Error("[*] Error : SetCookieHandler -> ", errCookie.Error())
		return nil, "", errCookie
	}

	//* Set Redis Profile
	cKeyAccount := fmt.Sprintf("%s_account_detail", responseToken.AccountId)
	if errRedis := srv.repo.SetRedisRepo(ctx, cKeyAccount, strDataProfileDetail); errRedis != nil {
		logrus.Error("[*] Error : SetRedisRepo -> ", errRedis.Error())
		return nil, "", errRedis
	}

	//* Find account in Database
	Id, errFindAccount := srv.repo.FindUserByAccountIdRepo(ctx, AccountOneId.ID)
	if errFindAccount != nil {
		logrus.Error("[*] Error : FindUserByAccountIdRepo -> ", errFindAccount.Error())
		return nil, "", errFindAccount
	}

	//* create or update account
	if Id != nil {
		//* update
		errUpdate := srv.repo.UpdateUserRepo(ctx, strDataProfileDetail, Id)
		if errUpdate != nil {
			tx.Rollback()
			logrus.Error("[*] Error : UpdateUserRepo -> ", errUpdate.Error())
			return nil, "", errUpdate
		}
	} else {
		//* create
		errCreate := srv.repo.CreateUserRepo(ctx, strDataProfileDetail)
		if errCreate != nil {
			tx.Rollback()
			logrus.Error("[*] Error : CreateUserRepo -> ", errCreate.Error())
			return nil, "", errCreate
		}
	}

	return &responseToken, "", nil
}
func (srv Service) LogoutUserService(ctxFiber *fiber.Ctx, ctx context.Context, keyCookie, accountId string) error {
	return nil
}
