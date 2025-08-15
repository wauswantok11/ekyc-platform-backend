package model

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	Id        uuid.UUID      `json:"id" gorm:"type:varchar(50);primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	if m.Id == uuid.Nil {
		m.Id, _ = uuid.NewV7()
	}
	return nil
}
