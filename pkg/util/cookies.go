package util

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetCookieHandler(ctx *fiber.Ctx, name string, token string) error {
	if ctx == nil {
		logrus.Error("fiber.Ctx is nil — cannot set cookie")
		return errors.New("fiber is null")
	}

	if name == "" || token == "" {
		return errors.New("bad request")
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    token,
		MaxAge:   86400,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return nil
}

// func GetCookieHandler(w http.ResponseWriter, r *http.Request) (string, error) {
// 	cookie, err := r.Cookie("exampleCookie")
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, http.ErrNoCookie):
// 			return errors.New("not found")
// 		default:
// 			log.Println(err)
// 			return err
// 		}

// 	}

// 	return w.Write([]byte(cookie.Value)), nil
// }

func DelCookieHandler() error {
	return nil
}
