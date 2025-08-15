package dto

type RequestRegisterUser struct {
	TitleTh        string `json:"account_title_th" validate:"required"`
	SpecialTitleTh string `json:"special_title_name_th,omitempty"`
	FirstNameTh    string `json:"first_name_th" validate:"required"`
	MiddleNameTh   string `json:"middle_name_th,omitempty"`
	LastNameTh     string `json:"last_name_th" validate:"required"`

	SpecialTitleEng string `json:"special_title_name_eng,omitempty"`
	TitleEng        string `json:"account_title_eng" validate:"required"`
	FirstNameEng    string `json:"first_name_eng" validate:"required"`
	MiddleNameEng   string `json:"middle_name_eng,omitempty"`
	LastNameEng     string `json:"last_name_eng" validate:"required"`
	Gender          string `json:"gender,omitempty"`

	Email     string `json:"email" validate:"required,email"`
	MobileNo  string `json:"mobile_no" validate:"required"`
	BirthDate string `json:"birth_date" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`

	IdCardNum string `json:"id_card_num" validate:"required"`
}
