package stores

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

/**
Table model
*/
type Acl struct {
	gorm.Model
	ID             uuid.UUID `gorm:"type:char(36);primary_key"`
	AclCode        string    `gorm:"type:varchar(20);uniqueIndex;not null"`
	AclName        string    `gorm:"type:varchar(100);index;not null"`
	AclDescription string    `gorm:"type:text"`
	IsActive       bool
}

/**
This function is a feature that gorm has for making hooks,
this hook function is used to generate uuid when the user
performs the create action
*/
func (*Acl) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}
