package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CasbinRule struct {
	ID    uuid.UUID `gorm:"type:char(36);primary_key"`
	Ptype string    `gorm:"size:512;uniqueIndex:unique_index"`
	V0    string    `gorm:"size:512;uniqueIndex:unique_index"`
	V1    string    `gorm:"size:512;uniqueIndex:unique_index"`
	V2    string    `gorm:"size:512;uniqueIndex:unique_index"`
	V3    string    `gorm:"size:512;uniqueIndex:unique_index"`
	V4    string    `gorm:"size:512;uniqueIndex:unique_index"`
	V5    string    `gorm:"size:512;uniqueIndex:unique_index"`
}

func (*CasbinRule) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
