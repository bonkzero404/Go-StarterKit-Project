package driver

import (
	"go-starterkit-project/domain/stores"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var CASBIN *casbin.Enforcer

func ConnectCasbin() *casbin.Enforcer {

	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(DB, &stores.CasbinRule{})

	CASBIN, _ := casbin.NewEnforcer("casbin_models/rbac_model.conf", adapter)

	// if err != nil {
	// 	// errMessage := fmt.Sprintf("Failed to connect database")
	// 	panic(err)
	// }

	// e.LoadPolicy()

	// // Check the permission.
	// e.Enforce("alice", "data1", "read")

	// // Modify the policy.
	// e.AddPolicy("alice", "data1", "read")
	// // e.RemovePolicy(...)

	// // Save the policy back to DB.
	// e.SavePolicy()
	return CASBIN
}
