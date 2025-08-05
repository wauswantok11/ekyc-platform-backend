package model

import (
	"gorm.io/gorm"
)

type Nationality struct {
	Id        uint            `gorm:"column:id;primaryKey;size:50"  json:"id"`
	Code      string          `gorm:"column:code;size:5"  json:"code"`
	Nation    string          `gorm:"column:nation;size:100"  json:"nation"`
	Note      string          `gorm:"column:note" json:"note"`
	CreatedBy string          `gorm:"column:created_by;size:50" json:"created_by"`
	UpdatedBy string          `gorm:"column:updated_by;size:50" json:"updated_by"`
	Status    int16           `gorm:"status;index;" json:"status"`
	gorm.Model
}

func (Nationality) TableName() string {
	return "nationality"
}


// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (1, 'AF', 'Afghan', 'อัฟกัน', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (2, 'AL', 'Albanian', 'แอลเบเนีย', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (3, 'DZ', 'Algerian', 'แอลจีเรีย', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (4, 'AD', 'Andorran', 'อันดอร์รา', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (5, 'AO', 'Angolan', 'แองโกลา', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (6, 'AR', 'Argentine', 'อาร์เจนตินา', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (7, 'AM', 'Armenian', 'อาร์เมเนีย', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (8, 'AU', 'Australian', 'ออสเตรเลีย', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (9, 'AT', 'Austrian', 'ออสเตรีย', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (10, 'AZ', 'Azerbaijani', 'อาเซอร์ไบจาน', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (11, 'BD', 'Bangladeshi', 'บังกลาเทศ', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (12, 'BE', 'Belgian', 'เบลเยียม', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (13, 'BJ', 'Beninese', 'เบนิน', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (14, 'BO', 'Bolivian', 'โบลิเวีย', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (15, 'BR', 'Brazilian', 'บราซิล', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (16, 'BN', 'Bruneian', 'บรูไน', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (17, 'BG', 'Bulgarian', 'บัลแกเรีย', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (18, 'KH', 'Cambodian', 'กัมพูชา', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (19, 'CA', 'Canadian', 'แคนาดา', 'system', 'system', 1, NOW(), NOW(), NULL);
// INSERT INTO nationalities (id, code, nation, note, created_by, updated_by, status, created_at, updated_at, deleted_at) VALUES (20, 'CL', 'Chilean', 'ชิลี', 'system', 'system', 1, NOW(), NOW(), NULL);
 
