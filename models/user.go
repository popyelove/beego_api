package models

import (
	"github.com/astaxie/beego/orm"
)

type T_rbac_user struct {
	Rbac_user_id          		int 				`orm:"pk"`
	Username        			string
	Password        			string
}
func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(T_rbac_user))
}
