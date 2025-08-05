package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountOneId string `gorm:"column:house_no;type:text;size:255" json:"account_id_one"`

	CidEncrypt      string `gorm:"column:cid_encrypt;type:text;" json:"cid_encrypt"`
	CidHash         string `gorm:"column:cid_hash;index;size:255;" json:"cid_hash"`
	LaserIdEncrypt  string `gorm:"column:laser_id_encrypt;type:text;" json:"laser_id_encrypt"`
	LaserIdHash     string `gorm:"column:laser_id_hash;index;size:255;" json:"laser_id_hash"`
	PassportEncrypt string `gorm:"column:passport_encrypt;type:text;" json:"passport_encrypt"`
	PassportHash    string `gorm:"column:passport_hash;index;size:255;" json:"passport_hash"`

	SpecialTitleTh string `gorm:"column:special_title_th;size:255" json:"special_title_th"`
	TitleTh        string `gorm:"column:title_th;size:255" json:"title_th"`
	FirstNameTh    string `gorm:"column:first_name_th;size:255" json:"first_name_th"`
	MiddleNameTh   string `gorm:"column:middle_name_th;size:255" json:"middle_name_th"`
	LastNameTh     string `gorm:"column:last_name_th;size:255" json:"last_name_th"`

	SpecialTitleEng string `gorm:"column:special_title_eng;size:255" json:"special_title_eng"`
	TitleEng        string `gorm:"column:title_eng;size:255" json:"title_eng"`
	FirstNameEng    string `gorm:"column:firstname_eng;size:255" json:"firstname_eng"`
	MiddleNameEng   string `gorm:"column:middle_name_eng;size:255" json:"middle_name_eng"`
	LastNameEng     string `gorm:"column:lastname_eng;size:255" json:"lastname_eng"`

	Gender       string       `gorm:"column:gender;size:255" json:"gender"`
	Email        string       `gorm:"column:email;size:255;index" json:"email"`
	Phone        string       `gorm:"column:phone;size:255;index" json:"phone"`
	BirthDate    string       `gorm:"column:birth_date" json:"birth_date"`
	NationCode   string       `gorm:"column:nation_code;size:255" json:"nation_code"`
	Nationality  *Nationality `gorm:"foreignKey:NationCode;references:code" json:"nation_detail"`
	TypeRegister string       `gorm:"column:type_register;size:50;index;" json:"type_register"`
	AddressId    uuid.UUID    `gorm:"column:address_id;size:50;index" json:"address_id"`
	Address      *Address     `gorm:"->;foreignKey:AddressId;references:Id" json:"address"`

	LastLogin time.Time `gorm:"column:last_login;" json:"last_login"`
	Username  string    `gorm:"column:username;size:50;index;" json:"username"`
	CreatedBy string    `gorm:"column:created_by;size:50"  json:"created_by"`
	UpdatedBy string    `gorm:"column:updated_by;size:50"  json:"updated_by"`
}

func (Account) TableName() string {
	return "account"
}
