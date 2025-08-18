package one_id

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/trace"

	"git.inet.co.th/ekyc-platform-backend/pkg/requests"
)

type Client struct {
	url          string
	clientId     string
	clientSecret string
	refCode      string
	timeOut      int
	http         *requests.HttpClient
	log          *logrus.Entry
	tracer       trace.Tracer
}

type ResponseErrorOneId struct {
	Result       string      `json:"result"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
	ResponseCode int64       `json:"responseCode"`
}

type ResponseSuccessLoginPWD struct {
	TokenType      string `json:"token_type"`
	ExpiresIn      int64  `json:"expires_in"`
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	ExpirationDate string `json:"expiration_date"`
	AccountId      string `json:"account_id"`
	Result         string `json:"result"`
	Username       string `json:"username"`
	LoginBy        string `json:"login_by"`
}

type RequestLoginPWD struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type ResponseApiAccountOneId struct {
	ID                  string `json:"id"`
	FirstNameTH         string `json:"first_name_th"`
	MiddleNameTH        string `json:"middle_name_th"`
	LastNameTH          string `json:"last_name_th"`
	FirstNameEng        string `json:"first_name_eng"`
	MiddleNameEng       string `json:"middle_name_eng"`
	LastNameEng         string `json:"last_name_eng"`
	SpecialTitleNameTH  string `json:"special_title_name_th"`
	AccountTitleTH      string `json:"account_title_th"`
	SpecialTitleNameEng string `json:"special_title_name_eng"`
	AccountTitleEng     string `json:"account_title_eng"`
	IDCardType          string `json:"id_card_type"`
	IDCardNum           string `json:"id_card_num"`
	HashIDCardNum       string `json:"hash_id_card_num"`
	AccountCategory     string `json:"account_category"`
	AccountSubCategory  string `json:"account_sub_category"`
	ThaiEmail           string `json:"thai_email"`
	ThaiEmail2          string `json:"thai_email2"`
	ThaiEmail3          string `json:"thai_email3"`
	StatusCD            string `json:"status_cd"`
	BirthDate           string `json:"birth_date"`
}
type ResponseLoginMobileOTP struct {
	Result string `json:"result"`
	Data   struct {
		Otp     string `json:"otp"`
		Refcode string `json:"refcode"`
	} `json:"data"`
	ErrorMessage *string `json:"errorMessage"`
	Code         int     `json:"code"`
}

type RequestLoginMobileOTP struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Refcode      string `json:"refcode"`
	MobileNo     string `json:"mobile_no"`
}

type ResponseCheckDupUsername struct {
	Message      string  `json:"message"`
	Data         *string `json:"data"`
	ErrorMessage *string  `json:"errorMessage"`
}