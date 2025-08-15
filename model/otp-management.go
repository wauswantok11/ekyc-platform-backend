package model

import uuid "github.com/google/uuid"

type OtpManagement struct {
	Model
	OtpCode         string     `gorm:"column:otp_code;type:varchar(10)" json:"otp_code"`
	RefCode         string     `gorm:"column:ref_code;type:varchar(20)" json:"ref_code"`
	OtpMakeBy       string     `gorm:"column:otp_make_by;type:varchar(50)" json:"otp_make_by"`
	MobileNo        string     `gorm:"column:mobile_no;type:varchar(250)" json:"mobile_no"`
	OtpFor          string     `gorm:"column:otp_for;type:varchar(100)" json:"otp_for"`
	OtpStatus       string     `gorm:"column:otp_status;type:varchar(50)" json:"otp_status"`
	AccountDetailId *uuid.UUID `gorm:"column:account_detail_id;index;default:NULL" json:"account_detail_id"`
	CreatedBy       string     `gorm:"column:created_by;type:varchar(50)" json:"created_by"`
	UpdatedBy       string     `gorm:"column:updated_by;type:varchar(50)" json:"updated_by"`
}

func (OtpManagement) TableName() string {
	return "otp_management"
}
