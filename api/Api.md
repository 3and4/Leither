Api文档
========
**简介**   
本文档描述Leither系统的所有接口函数。  
接口包括：RPC方式的接口;命令行方式的接口;Http方式.

**1、命令行接口**
+ <a href="./LPki.md"> 安全认证(lpki)</a>    
+ 网络节点(swarm)
+ 分布式哈希表(dht)</a>
+ 弥媒管理(mimei)
+ <a href="./LApp.md"> 应用部署(lapp)</a>    
+ ipfs命令(ipfs)
+ 本地ipfs文件系统（files）


**2、RPC接口**  
采用Hprose引擎，支持常见四十多种语言  
文档会以golang和html5作为示例,其它语言类似  
后端支持页面模板和lua.   
   
接口包含以下部分：
+ <a href="./Auth.md"> 认证部分</a>  
+ <a href="./App.md" > 应用部分</a> 
+ <a href="./MiMei.md"> 弥媒部分</a>  
+ <a href="./Net.md" > 网络部分</a>    
+ <a href="./LTask.md"> 任务部分</a>    
+ <a href="./Message.md"> 消息部分</a>        
+ <a href="./VarAct.md"> 系统变量</a>    


**3、http接口**  
web入口  
webdav  
IP相关  
Mac相关    
系统变量获取    
应用入口  
文件相关  
pprof  
 