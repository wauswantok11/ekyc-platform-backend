package dto

type RequestForgetWithEmail struct {
	Email         string `json:"email" validate:"required"`
	UsernameOrCid string `json:"username_or_cid" validate:"required,min=6"`
}

type RequestForgetWithMobile struct {
	Mobile        string `json:"mobile" validate:"required,min=10,mix=10"`
	UsernameOrCid string `json:"username_or_cid" validate:"required,min=6"`
}

type RequestForgetPassword struct {
	Otp string `json:"otp" validate:"required,min=6,mix=6"`
}
