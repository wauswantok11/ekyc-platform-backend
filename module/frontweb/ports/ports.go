package ports

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"

	"git.inet.co.th/ekyc-platform-backend/config"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	"git.inet.co.th/ekyc-platform-backend/pkg/cache"
	oneId "git.inet.co.th/ekyc-platform-backend/pkg/one-id"
)

type Repository interface {
	Module() string
	AppCfg() *config.Config
	Log() *logrus.Entry
	OneId() *oneId.Client
	Cache() *cache.Redis
	Trace(ctx context.Context, spanName string, attributes ...trace.SpanStartOption) (context.Context, trace.Span)

	GetAccountByAccountIdOneRepo(ctx context.Context, accountIdOne string) (string, error)
	GenJwtTokenRepo(ctx context.Context, dataToken map[string]interface{}) (string, error)
	//* Redis Repo
	// SetRedisRepo(ctx context.Context, cKey, Accesstoken string) error
	SetRedisRepo(ctx context.Context, cKey string, userProfile map[string]interface{}) error
}

type Service interface {
	//* Login Account One Id
	LoginUserOneService(ctxFiber *fiber.Ctx, ctx context.Context, payload dto.RequestLoginUser) (*dto.ResponseLoginUser, error)
	//* Get Profile One Id
	GetProfileOneIdService(ctx context.Context, accessToken string) (dto.ResponseUserProfile, error)
}
