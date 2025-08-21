package services

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (srv Service) PostCheckUsernameDupService(ctx context.Context, username string) (string, string, error) {

	checkDupUsername, err := srv.repo.FindCheckUsernameRepo(ctx, username)
	if err != nil {
		logrus.Error("[*] Error : FindChackUsernameRepo -> ", err.Error())
		return "", "", err
	}
	if checkDupUsername != "username duplicate" {
		// //! check to one-id
		resp, err := srv.repo.OneId().CheckUsernameDup(ctx, username)
		if err != nil {
			logrus.Error("[*] Error Service : One CheckUsernameDup -> ", err.Error())
			return "error one", "err.Error()", err
		}

		msgOpenApi := resp.ErrorMessage
		if msgOpenApi != nil && *msgOpenApi == "username duplicate" {
			return "username duplicate", "", nil
		}
	}
	return "username not found", "", nil

}

func (srv Service) PostCheckCidDupService(ctx context.Context, cid string) (string, string, error) {

	checkDupCid, err := srv.repo.FindCheckCidRepo(ctx, cid)
	if err != nil {
		logrus.Error("[*] Error : FindCheckCidRepo -> ", err.Error())
		return "", "", err
	}

	if checkDupCid != "id duplicate" {
		// //! check to one-id
		resp, err := srv.repo.OneId().CheckIdCardDup(ctx, cid)
		if err != nil {
			logrus.Error("[*] Error Service : One CheckIdCardDup -> ", err.Error())
			return "error one", "err.Error()", err
		}

		msgOpenApi := resp.ErrorMessage
		if msgOpenApi != nil && *msgOpenApi == "id duplicate" {
			return "id duplicate", "", nil
		}
	}
	return "id not found", "", nil

}

func (srv Service) PostCheckEmailDupService(ctx context.Context, email string) (string, string, error) {

	checkDupEmail, err := srv.repo.FindCheckEmailRepo(ctx, email)
	if err != nil {
		logrus.Error("[*] Error : FindCheckEmailRepo -> ", err.Error())
		return "", "", err
	}
	if checkDupEmail != "email duplicate" {
		// //! check to one-id
		resp, err := srv.repo.OneId().CheckEmailDup(ctx, email)
		if err != nil {
			logrus.Error("[*] Error Service : One CheckEmailDup -> ", err.Error())
			return "error one", "err.Error()", err
		}

		msgOpenApi := resp.ErrorMessage
		if msgOpenApi != nil && *msgOpenApi == "email duplicate" {
			return "email duplicate", "", nil
		}
	}
	return "email not fount", "", nil

}
