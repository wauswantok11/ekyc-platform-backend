package model

type Nationality struct {
	Model
	Code      string `gorm:"column:code;size:5"  json:"code"`
	Nation    string `gorm:"column:nation;size:100"  json:"nation"`
	Note      string `gorm:"column:note" json:"note"`
	CreatedBy string `gorm:"column:created_by;size:50" json:"created_by"`
	UpdatedBy string `gorm:"column:updated_by;size:50" json:"updated_by"`
	Status    int16  `gorm:"status;index;" json:"status"`
}

func (Nationality) TableName() string {
	return "nationality"
}
