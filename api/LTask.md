任务模块
========  
## 一、根据代码和参数运行角本 
RunScript(script string, glob map[string]string) (ret interface{}, err error) 
## 二、注册角本，生成角本ID 
RegisterScript(sid, script string, info *ScriptInfo) (id string, err error)
## 三、通过角本ID运行角本 
RunByScriptID       func(sid, scriptId string, glob map[string]string) (ret interface{}, err error)
## 四、设置后台任务 
TaskRunInBackground func(sid string, setting TaskSetting) error   
