App
========  
应用部分Api  
本Api基于MiMei部分完成，同时被lpki命令集使用  
应用本质是一种特殊的弥媒，共用弥媒相关的api  

### 一、上传应用
上传一个应用
```golang
UploadApp(sid, fileid, tp string) (app *App, err error)
```
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取|
|fileid|文件id|通过MFTemp2MacFile生成 
|tp|文件类型|支持zip和tar两种类型，缺省是tar

|返回值|描述|
|--|--|
|app|应用信息|  

App类型定义如下：
```golang
type App struct {
	ID      string `json:"id"` //应用id
	Name    string //应用名称
	Author  string //开发者id
	Release string //release
	Last    string //last ver
}
```

### 二、更新应用文件
```golang
UploadAppfile(sid, AppName string, filename, fileid string) error
```
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取|
|AppName|应用名|
|filename|文件名| 
|fileid|文件id|通过MFTemp2MacFile生成 

|返回值|描述|
|--|--|
|app|应用信息| 
  
<!--
还不够完备，先注解掉
### 三、卸载应用
```golang
UninstallApp(sid, AppName string) error
```
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取|
|AppName|应用名|
-->  
### 其它  
应用本质是一种弥媒，其它操作参考弥媒api  
查询信息的Api参考变量Api  
