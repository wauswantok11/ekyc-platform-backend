package frontweb

import (
	"github.com/gofiber/fiber/v2"

	"git.inet.co.th/ekyc-platform-backend/app"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/handler"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/repositories"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/services"
)

func Create(app *app.Context) error {
	repo, err := repositories.New(app)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	svc := services.New(repo)
	handler := handler.NewHandler(svc)
	router := app.Router.Group(app.Config.App.PrefixPath)
	addRouter(router, handler)
	return nil
}

func addRouter(router fiber.Router, handler *handler.Handler) {
	v1 := router.Group("/v1")

	session := v1.Group("/session")
	session.Get("/", handler.GetSessionHandler)

	loginUser := v1.Group("/login")
	loginUser.Post("/user", handler.PostLoginUserHandler)
	loginUser.Post("/cid-mobile", handler.PostLoginUserHandler)
	loginUser.Post("/", handler.PostLoginUserHandler)

	sharedToken := v1.Group("/shared-token")
	sharedToken.Get("/:shared_token", handler.GetSharedTokenHandler)

	forgetPassword := v1.Group("/forget-password")
	forgetPassword.Post("/email", handler.PostForgetPasswordEmailUserHandler)
	forgetPassword.Post("/mobile", handler.PostForgetPasswordMobileUserHandler)

	register := v1.Group("/register")
	register.Post("/", handler.PostRegisterUserHandler)
}
