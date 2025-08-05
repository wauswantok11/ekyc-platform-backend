package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Address struct {
	Id            uuid.UUID `gorm:"column:id;index" json:"id"`
	HouseNo       string    `gorm:"column:house_no;type:text;size:255" json:"house_no"`
	Soi           *string   `gorm:"column:soi;type:text;size:255" json:"soi"`
	Village       *string   `gorm:"column:village;type:text;size:255" json:"village"`
	Street        *string   `gorm:"column:street;type:text;size:255" json:"street"`
	ZipcodeId     uuid.UUID `gorm:"column:zipcode_id;size:50;index" json:"zipcode_id"`
	Zipcode       Zipcode   `gorm:"->;foreignKey:ZipcodeId;references:Id" json:"zipcode"`
	AddressDetail string    `gorm:"column:address_detail;type:text" json:"address_detail"`
	CreatedBy     string    `gorm:"column:created_by;size:50"`
	UpdatedBy     string    `gorm:"column:updated_by;size:50"`
	gorm.Model
}

func (Address) TableName() string {
	return "address"
}
