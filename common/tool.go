package common

func SuccessData(data map[string]interface{}) (interface{}) {
 //str :=`{"ret":0,"msg":"ok","data":[]}`
 str :=make(map[string]interface{})
 str["ret"]=0
 str["msg"]="成功"
 str["data"]=data
 return str
}
func Success() (interface{}) {
 str :=make(map[string]interface{})
 str["ret"]=0
 str["msg"]="成功"
 str["data"]=make(map[string]interface{})
 return str
}
func Error(errno int,msg string)(interface{})  {
 str :=make(map[string]interface{})
 str["ret"]=errno
 str["msg"]=msg
 str["data"]=make(map[string]interface{})
 return str
}
