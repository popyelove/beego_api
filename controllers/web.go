package controllers

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"api/common"
)

// Operations about Users
type WebController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *WebController) GetAll() {
	rd,err := redis.Dial(beego.AppConfig.String("redisnetwork"), beego.AppConfig.String("redishost")+":"+ beego.AppConfig.String("redisport"),	redis.DialPassword("lichao"))
	if err != nil{
		fmt.Println(err)
		return
	}
	env,err :=rd.Do("del","name2")
	if err !=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(env)
	fmt.Println(common.Add(1,2))
	defer rd.Close()
}
// @Title schemeinfo
// @Description get all Users
// @Success 200 {object} models.User
// @router /schemeinfo [get]
func (u *WebController) Schemeinfo()  {
	log :=common.Conf("host")
	println(log)
}


