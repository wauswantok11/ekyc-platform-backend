package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Model struct {
	Id        uuid.UUID      `json:"id" gorm:"primaryKey;index"`
	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	if uuid.Equal(m.Id, uuid.Nil) {
		m.Id = uuid.NewV4()
	}
	return nil
}
