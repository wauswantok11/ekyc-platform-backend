package repositories

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
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
