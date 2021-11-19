package stores

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActivationType string

const (
	ACTIVATION_CODE ActivationType = "ACTIVATION_CODE"
	FORGOT_PASSWORD ActivationType = "FORGOT_PASSWORD"
)

type UserActivation struct {
	gorm.Model
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	UserId    uuid.UUID      `gorm:"type:char(36):index"`
	User      User           `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Code      string         `gorm:"type:char(32);index;not null"`
	ActType   ActivationType `gorm:"type:char(30);index;not null"`
	ExpiredAt *time.Time
}

func (*UserActivation) BeforeCreate(tx *gorm.DB) error {
	t := time.Now()
	newT := t.Add(time.Hour * 2)

	tx.Statement.SetColumn("ID", uuid.New())
	tx.Statement.SetColumn("ExpiredAt", newT)
	return nil
}
