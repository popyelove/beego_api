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
	json2 "encoding/json"
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
func (this *WebController) CheckLogin() {
	user_id :=this.GetSession("rbac_user_id")
	if user_id==nil {
		res :=common.Error(11,"缺失 openid")
		this.Data["json"]=res
		this.ServeJSON()
	}else {
		res :=`{"ret":0,"msg":"\u6210\u529f","data":{"privilege":[{"name":"\u516c\u5171\u57fa\u7840","status":0,"menu":1,"icon":"user","children":[{"name":"\u89d2\u8272\u7ba1\u7406","status":0,"url":"\/web\/public\/role","menu":101,"children":[{"name":"\u65b0\u589e","status":1,"menu":10101,"path":"Rbac_Role\/add"},{"name":"\u7f16\u8f91","status":1,"menu":10102,"path":"Rbac_Role\/edit"},{"name":"\u67e5\u8be2","status":1,"menu":10103,"path":"Rbac_Role\/all"}]},{"name":"\u7528\u6237\u7ba1\u7406","status":0,"menu":102,"url":"\/web\/public\/user","children":[{"name":"\u65b0\u589e","status":1,"menu":10201,"path":"Rbac_User\/add"},{"name":"\u5220\u9664","status":1,"menu":10202,"path":"Rbac_User\/del"},{"name":"\u7f16\u8f91","status":1,"menu":10203,"path":"Rbac_User\/edit"},{"name":"\u89e3\u7ed1","status":1,"menu":10204,"path":"Rbac_User\/resetPassword"},{"name":"\u67e5\u8be2","status":1,"menu":10205,"path":"Rbac_User\/all"},{"name":"\u4fee\u6539\u5bc6\u7801","status":1,"menu":10206,"path":"Rbac_User\/resetPassword"},{"name":"\u6279\u91cf\u5bfc\u5165","status":1,"menu":10207,"path":"Rbac_User\/resetPassword"}]},{"name":"\u5ba1\u6838\u8bbe\u7f6e","status":0,"url":"\/web\/public\/audit","menu":103,"children":[{"name":"\u7f16\u8f91","status":1,"menu":10301,"path":"Agent\/update_agent_status"},{"name":"\u67e5\u8be2","status":1,"menu":10302,"path":"Agent\/get_all_agent"}]}]},{"name":"\u6e20\u9053\u7ba1\u7406","status":0,"menu":2,"icon":"appstore-o","children":[{"name":"\u6e20\u9053\u5217\u8868","status":0,"menu":201,"url":"\/web\/agent","children":[{"name":"\u65b0\u589e\/\u7f16\u8f91\u6e20\u9053","status":1,"menu":20101,"path":"Agent\/edit"},{"name":"\u67e5\u8be2\u6e20\u9053","status":1,"menu":20102,"path":"Agent\/all"},{"name":"\u8d39\u7387","status":1,"menu":20103,"path":"Agent_Payment\/all"},{"name":"\u6e20\u9053\u51bb\u7ed3","status":1,"menu":20104,"path":"Agent_Payment\/all"}]},{"name":"\u6e20\u9053\u5ba1\u6838","status":0,"menu":202,"url":"\/web\/audit\/agent","children":[{"name":"\u67e5\u8be2","status":1,"menu":20201,"path":"Agent\/audit_query"},{"name":"\u5ba1\u6838","status":1,"menu":20202,"path":"Agent\/audit"}]}]},{"name":"\u5546\u6237\u7ba1\u7406","status":0,"menu":3,"icon":"inbox","children":[{"name":"\u5546\u6237\u5217\u8868","status":0,"menu":301,"url":"\/web\/merchant","children":[{"name":"\u65b0\u589e\/\u7f16\u8f91\u5546\u6237","status":1,"menu":30101,"path":"Merchant\/edit"},{"name":"\u67e5\u8be2\u5546\u6237","status":1,"menu":30102,"path":"Merchant\/all"},{"name":"\u5ba1\u6838\u5546\u6237","status":1,"menu":30103,"path":"Merchant\/audit"},{"name":"\u5546\u6237\u652f\u4ed8\u914d\u7f6e","status":1,"menu":30104,"path":"Merchant_Payment\/all"},{"name":"\u5bfc\u51fa","status":1,"menu":30105,"path":"Merchant_Payment\/all"}]},{"name":"\u5546\u6237\u5ba1\u6838","status":0,"menu":302,"url":"\/web\/audit\/merchant","children":[{"name":"\u67e5\u8be2","status":1,"menu":30201,"path":"Merchant\/audit_query"},{"name":"\u5ba1\u6838","status":1,"menu":30202,"path":"Merchant\/audit"},{"name":"\u5bfc\u51fa","status":1,"menu":30203,"path":"Merchant\/audit"}]},{"name":"\u8349\u7a3f\u7bb1","status":0,"menu":303,"url":"\/web\/merchant\/draft","children":[{"name":"\u67e5\u8be2","status":1,"menu":30301,"path":"Merchant\/audit_query"},{"name":"\u7f16\u8f91","status":1,"menu":30302,"path":"Merchant\/audit_query"}]},{"name":"\u5546\u6237\u5ba1\u6838\u8bb0\u5f55","status":0,"menu":304,"url":"\/web\/audit\/merchant\/record","children":[{"name":"\u67e5\u8be2","status":1,"menu":30401,"path":"Merchant\/audit_query"}]}]},{"name":"\u4ea4\u6613\u7ba1\u7406","status":0,"menu":4,"icon":"pay-circle-o","children":[{"name":"\u6c47\u603b\u6570\u636e","status":0,"menu":401,"url":"\/web\/transaction\/total","children":[{"name":"\u67e5\u8be2","status":1,"menu":40101,"path":"Transaction_Flow\/total_all"}]},{"name":"\u8ba2\u5355\u67e5\u8be2","status":0,"menu":402,"url":"\/web\/transaction\/flow","children":[{"name":"\u67e5\u8be2","status":1,"menu":40201,"path":"Transaction_Flow\/flow_all"},{"name":"\u5bfc\u51fa","status":1,"menu":40202,"path":"Transaction_Flow\/flow_export"}]},{"name":"\u4ee3\u4ed8\u8bb0\u5f55","status":0,"menu":403,"url":"\/web\/transaction\/PayedRecord","children":[{"name":"\u67e5\u8be2","status":1,"menu":40301,"path":"Transaction_Flow\/flow_all"},{"name":"\u5bfc\u51fa","status":1,"menu":40302,"path":"Transaction_Flow\/flow_export"}]}]},{"name":"\u7ed3\u7b97\u7ba1\u7406","status":0,"menu":5,"icon":"calculator","children":[{"name":"\u5546\u6237\u624b\u7eed\u8d39\u7edf\u8ba1","status":0,"menu":501,"url":"\/web\/statistics\/merchant","children":[{"name":"\u67e5\u8be2","status":1,"menu":50101,"path":"Merchant_Profit\/all"}]},{"name":"\u6e20\u9053\u5206\u6da6\u7edf\u8ba1","status":0,"menu":502,"url":"\/web\/statistics\/agent","children":[{"name":"\u67e5\u8be2","status":1,"menu":50201,"path":"Agent_Profit\/all"},{"name":"\u8be6\u60c5","status":1,"menu":50202,"path":"Agent_Profit\/one"}]},{"name":"\u5bf9\u8d26","status":0,"menu":503,"url":"\/web\/statistics\/balance","children":[{"name":"\u67e5\u8be2","status":1,"menu":50301,"path":"Agent_Profit\/all"}]},{"name":"\u5212\u8d26","status":0,"menu":504,"url":"\/web\/statistics\/agent\/distribute","children":[{"name":"\u67e5\u8be2","status":1,"menu":50401,"path":"Agent_Profit\/all"}]}]},{"name":"\u7ba1\u7406\u4e2d\u5fc3","status":0,"menu":6,"icon":"bars","children":[{"name":"\u4e0b\u8f7d\u4e2d\u5fc3","status":0,"menu":601,"url":"\/web\/management\/download","children":[{"name":"\u67e5\u8be2","status":1,"menu":60101,"path":"Statistics_Agent\/result_export"}]},{"name":"\u8054\u884c\u53f7\u67e5\u8be2","status":0,"menu":602,"url":"\/web\/management\/bankNo","children":[{"name":"\u67e5\u8be2","status":1,"menu":60201,"path":"Statistics_Agent\/result_export"}]}]}],"userInfo":{"rbac_user_id":"1","username":"admin","created":"2017-11-23 11:37:10","name":null,"address":null,"updated":"2017-11-23 11:37:10","mobileRole":"9","mchRole":{"boss":"16","clerk":{"3":"17","2":"18","1":"19"}},"agent_id":"1","apartment":"\u7814\u53d1\u90e8","agent_roles":["8"],"isMobileUser":false,"uid":"1"}}}`
		this.Ctx.WriteString(res)
	}

}
func CheckPrivilegeSetInfo(userid int,usertype int)  {

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
	log.Println(user.Username)
	log.Println(user.Rbac_user_id)
	o.SetSession("rbac_user_id",user.Rbac_user_id)
	o.SetSession("username",user.Username)
	res :=common.Success()
	o.Data["json"]=res
	o.ServeJSON()
}
// @Title Roleall
// @Description 角色列表
// @Success 200
// @router /role/all [post]
func (this *WebController) Roleall() {
	req :=this.Ctx.Input.RequestBody
	params,err:=simplejson.NewJson(req)
	if err!=nil {
		log.Println(err)
		return
	}
	type RbacUser struct {
		Rbac_role_id int `json:"rbac_role_id"`
		Name string		`json:"name"`
		Used int	`json:"used"`
		Comment string	`json:"comment"`
		Created string	`json:"created"`
		Author_id int	`json:"author_id"`

		Rbac_user_id int	`json:"rbac_user_id"`
		Author_name string	`json:"author_name"`

	}
	var users []RbacUser

	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	name,_:=params.Get("name").String()
	status,_:=params.Get("status").String()
	qb.Select(
		"t_rbac_role.rbac_role_id",
		"t_rbac_role.name",
		"t_rbac_role.used",
		"t_rbac_role.comment",
		"t_rbac_role.created",
		"t_rbac_role.author_id",
		"t_rbac_user.rbac_user_id",
		"t_rbac_user.username as author_name").From("t_rbac_role").
			InnerJoin("t_rbac_user").On("t_rbac_role.author_id = t_rbac_user.rbac_user_id")
	qb.Where("1=1")
	if name!=""{
		qb.And("name like "+`"%`+name+`"`)
	}
	if status!=""{
		qb.And("t_rbac_user.used ="+status)
	}
	// 导出 SQL 语句
	sql := qb.String()
	log.Println(sql)
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Using("psys") // 默认使用 default，你可以指定为其他数据库
	o.Raw(sql).QueryRows(&users)
	ret :=common.SuccessData(users)
	json,_:=json2.Marshal(ret)
	this.Ctx.WriteString(string(json))

}
// @router /test [post]
func (this *WebController) Test()  {
	c = make(chan int)
	go ready("tee",2)
	go ready("chao",5)
	fmt.Println("i am waiting but not too long")
	a :=<-c
	b :=<-c
	fmt.Println(a,b)
}
var c chan int
func ready(w string,sec int)  {
	fmt.Println(w,"is ready")
	c<-1
}