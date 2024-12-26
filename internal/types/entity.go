package types

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IEntity interface {
	TableName() string
}

// Entity struct
type Entity struct {
	UUID      uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"uuid"`
	CreatedAt time.Time      `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName return table name
func (e Entity) TableName() string {
	return "entities"
}
