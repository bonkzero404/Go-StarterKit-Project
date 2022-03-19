package driver

import (
	"go-starterkit-project/config"
	"go-starterkit-project/domain/stores"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func CasbinAdapterConnect() (*gormadapter.Adapter, error) {
	CasbinAdapter, CasbinAdapterError := gormadapter.NewAdapterByDBWithCustomTable(DB, &stores.CasbinRule{})

	if CasbinAdapterError != nil {
		return nil, CasbinAdapterError
	}

	return CasbinAdapter, CasbinAdapterError
}

func Casbin() (*casbin.Enforcer, error) {
	// CasbinRbac = &casbin.Enforcer{}
	adapter, _ := CasbinAdapterConnect()
	CasbinRbac, err := casbin.NewEnforcer(config.Config("CASBIN_MODEL"), adapter)

	if err != nil {
		return nil, err
	}

	//CasbinRbac.LoadPolicy()

	return CasbinRbac, err
}
