package driver

import (
	"go-starterkit-project/config"
	"go-starterkit-project/domain/stores"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var CASBIN *casbin.Enforcer

func ConnectCasbin() *casbin.Enforcer {
	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(DB, &stores.CasbinRule{})
	CASBIN, _ := casbin.NewEnforcer(config.Config("CASBIN_MODEL"), adapter)
	return CASBIN
}
