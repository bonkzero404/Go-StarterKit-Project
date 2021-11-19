package database

import (
	"go-boilerplate-clean-arch/domain/stores"
)

func MigrateDB() {
	DB.AutoMigrate(
		&stores.User{},
		&stores.UserActivation{},
	)
}
