任务模块
========  
## 一、根据代码和参数运行角本 
RunScript(script string, glob map[string]string) (ret interface{}, err error) 
## 二、注册角本，生成角本ID 
RegisterScript(sid, script string, info *ScriptInfo) (id string, err error)  
type ScriptInfo struct {  
	ID   string  
	Name string  
	Memo string  
}  

## 三、通过角本ID运行角本 
RunByScriptID(sid, scriptId string, glob map[string]string) (ret interface{}, err error)
## 四、设置后台任务 
TaskRunInBackground func(sid string, setting TaskSetting) error   
type TaskSetting struct {  
	Name     string //缺省唯一  
	ScriptID string //运行的角本id  
	Start    int64  //0，表示立刻执行，正表示unix秒数，负值以后定义，目前考虑事件触发：｛login，消息等｝  
	Period   int    //单位秒，-1表示每月，-2表示每年；其它数值可以对应别的特殊信息  
	Param    map[string]string  
	Status   byte //0 执行中，1 暂停，/*2，等待任务*/  
	//后续添加权限相关  
}  

