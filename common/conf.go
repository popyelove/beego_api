package common

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"fmt"
	"log"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
)

func Conf(key string,schmeid string) (res string){
	redis_key :=key+"_"+schmeid
	log.Println(redis_key)
	rds,err := redis.Dial(
		beego.AppConfig.String("redisnetwork"),
		beego.AppConfig.String("redishost")+":"+
		beego.AppConfig.String("redisport"),
		redis.DialPassword(beego.AppConfig.String("redispwd")))
	if err != nil{
		fmt.Println(err)
		return
	}
	value,err :=redis.String(rds.Do("get",redis_key))
	if err!=nil{
		log.Println(err)
	}
	if value!=""{
		return value
	}else {
		post :=make(map[string]string)
		post["key"]=key
		post["scheme_id"]=schmeid
		json_post,err :=json.Marshal(post)
		if err != nil {
			fmt.Println("error:", err)
		}
		client := &http.Client{}
		req_new := bytes.NewBuffer([]byte(json_post))
		request, _ := http.NewRequest("POST", "http://configcenter.coincard.cc/api/conf", req_new)
		request.Header.Set("Content-type", "application/json")
		response, _ := client.Do(request)
		if response.StatusCode == 200 {
			body, _ := ioutil.ReadAll(response.Body)
			js,err :=simplejson.NewJson(body)
			if err!=nil {
				log.Println(err)
			}
			//rds.Do("set",key+"_"+schmeid,js.Get("data").Get("value").MustString())
			if key!="all"{
				return js.Get("data").Get("value").MustString()
			}else {
				rds_arr, _ :=js.Get("data").Array()
				rds_map :=make(map[string]string)
				for i,_ := range rds_arr{
					arr :=js.Get("data").GetIndex(i)
					rds_map[arr.Get("key").MustString()]=arr.Get("value").MustString()
				}
				ret_json,_ :=json.Marshal(rds_map)
				return string(ret_json)
				//return js.Get("data").Array()
			}

		}
		return
	}
}

