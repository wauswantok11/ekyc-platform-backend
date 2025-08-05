package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Zipcode struct {
	Id              uuid.UUID `gorm:"column:id;primaryKey;size:50;index;"  json:"id"`
	District        string    `gorm:"column:district;size:255;index;"  json:"district"`
	Subdistrict     string    `gorm:"column:subdistrict;size:255;index;"  json:"subdistrict"`
	Province        string    `gorm:"column:province;size:255;index;"  json:"province"`
	Zipcode         string    `gorm:"column:zipcode;size:50;index;"  json:"zipcode"`
	SubdistrictCode string    `gorm:"column:subdistrict_code;size:10;index;"  json:"subdistrict_code"`
	DistrictCode    string    `gorm:"column:district_code;size:10;index;"  json:"district_code"`
	ProvinceCode    string    `gorm:"column:province_code;size:10;index;" json:"province_code"`
	CreatedBy       string     `gorm:"column:created_by;size:50" json:"created_by"`
	UpdatedBy       string     `gorm:"column:updated_by;size:50" json:"updated_by"`
	gorm.Model

}

func (Zipcode) TableName() string {
	return "zipcode"
}
 
