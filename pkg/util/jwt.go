package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(jwtSecretKey string, payload map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{}
	for k, v := range payload {
		claims[k] = v
	}
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
 
	return token.SignedString([]byte(jwtSecretKey))
}
func ParseJWT(jwtSecretKey, tokenStr string) (map[string]interface{}, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if expRaw, exists := claims["exp"]; exists {
			switch exp := expRaw.(type) {
			case float64:
				if time.Unix(int64(exp), 0).Before(time.Now()) {
					return nil, errors.New("expired jwt")
				}
			default:
				return nil, errors.New("invalid exp claim type")
			}
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
