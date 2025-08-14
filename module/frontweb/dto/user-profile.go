package dto

type RequestLoginUser struct {
	Username string `json:"username"  validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
}

type ResponseLoginUser struct {
	TokenType      string              `json:"token_type,omitempty"`
	ExpiresIn      int64               `json:"expires_in ,omitempty"`
	AccessToken    string              `json:"access_token ,omitempty"`
	RefreshToken   string              `json:"refresh_token ,omitempty"`
	ExpirationDate string              `json:"expiration_date ,omitempty"`
	AccountId      string              `json:"account_id ,omitempty"`
	Result         string              `json:"result ,omitempty"`
	Username       string              `json:"username ,omitempty"`
	LoginBy        string              `json:"login_by ,omitempty"`
	UserProfile    ResponseUserProfile `json:"user_profile ,omitempty"`
}

type RequestUserProfile struct {
	Token string `json:"token"`
}

type ResponseUserProfile struct {
	AccessId            string `json:"account_id ,omitempty"`
	FirstNameTH         string `json:"first_name_th ,omitempty"`
	MiddleNameTH        string `json:"middle_name_th,omitempty"`
	LastNameTH          string `json:"last_name_th ,omitempty"`
	FirstNameEng        string `json:"first_name_eng ,omitempty"`
	MiddleNameEng       string `json:"middle_name_eng ,omitempty"`
	LastNameEng         string `json:"last_name_eng ,omitempty"`
	SpecialTitleNameTH  string `json:"special_title_name_th ,omitempty"`
	AccountTitleTH      string `json:"account_title_th ,omitempty"`
	SpecialTitleNameEng string `json:"special_title_name_eng ,omitempty"`
	AccountTitleEng     string `json:"account_title_eng ,omitempty"`
	IdCardType          string `json:"id_card_type ,omitempty"`
	IdCardNum           string `json:"id_card_num ,omitempty"`
	HashIdCardNum       string `json:"hash_id_card_num ,omitempty"`
	AccountCategory     string `json:"account_category ,omitempty"`
	AccountSubCategory  string `json:"account_sub_category ,omitempty"`
	ThaiEmail           string `json:"thai_email ,omitempty"`
	ThaiEmail2          string `json:"thai_email2 ,omitempty"`
	ThaiEmail3          string `json:"thai_email3 ,omitempty"`
	StatusCD            string `json:"status_cd ,omitempty"`
	BirthDate           string `json:"birth_date ,omitempty"`
}
