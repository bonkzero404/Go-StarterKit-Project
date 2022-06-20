package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/**
Table model
*/
type AclRole struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	AclId      uuid.UUID `gorm:"type:char(36):index"`
	RoleUserId uuid.UUID `gorm:"type:char(36):index"`
	Acl        Acl       `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RoleUser   RoleUser  `gorm:"references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsActive   bool
}

/**
This function is a feature that gorm has for making hooks,
this hook function is used to generate uuid when the user
performs the create action
*/
func (*AclRole) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
