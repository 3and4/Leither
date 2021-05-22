任务模块
========  
## 一、根据代码和参数运行角本 
执行这个角本的时候，会先解析上下文信息到参数表glob    
从参数表中获取认证相关的信息（sid,用户名密码，通行证等）  
认证通过以认证用户身份执行角本  
不通过则以Guest身份执行角本  
```golang
RunScript(script string, glob map[string]string) (ret interface{}, err error) 
```
|参数|名称|说明|
|--|--|--|
|script|要执行的角本||
|golb|参数| 

|返回值|描述|
|--|--|
|ret|返回值|  


## 二、注册角本 
注册角本，生成角本ID  
```golang
RegisterScript(sid, script string, info *ScriptInfo) (id string, err error)  
```  

|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取|  
|script|要执行的角本|  
|info|角本参数|  

|返回值|描述|说明|
|--|--|--|
|id|角本id|用于RunByScriptID等函数

角本类型定义如下：  
```golang
type ScriptInfo struct {  
	ID   string  
	Name string  
	Memo string  
}  
```

## 三、通过角本ID运行角本 
```golang
RunByScriptID(sid, scriptId string, glob map[string]string) (ret interface{}, err error)
```
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取|  
|scriptid|要执行的角本的id||
|golb|参数| 

|返回值|描述|
|--|--|
|ret|返回值|    

## 四、设置后台任务 
设置一个后台任务  
```golang
TaskRunInBackground func(sid string, setting TaskSetting) error   
```
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取|  
|setting|任务设置||

```golang
type TaskSetting struct {  
	Name     string //缺省唯一  
	ScriptID string //运行的角本id  
	Start    int64  //0，表示立刻执行，正表示unix秒数，负值以后定义，目前考虑事件触发：｛login，消息等｝  
	Period   int    //单位秒，-1表示每月，-2表示每年；其它数值可以对应别的特殊信息  
	Param    map[string]string  
	Status   byte //0 执行中，1 暂停，/*2，等待任务*/  
	//后续添加权限相关  
}  
```