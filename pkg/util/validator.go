package util

import (
	"fmt"
	"regexp"
	"slices"
	"time"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/rand"
)

func ValidatorStruct(payload interface{}) (string, error) {
	var validate = validator.New()
	validate.RegisterValidation("regexp", Regexp)
	validate.RegisterValidation("ThaiOnly", ThaiOnly)
	validate.RegisterValidation("ThaiAndDashOnly", ThaiAndDashOnly)
	validate.RegisterValidation("EnglishOnly", EnglishOnly)
	if err := validate.Struct(payload); err != nil {
		logrus.Infoln("err validate struct==>", err)
		errType := make(map[string][]string)
		for _, err := range err.(validator.ValidationErrors) {
			errType[err.Tag()] = append(errType[err.Tag()], err.StructNamespace())
		}
		msg := formatValidationErrors(errType)
		return msg, err
	}
	return "", nil
}

func formatValidationErrors(errType map[string][]string) string {
	format := []string{"ThaiOnly", "regexp", "ThaiAndDashOnly", "EnglishOnly"}
	msg := ""
	i := 0
	for e, m := range errType {
		namespaces := joinNamespaces(m)
		switch {
		case e == "required":
			msg += "[missing parameters]: " + namespaces
		case slices.Contains(format, e):
			msg += "[invalid format]: " + namespaces
		default:
			msg += e + ": " + namespaces
		}
		if i < len(errType)-1 {
			msg += " | "
		}
		i++
	}
	return msg
}

func joinNamespaces(namespaces []string) string {
	if len(namespaces) == 0 {
		return ""
	}
	result := namespaces[0]
	for i := 1; i < len(namespaces); i++ {
		result += ", " + namespaces[i]
	}
	return result
}

func Regexp(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(fl.Param())
	return re.MatchString(fl.Field().String())
}

func RandomNumber(max, length int) string {
	number := rand.Intn(max)
	return fmt.Sprintf("%0*d", length, number)
}

func RandomNumberPin(length int) string {
	UnixNano := time.Now().UnixNano()
	rand.Seed(uint64(UnixNano))
	digits := "0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = digits[rand.Intn(len(digits))]
	}

	return string(result)
}

func ThaiOnly(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString(`^[ก-๙]+$`, fl.Field().String())
	return match
}

func EnglishOnly(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString(`^[A-Za-z-]+$`, fl.Field().String())
	return match
}

func ThaiAndDashOnly(fl validator.FieldLevel) bool {
	match, _ := regexp.MatchString(`^[ก-๙\-]+$`, fl.Field().String())
	return match
}
