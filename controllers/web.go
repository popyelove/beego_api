package controllers

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"api/common"
	"log"
	"github.com/bitly/go-simplejson"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"api/models"
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
	client_schmeid :=beego.AppConfig.String("clientschmeid")
	str :="window.CONFIG="
	res :=common.Conf("all",client_schmeid)
	str+=res+";"
	u.Ctx.WriteString(str)

}
// @Title CheckLogin
// @Description find object by objectid
// @Success 200
// @router /checkLogin [post]
func (o *WebController) CheckLogin() {
	req :=o.Ctx.Input.RequestBody
	req1,_:=simplejson.NewJson(req)
	log.Println(req1.Get("name"))
	res :=common.Error(11,"缺失 openid")
	o.Data["json"]=res
	o.ServeJSON()
}
// @Title CheckLogin
// @Description find object by objectid
// @Success 200
// @router /login [post]
func (o *WebController) Login() {
	req :=o.Ctx.Input.RequestBody
	//res :=common.Success()
	//o.Data["json"]=res
	//o.ServeJSON()
	userinfo,_ :=simplejson.NewJson(req)

	ob := orm.NewOrm()
	ob.Using("psys") // 默认使用 default，你可以指定为其他数据库


	//user := new(models.T_rbac_user)
	//user.Username = "chao"
	//fmt.Println(ob.Delete(user,"Username"))

	//user := new(models.T_rbac_user)
	//user.Username = "chao"
	//fmt.Println(ob.Insert(user))


	//user := new(models.T_rbac_user)
	username,_ :=userinfo.Get("username").String()
	password,_ :=userinfo.Get("password").String()
	passworddata :=[]byte(password)
	passwordhas :=md5.Sum(passworddata)
	passwordmd5 :=fmt.Sprintf("%x",passwordhas)
	//user := models.T_rbac_user{Username:username}
	//
	//err := ob.Read(&user,"Username")
	//
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	fmt.Println(user.Rbac_user_id, user.Username)
	//}
	var user models.T_rbac_user
	err := ob.Raw("select * from t_rbac_user where username = ? and password = ?",username,passwordmd5).QueryRow(&user)

	if err!=nil {
		log.Println(err)
		res :=common.Error(-1,"登录失败")
		o.Data["json"]=res
		o.ServeJSON()
	}

}