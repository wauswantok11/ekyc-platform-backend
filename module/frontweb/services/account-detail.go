package services

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (srv Service) PostCheckUsernameDupService(ctx context.Context, username string) (string, string, error) {

	checkDupUsername, err := srv.repo.FindChackUsernameRepo(ctx, username)
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
