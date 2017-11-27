package main

import (
	_"api/routers"
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/session/redis"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)
func init() {
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 30
	maxConn := 30
	//root:lichao@tcp(127.0.0.1:3306)/beego?charset=utf8
	mysqluser :=beego.AppConfig.String("mysqluser")
	mysqlpass :=beego.AppConfig.String("mysqlpass")
	mysqlurls :=beego.AppConfig.String("mysqlurls")
	mysqldb :=beego.AppConfig.String("mysqldb")
	mysqlport :=beego.AppConfig.String("mysqlport")
	link :=mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+":"+mysqlport+")/"+mysqldb+"?charset=utf8"
	orm.RegisterDataBase("default", "mysql",link, maxIdle, maxConn)
}
func main() {
	beego.Run()
}
