package repositories

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"

	"git.inet.co.th/ekyc-platform-backend/pkg/util"
)

func (r Repository) GenJwtTokenRepo(ctx context.Context, dataToken map[string]interface{}) (string, error) {
	accountId, _ := dataToken["account_id"].(string)
	_, span := r.Trace(ctx, "GenJwtTokenRepo", oteltrace.WithAttributes(
		attribute.String("Generate JWT Data Account id One", accountId),
	))
	defer span.End()
	TokenJwt, err := util.GenerateJWT(r.app.Config.Secret.JwtKey, dataToken)
	if err != nil {
		return "", err
	}

	return TokenJwt, nil
}

func (r Repository) SetRedisRepo(ctx context.Context, cKey string, userProfile map[string]interface{}) error {
	var AccountId = userProfile["account_id"].(string)
	_, span := r.Trace(ctx, "SetRedisRepo", oteltrace.WithAttributes(
		attribute.String("Set Redis Token One", AccountId),
	))
	defer span.End()
	timeout := 24 * time.Hour

	if errRedis := r.cache.Set(cKey, userProfile, timeout); errRedis != nil {
		return errRedis

	}
	return nil
}
