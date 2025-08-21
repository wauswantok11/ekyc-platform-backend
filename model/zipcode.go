package model

type Zipcode struct {
	Model
	District        string `gorm:"column:district;size:255;index;"  json:"district"`
	Subdistrict     string `gorm:"column:subdistrict;size:255;index;"  json:"subdistrict"`
	Province        string `gorm:"column:province;size:255;index;"  json:"province"`
	Zipcode         string `gorm:"column:zipcode;size:50;index;"  json:"zipcode"`
	SubdistrictCode string `gorm:"column:subdistrict_code;size:10;index;"  json:"subdistrict_code"`
	DistrictCode    string `gorm:"column:district_code;size:10;index;"  json:"district_code"`
	ProvinceCode    string `gorm:"column:province_code;size:10;index;" json:"province_code"`
	CreatedBy       string `gorm:"column:created_by;size:50" json:"created_by"`
	UpdatedBy       string `gorm:"column:updated_by;size:50" json:"updated_by"`
}

func (Zipcode) TableName() string {
	return "zipcode"
}
