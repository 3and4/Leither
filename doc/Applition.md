应用体系  
========   
本文档所讲的应用，主要指可以通过浏览器访问的互联网应用。
通过调用API生成可执行的客户端程序和小程序等应用另行叙述。

### 一、应用的种类  
**1.1、静态页面模式**  
这种模式是通过静态页面访问应用  
需要在应用中写上api的地址，代码中根据地址生成api调用句柄   
示例如下：  
```javascript  
var hosturl = "ws://127.0.0.1:4800/ws/"
var baseurl = "http://127.0.0.1:4800/"
//生成操作句柄
var client = hprose.Client.create(hosturl, 
    ["Login","MFOpenByPath", "MFStat","MFReaddir", "MFGetData", "MFGetMimeType"]);

```

**1.2、传统Web容器模式**  
这种模式同Apache和Tomcat类似，访问节点地址，显示html目录中的应用。  
需要在应用中增加地址自动获取  
```javascript  
hosturl = "ws://" + window.location.host + "/ws/"
baseurl = "http://" +  window.location.host + "/"
//生成操作句柄
var client = hprose.Client.create(hosturl, 
    ["Login","MFOpenByPath", "MFStat","MFReaddir", "MFGetData", "MFGetMimeType"]);
```  

**1.3、弥媒应用模式**  
这种模式是Leither的推荐模式，应用以弥媒的形式发布到节点。   
支持镜像版本能功能，可以自由在节点之间流动。  
可以通过域名或链接直接访问  

需要在应用中增加地址自动获取  
```javascript  
p=window.getParam()
console.log("p=", p)
hosturl = "ws://" + p["ips"][p.CurNode] + "/ws/"
baseurl = "http://" + p["ips"][p.CurNode] + "/"
//生成操作句柄
var client = hprose.Client.create(hosturl, 
    ["Login","MFOpenByPath", "MFStat","MFReaddir", "MFGetData", "MFGetMimeType"]);
```  
getParam函数定义在模板页中，用于获取动态的节点信息

**1.4、兼容汇总模式**  
把上面三种情况合并起来，自动兼容
```javascript  
//缺省的地址，用于本地调试程序
var hosturl = "ws://127.0.0.1:4800/ws/"
var baseurl = "http://127.0.0.1:4800/"

//获取节点链接
if (window.getParam != null){
    p=window.getParam()
    console.log("p=", p)
    hosturl = "ws://" + p["ips"][p.CurNode] + "/ws/"
    baseurl = "http://" + p["ips"][p.CurNode] + "/"
} else if (window.location.host != ""){
    hosturl = "ws://" + window.location.host + "/ws/"
    baseurl = "http://" +  window.location.host + "/"
}

//生成操作句柄
var client = hprose.Client.create(hosturl, 
    ["Login","MFOpenByPath", "MFStat","MFReaddir", "MFGetData", "MFGetMimeType"]);
```  


### 二、应用管理  
lapp是一套指令集，用于管理应用。    
指令集是命令行形式，可以和开发工具整体在一起  

指令包括：  
+ 上传应用  
+ 更新文件  
+ 显示应用  
+ 域名设置  
+ 版本备份  
+ 版本发布  
+ 版本恢复  

<a href="../api/LApp.md"> 详细文档</a>    


### 三、应用版本   
应用的初始版本为1，每次发布的时候版本号增加1  
有三个特殊版本：  
**3.1、当前版本**  
编号cur  
表示是当前在开发的应用版本  
每次上传和更新的时候，操作的都是cur版本  
**3.2、最后版本**  
编号last  
表示是最后一次备份的版本  
备份应用是把当前版本转变为last版本  
**3.3、发布版本**  
编号release    
表示是稳定对外的版本

### 四、访问应用     
可以通过链接访问一个应用，格式如下：  
http://192.168.1.7:8000/entry?author=h7Nkwe4rfb1d15En8yQ65Lhhq9b&app=dav&ver=cur

|键名|描述|
|--|--|
|entry|应用入口|  
|appid|应用id|    
|author|开发者|  
|app|应用名|  
|ver|应用版本|  


注：appid参数等效于author加app  
  
### 五、应用相关的api
**5.1、上传应用**  
上传一个应用 
```golang
UploadApp(sid, fileid, tp string) (app *App, err error)
```
**5.2、更新应用文件**  
```golang
UploadAppfile(sid, AppName string, filename, fileid string) error
```
**5.3、删除应用版本**
```golang
MMDelVers(sid, mid string, vers ...string) (int64, error)  
```
**5.4、获取应用信息**  
获取应用信息  
```golang
err := stub.GetVarObj(api.ApiVarAPPInfo, &app, appName)
```
应用版本信息
```golang  
var ayVersions []string
err = stub.GetVarObj(api.ApiVarAPPVersions, &ayVersions, appName)
```  

所有的应用名  
```golang  
var ret []*api.App
err := stub.GetVarObj(api.ApiVarAPPInfos, &ret)  
```  
**5.4、域名设置** 
```golang  
SetDomain(sid, domain string, info map[string]string) error
```  

**5.5、版本备份** 
```golang  
MMBackup(sid, mid, memo string) (ver string, err error)
```  

**5.6、版本发布** 
```golang  
MMRelease(sid, mid, ver string) (string, error)
```  

**5.6、版本恢复** 
```golang  
MMRestore(sid, mid, ver string) error
``` 
