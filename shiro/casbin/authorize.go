package casbin

import (
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
)

//CasbinInstance is a casbininstance
type CasbinInstance struct {
	enforcer *casbin.SyncedEnforcer
}

//NewCasbinInstance initialize a casbininstance
func NewCasbinInstance() *CasbinInstance {
	//e, err := casbin.NewEnforcerSafe("casbin_model.conf", "casbin_policy.csv")
	a := gormadapter.NewAdapter("sqlite3", "casbin_policy.db")
	//e := casbin.NewSyncedEnforcer("casbin_model.conf", "casbin_policy.csv")
	e := casbin.NewSyncedEnforcer("casbin_model.conf", a)
	//load the policy from db
	e.LoadPolicy()
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

//AddPolicy add one policy to the policys
func (c *CasbinInstance) AddPolicy(principal string, obj string, act string) bool {
	ok := c.enforcer.AddPolicy(principal, obj, act)
	//save the policys back to DB
	c.enforcer.SavePolicy()
	return ok
}

//RemovePolicy remove one policy from the policys
func (c *CasbinInstance) RemovePolicy(principal string, obj string, act string) bool {
	ok := c.enforcer.RemovePolicy(principal, obj, act)
	//save the policys back to DB
	c.enforcer.SavePolicy()
	return ok
}
