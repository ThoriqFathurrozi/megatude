package migrations

import (
	"fmt"

	earthquake "github.com/ThoriqFathurrozi/megatude/internal/domains/earthquake/entity"
	"github.com/ThoriqFathurrozi/megatude/internal/types"
	"gorm.io/gorm"
)

type Migrator struct {
	db       *gorm.DB
	entities []types.IEntity
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{
		db: db,
		entities: []types.IEntity{
			earthquake.Earthquake{},
		},
	}
}

func (m *Migrator) Migrate(db *gorm.DB) error {
	var entities []interface{}

	for _, entity := range m.entities {
		entities = append(entities, entity)
	}

	if err := db.AutoMigrate(entities...); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err.Error())
	}

	return nil
}

func (m *Migrator) Purge() error {
	var entities []interface{}

	for _, entity := range m.entities {
		entities = append(entities, entity)
	}

	if err := m.db.Migrator().DropTable(entities...); err != nil {
		return fmt.Errorf("failed to purge database: %v", err.Error())
	}

	return nil
}
