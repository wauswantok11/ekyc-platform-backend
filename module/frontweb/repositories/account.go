package repositories

import (
	"context"
	"errors"

	"git.inet.co.th/ekyc-platform-backend/model"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/mapper"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

func (r Repository) GetAccountByAccountIdOneRepo(ctx context.Context, accountIdOne string) (string, error) {
	_, span := r.Trace(ctx, "GetAccountByAccountIdOneRepo", oteltrace.WithAttributes(
		attribute.String("AccountIdOne", accountIdOne),
	))
	defer span.End()
	// r.app.Config.Secret.JwtKey
	// Example of correct usage

	return "", nil
}

func (r Repository) FindUserByAccountIdRepo(ctx context.Context, accountId string) (*string, error) {
	var account model.Account
	err := r.dbMain.Ctx().WithContext(ctx).
		Model(&model.Account{}).
		Where("account_one_id = ?", accountId).
		First(&account).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Error("Unexpected DB error:", err)
			return nil, err
		}
		logrus.Warn("Account not found")
		return nil, nil
	}
	logrus.Infof("Found account: %+v", account)
	Id := account.Id.String()
	return &Id, nil
}

func (r Repository) CreateUserRepo(ctx context.Context, userProfile map[string]interface{}) error {
	newAccount, errMapper := mapper.MapToAccount(userProfile)
	if errMapper != nil {
		logrus.Error("Mapping error: ", errMapper)
		return errMapper
	}

	err := r.dbMain.Ctx().WithContext(ctx).Create(newAccount).Error
	if err != nil {
		logrus.Error("Failed to create account: ", err)
		return err
	}
	return nil
}

func (r Repository) UpdateUserRepo(ctx context.Context, userProfile map[string]interface{}, id *string) error {
	if id == nil {
		return errors.New("id is invalid")
	}
	updatedAccount, errMapper := mapper.MapToAccount(userProfile)
	if errMapper != nil {
		logrus.Error("Mapping error: ", errMapper)
		return errMapper
	}

	// Perform the update
	err := r.dbMain.Ctx().WithContext(ctx).
		Model(&model.Account{}).
		Where("id = ?", *id).
		Updates(updatedAccount).Error

	if err != nil {
		logrus.Error("Failed to update account: ", err)
		return err
	}
	return nil
}
