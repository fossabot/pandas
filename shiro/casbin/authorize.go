package casbin

import (
	"github.com/casbin/casbin"
)

//CasbinInstance is a casbininstance
type CasbinInstance struct {
	enforcer *casbin.SyncedEnforcer
}

//NewCasbinInstance initialize a casbininstance
func NewCasbinInstance() *CasbinInstance {
	//e, err := casbin.NewEnforcerSafe("casbin_model.conf", "casbin_policy.csv")
	e := casbin.NewSyncedEnforcer("casbin_model.conf", "casbin_policy.csv")
	//e, err := casbin.NewEnforcer("casbin_model.conf", "casbin_policy.csv")
	return &CasbinInstance{
		enforcer: e,
	}
}

//Authroize return whether principal has the right to access
func (c *CasbinInstance) Authroize(principal string, obj string, act string) bool {
	ok := c.enforcer.Enforce(principal, obj, act)
	return ok
}
