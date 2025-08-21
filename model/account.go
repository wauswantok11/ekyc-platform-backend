package model

import (
	uuid "github.com/google/uuid"
	"time"
)

type Account struct {
	Model
	AccountOneId    string    `gorm:"column:account_one_id;type:varchar(255)" json:"account_one_id"`
	CidType         string    `gorm:"column:cid_type;type:varchar(255)" json:"cid_type"`
	CidEncrypt      string    `gorm:"column:cid_encrypt;type:text" json:"cid_encrypt"`
	CidHash         string    `gorm:"column:cid_hash;type:varchar(255);index" json:"cid_hash"`
	LaserIdEncrypt  string    `gorm:"column:laser_id_encrypt;type:text" json:"laser_id_encrypt"`
	LaserIdHash     string    `gorm:"column:laser_id_hash;type:varchar(255);index" json:"laser_id_hash"`
	PassportEncrypt string    `gorm:"column:passport_encrypt;type:text" json:"passport_encrypt"`
	PassportHash    string    `gorm:"column:passport_hash;type:varchar(255);index" json:"passport_hash"`
	SpecialTitleTh  string    `gorm:"column:special_title_th;type:varchar(255)" json:"special_title_th"`
	TitleTh         string    `gorm:"column:title_th;type:varchar(255)" json:"title_th"`
	FirstNameTh     string    `gorm:"column:first_name_th;type:varchar(255)" json:"first_name_th"`
	MiddleNameTh    string    `gorm:"column:middle_name_th;type:varchar(255)" json:"middle_name_th"`
	LastNameTh      string    `gorm:"column:last_name_th;type:varchar(255)" json:"last_name_th"`
	SpecialTitleEng string    `gorm:"column:special_title_eng;type:varchar(255)" json:"special_title_eng"`
	TitleEng        string    `gorm:"column:title_eng;type:varchar(255)" json:"title_eng"`
	FirstNameEng    string    `gorm:"column:firstname_eng;type:varchar(255)" json:"firstname_eng"`
	MiddleNameEng   string    `gorm:"column:middle_name_eng;type:varchar(255)" json:"middle_name_eng"`
	LastNameEng     string    `gorm:"column:lastname_eng;type:varchar(255)" json:"lastname_eng"`
	Gender          string    `gorm:"column:gender;type:varchar(255)" json:"gender"`
	Email           string    `gorm:"column:email;type:varchar(255);index" json:"email"`
	Phone           string    `gorm:"column:phone;type:varchar(255);index" json:"phone"`
	BirthDate       string    `gorm:"column:birth_date;type:varchar(50)" json:"birth_date"`
	NationCode      string    `gorm:"column:nation_code;type:varchar(255)" json:"nation_code"`
	TypeRegister    string    `gorm:"column:type_register;type:varchar(50);index" json:"type_register"`
	AddressId       uuid.UUID `gorm:"column:address_id;type:varchar(50);index" json:"address_id"`
	LastLogin       time.Time `gorm:"column:last_login" json:"last_login"`
	Username        string    `gorm:"column:username;type:varchar(50);index" json:"username"`
	CreatedBy       string    `gorm:"column:created_by;type:varchar(50)" json:"created_by"`
	UpdatedBy       string    `gorm:"column:updated_by;type:varchar(50)" json:"updated_by"`
}

func (Account) TableName() string {
	return "account"
}
