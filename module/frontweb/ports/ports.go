package ports

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"

	"git.inet.co.th/ekyc-platform-backend/config"
	"git.inet.co.th/ekyc-platform-backend/model"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/dto"
	"git.inet.co.th/ekyc-platform-backend/pkg/cache"
	"git.inet.co.th/ekyc-platform-backend/pkg/database"
	oneId "git.inet.co.th/ekyc-platform-backend/pkg/one-id"
)

type Repository interface {
	Module() string
	AppCfg() *config.Config
	Log() *logrus.Entry
	OneId() *oneId.Client
	DB() *database.Client
	Cache() *cache.Redis
	Trace(ctx context.Context, spanName string, attributes ...trace.SpanStartOption) (context.Context, trace.Span)

	GetAccountByAccountIdOneRepo(ctx context.Context, accountIdOne string) (string, error)
	GenJwtTokenRepo(ctx context.Context, dataToken map[string]interface{}) (string, error)
	//* Redis Repo
	// SetRedisRepo(ctx context.Context, cKey, Accesstoken string) error
	SetRedisRepo(ctx context.Context, cKey string, userProfile map[string]interface{}) error
	DelRedisRepo(ctx context.Context, cKey string) error
	//* CRUD UserRepo
	FindUserByAccountIdRepo(ctx context.Context, accountId string) (*string, error)
	FindUserDetailByAccountIdRepo(ctx context.Context, accountId string) (*model.Account, error)

	CreateUserRepo(ctx context.Context, userProfile map[string]interface{}) error
	UpdateUserRepo(ctx context.Context, userProfile map[string]interface{}, id *string) error
}

type Service interface {
	//* Login Account One Id
	LoginUserOneService(ctxFiber *fiber.Ctx, ctx context.Context, payload dto.RequestLoginUser) (*dto.ResponseLoginUser, string, error)
	LogoutUserService(ctxFiber *fiber.Ctx, ctx context.Context, keyCookie, accountId string) error
	
	//* Get Profile One Id
	GetProfileOneIdService(ctx context.Context, accountId, token string) (*dto.ResponseUserProfile, string, error)
	GetProfileOneAvatarByAccountOneIdService(ctx context.Context, accountOneId string) (string, error)
}
