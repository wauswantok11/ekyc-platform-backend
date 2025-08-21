package repositories

import (
	"context"
	"errors"

	"git.inet.co.th/ekyc-platform-backend/model"
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
func (r Repository) FindUserDetailByAccountIdRepo(ctx context.Context, accountId string) (*model.Account, error) {
	_, span := r.Trace(ctx, "FindUserDetailByAccountIdRepo", oteltrace.WithAttributes(
		attribute.String("AccountIdOne", accountId),
	))
	defer span.End()

	var account model.Account
	err := r.dbMain.Ctx().WithContext(ctx).
		Model(&model.Account{}).
		Where("account_one_id = ?", accountId).
		First(&account).Error

	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (r Repository) FindUserByAccountIdRepo(ctx context.Context, accountId string) (*string, error) {
	_, span := r.Trace(ctx, "FindUserByAccountIdRepo", oteltrace.WithAttributes(
		attribute.String("AccountIdOne", accountId),
	))
	defer span.End()

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
	Id := account.Id.String()
	return &Id, nil
}

func (r Repository) CreateUserRepo(ctx context.Context, newAccount model.Account) error {
	err := r.dbMain.Ctx().WithContext(ctx).Create(newAccount).Error
	if err != nil {
		logrus.Error("Failed to create account: ", err)
		return err
	}
	return nil
}

func (r Repository) UpdateUserRepo(ctx context.Context, updatedAccount model.Account, id *string) error {
	if id == nil {
		return errors.New("id is invalid")
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

func (r Repository) CreateOtpManagemontRepo(ctx context.Context, reqStu model.OtpManagement) error {
	_, span := r.Trace(ctx, "CreateOtpManagemontRepo", oteltrace.WithAttributes(
		attribute.String("mobile_no", reqStu.MobileNo),
	))
	defer span.End()

	// if reqStu.Id == (uuid.UUID{}) {
	// 	reqStu.Id = uuid.New()
	// }

	if err := r.dbMain.Ctx().WithContext(ctx).Create(&reqStu).Error; err != nil {
		return err
	}
	return nil
}
