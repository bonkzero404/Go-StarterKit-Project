package database

import (
	"go-boilerplate-clean-arch/database/driver"
	"go-boilerplate-clean-arch/domain/stores"
)

/**
This function is used for auto migration and is loaded
into the main function
*/
func MigrateDB() {
	driver.DB.AutoMigrate(
		&stores.User{},
		&stores.UserActivation{},
	)
}
