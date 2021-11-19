package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserActivation struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:char(36);primary_key"`
	UserId uuid.UUID `gorm:"type:char(36):index"`
	User   User      `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Code   string    `gorm:"type:char(32);index;not null"`
	IsUsed bool
}

func (*UserActivation) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
