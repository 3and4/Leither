Leither
========
Leither是一个超轻量级的去中心化的云操作系统  
系统能够以极低的资源实现常见的互联网应用，包括不限于网站，公众号，小程序，APP    
它的大小只有12兆，可以运行在pc、服务器、手机、NAS、路由、树莓派等各种设备上  
支持常见的操作系统，包括不限于windows、linux、freebsd、Darwin、Android     
系统实现了认证体系、应用体系、文件和数据库系统，域名和网络系统  
Api支持常见四十多种语言    

一个节点，你可以当成轻量级的主机使用，可以在家用网络上跑出idc主机效果  
两个节点，可以进行互助，包含不限于数据互备、容错和均衡，域名解析和路由  
多个节点，可以构建去中心化的分布式网络，网络上能够以去中化心化的方式实现常见的互联网应用。  

通过弥媒和应用体系，开发者能够以去中心化的方式肢解互联网平台业务  
通过组织和共识机制，组织者可以协调参与各方的利益，促使更多的人参与进来    
两者结合，可以从互联网寡头手中把互联网交还给用户  


### **安装运行**
下载程序到本地，然后解压  
<a href="./bin/README.md"> 各版本的应用下载</a>  
了解各文件细节参考<a href="./doc/Directory.md"> 系统目录结构</a>  


设置好Leither的运行权限  
```bash
chmod +x ./Leither
```

直接运行  
```bash
./Leither&
```


### **网络**
节点启动之后，通过引导节点进入网络  
显示本节点id  
```bash
./Leither lpki id
```
返回：
```bash
Ngacq50-IRX_DfcndFwZ0c9S0Nh
```

显示本机地址
```bash
./Leither swarm local
```

返回：  
```bash
/ip4/192.168.3.7/tcp/8000
/ip4/60.186.9.237/tcp/8000
/ip6/240e:390:86b:f630:2900:1e4:50f8:cd6a/tcp/8000
```

显示附近的网络节点
```bash
./Leither swarm addrs
```
返回:
```bash
Ngacq50-IRX_DfcndFwZ0c9S0Nh (3)
        /ip4/192.168.3.7/tcp/8000
        /ip4/60.186.9.237/tcp/8000
        /ip6/240e:390:86b:f630:2900:1e4:50f8:cd6a/tcp/8000
l86HuY4FuRDezLEPHOHBjnaQczp (2)
        /ip4/172.31.47.58/tcp/80
        /ip4/18.222.243.60/tcp/80
...
略
....
```

查找节点
```bash
./Leither dht findpeer tNP93yuZhNXd-om4izWQkYHfS50
```
返回：

```bash
/ip4/99.79.46.219/tcp/80
/ip6/2600:1f11:ec1:3001:4d57:55c2:aec6:e279/tcp/80
/ip4/10.0.17.253/tcp/80
```

### **ipfs文件**
ipfs是知名的去中心化文件存储系统  
Leither网络支持ipfs协议，兼容ipfs网络。

**添加ipfs文件到网络**
```bash
./Leither ipfs add Leither.txt
```
输出结果如下:
```
IpfsAdd  /home/pi/sdb/Leither/Leither.txt
wa.sid= f6503de3a81ee92c2e0c8f04c1c1be71ce6af912
add /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2 50773
ipfs add ok  /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

可以在网络内任何一个节点，通过网址查看文件
```
curl 127.0.0.1:8000/ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
......
文件内容略
......
```
### **弥媒**
弥媒是一个信息容器,里面可以存放文件或数据库。    
弥媒实现了数据的存取、展现  
可以在各网络节点间流动  
可以通过资源引用来描述复杂的内容    
可以指定操作数据的关联应用  

传统互联网服务应用和数据都是高度耦合在一起的  
弥媒的主要目的是对这些内容进行解构再重构  
通过弥媒可以实现大部分传统互联网的功能      

**生成一个弥媒**   
```bash
./Leither mimei create
```

返回:
```
Create MiMei  ok 
mid= RXN74QNeiY08LRSaoeQhx3nOLTC
```
这里生成了一个弥媒  
id为RXN74QNeiY08LRSaoeQhx3nOLTC  
缺省类型是文件

弥媒生成之后，就可以往弥媒中填充数据  

**填充ipfs文件到弥媒**  
弥媒支持ipfs文件和文件系统

把ipfs文件放入弥媒
```bash
./Leither mimei setcid RXN74QNeiY08LRSaoeQhx3nOLTC QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

返回:
```bash
mid= RXN74QNeiY08LRSaoeQhx3nOLTC
cid= QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
MiMeiSetCid ver= 1
```

这时候弥媒的最后版本为1  
版本1指向了上面的ipfs文件


**直接填充文件到弥媒**  
```bash
./Leither mimei add RXN74QNeiY08LRSaoeQhx3nOLTC Leither.txt
```

返回:
```bash
add /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2 50773
MiMeiAdd cid= /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
MiMeiAdd ver= 2
```
这时候弥媒的最后版本为2  
版本2指向了上面的ipfs文件


**弥媒发布**  
发布弥媒信息到网络
```bash
./Leither mimei publish RXN74QNeiY08LRSaoeQhx3nOLTC
```

返回：   
```bash
MiMeiPublish mids= [RXN74QNeiY08LRSaoeQhx3nOLTC] EOL = 168h
MiMeiPublish ok
```
这时候弥媒的信息发布到了网络上  
所有保存弥媒数据的支撑节点会实时更新相应内容。  
  
  
**查看弥媒信息**  
查询本地和网络上的弥媒信息  
```bash
./Leither mimei  show RXN74QNeiY08LRSaoeQhx3nOLTC
```

返回：

```bash
MiMeiShow mid= RXN74QNeiY08LRSaoeQhx3nOLTC
Author  : Ngacq50-IRX_DfcndFwZ0c9S0Nh
AppType :
Ext     :
Mark    : 16650373773415379251557603663
Create  : 2022-10-06T14:22:57+08:00
Right   :7276704

from local
-------------------------------------------------------------       
2        QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
last     2

from net
-------------------------------------------------------------       
2        QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
last     2

conflict
-------------------------------------------------------------       
no conflict

compare version
-------------------------------------------------------------       
local version is same with net

mimei show ok
```

**查看弥媒内容**  
查询弥媒内容  
```bash
./Leither mimei  get RXN74QNeiY08LRSaoeQhx3nOLTC
```
返回：  
```bash
.......
文件内容略
.......
```


**弥媒同步**  
从网络或指定节点上同步弥媒    
在另一个节点上输入下面指令  
```bash
./Leither mimei sync RXN74QNeiY08LRSaoeQhx3nOLTC
```  
返回:  
```bash
MiMeiSync mid = RXN74QNeiY08LRSaoeQhx3nOLTC
mimei sync ok
```  
弥媒信息和数据就同步到了本地节点  

**弥媒支撑**  
提供弥媒数据的节点我们叫支撑节点，或者叫数据提供者
provide是向网络广播：本节点提供这个弥媒的所有数据  
```bash
./Leither mimei provide RXN74QNeiY08LRSaoeQhx3nOLTC
```  
返回:
```bash
MiMeiProvide cids= [RXN74QNeiY08LRSaoeQhx3nOLTC]
MiMeiProvide ok
```  

查看支撑节点  
```bash
./Leither mimei findprovs RXN74QNeiY08LRSaoeQhx3nOLTC
```
返回:
```bash
MiMeiFindProvide mid= RXN74QNeiY08LRSaoeQhx3nOLTC
mimei findprovs  ver=2
mimei findprovs  addrs=[{Ngacq50-IRX_DfcndFwZ0c9S0Nh [/ip6/240e:390:86b:f630:2900
:1e4:50f8:cd6a/tcp/8000 /ip4/192.168.3.7/tcp/8000 /ip4/60.186.9.237/tcp/8000]}]
mimei findprovs ok
```  

用户发布新的弥媒数据时，所有在线的支撑节点的数据会实时同步

**弥媒使用**  
除了api和命令直接读取弥媒中的数据  
弥媒可以指定应用进行展示    
详细参见后面的“应用”和“域名显示弥媒”  


### **应用**
用户可以根据api开发应用。api协议基于hprose协议，支持常见大部分开发语言。  
用户可以直接在自己的应用中直接使用系统功能。
对于html5应用，系统进行了特殊优化。

**应用备份**  
```bash
./Leither lapp backup -a dav -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```  
  

**上传到节点**
```bash
./Leither lapp uploadapp -i dist/dav -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```  

**查看应用信息**  
```bash
./Leither lapp showapp -a dav -v cur -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```  

**发布应用到网络**  
```bash
./Leither mimei publish htpoEXiE6IlCAqVbCvjvkY_XNfu -p newuserforlogin.ppt
```  
发布应用到网络之后，可以通过浏览器访问任何一个节点运行该应用


应用存放在dist/dav目录, 应用名就是目录名dav  
开发者签发的通行证是newuserforlogin.ppt  
节点地址是http://127.0.0.1:4800/  
上面命令可以做成角本，写在开发工具的配置文件中（例如package.json） 

**同步应用到节点**  
从网络或指定节点上同步应用    
在另一个节点上输入下面指令  
```bash
./Leither mimei sync htpoEXiE6IlCAqVbCvjvkY_XNfu
```  
返回:  
```bash
MiMeiSync mid = htpoEXiE6IlCAqVbCvjvkY_XNfu
mimei sync ok
```  
应用信息和数据就同步到了本地节点  

**指定弥媒缺省应用**  
弥媒创建的时候，可以指定缺省应用
```bash
./Leither mimei create -a 应用id
```  
打开弥媒的时候，会缺省使用指定的应用  
也可以在打开弥媒的时候再指定应用

### **域名显示弥媒**
通过域名节点，可以把用户家里的节点通过域名展示给其他人。  
整个体验象idc机房里主机一样

以下链接用缺省应用打开一个弥媒，弥媒id为dJM6X7OTmJXbGqPQaFdAZ3kGpBl  
http://www.leither.cn/tpt/dJM6X7OTmJXbGqPQaFdAZ3kGpBl/  


以下链接直接运行一个应用，应用id为htpoEXiE6IlCAqVbCvjvkY_XNfu  
http://www.leither.cn/tpt/,htpoEXiE6IlCAqVbCvjvkY_XNfu/  

注：要确保外网可以访问家用节点    
所有运营商都支持ipv6,电信运营商支持有动态的ipv4地址  
两者都要设置本地路由或光猫，确保外部可以访问  
如果本地没有可访问的地址，可以通过tunnel功能借助他人的网络对外提供服务。    

### **弥媒数据库**
弥媒创建的时候可以指定数据库类型  
数据库兼容绝大部分redis的功能  
html5中的数据库功能都对应实现
可以使用命令行和api操作这个数据库  
使用上面的备份和同步  
可以把数据库发布到网络上
支撑节点实时同步数据库的变化

### **均衡容错**  
**容错**  
分布式网络中的节点地址比较复杂  
有些是ipv4,有些是ipv6,有些是通过nat.  
用户访问域名会返回模板页，在模板页中检测最优的网络路径访问节点    
**均衡**  
通过provide机制， 可以有多个节点支撑用户的数据和应用  
用户通过域名或链接访问应用时
域名和路由节点会把所有支撑节点信息填充在模板页上  
通过浏览器，在用户本地选择最优的访问节点
从而实现均衡功能。
用户可以设置或修改模板页，从而指定不同的均衡策略


### **更多的功能使用**
可以通过以下方式和系统进行交互：  
系统应用界面、命令行、系统API、浏览器    
为了方便，文档通过命令行介绍和演示各模块的功能  
开发者可以通过api在自己的应用中使用这些功能  

细节参考 <a href="./api/Api.md"> Api文档</a>  

### **安全**  
节点在启动时，自动创建一对秘钥  
并根据秘钥生成用户id,这也是用户的身份标识  
lpki命令集中包含了密钥的管理功能  
也包括了一些身份认证相关的操作  


### **网络**  
服务启动的时候,会通过配置中的引导节点进入网络。  
swarm命令集中包含了节点地址连接过滤等功能    

整个网络是一个dht网络
dht命令集中包含了网络的读写功能  
也包含了节点查找的网络路由功能  


### **弥媒**  
系统中所有的应用和数据都是弥媒类型  
弥媒支持文件和数据库类型    
弥媒可以发布在网络上  
网络中的其他用户可以通过弥媒id直接访问到弥媒  
也可以把数据同步到本地节点    
用户也可以对他人的弥媒提供数据支撑  
弥媒的每次变化都会产生一个版本  
这些变化会实时更新到各个支撑节点（数据提供者）上。

再结合域名和负载均衡机制，可以实现容错和负载均衡功能  


mimei命令集相关的操作


系统兼容ipfs网络  
ipfs命令集中包含了ipfs常用命令  
files命令集包含了如何操作本地的ipfs文件系统

### **应用体系**  
弥媒创建的时候都会指定一个应用类型    
用户在打开弥媒的时候会调用这个应用展示弥媒  
应用本身也是一个弥媒  
lapp命令集管理应用的备份发布，也包含版本和域名管理。

### **域名和负载均衡**    
系统可以运行在家用节点  
家用网络优点是速度快，成本低  
家用网络缺点是没有稳定ip和80端口，稳定性也不如idc机房  
系统通过域名节点解决这个问题  
用户可以把资源（节点，应用，内容）绑定在一个域名上  
把域名解析指向一个域名节点  
第三方用户通过浏览器访问域名节点
域名节点返回一个极小（3k）的数据包
浏览器解析包里的数据和节点信息  
选择性能最好的节点提供服务给第三方用户  
浏览器筛选节点的过程就是一个容错和负载均衡的过程  
整个过程都是在浏览器中完成，没有额外使用域名节点的流量    

对于浏览者而言，没有页面跳转  
相当于直接访问用户节点   
整个体验和idc机房的服务相当  


### **相关技术原理介绍**  
<a href="./doc/MiMei.md"> 弥媒和应用体系</a>  
介绍如何通过去中心化应用的方式肢解互联网平台业务  
<a href="./doc/GongShi.md"> 组织和共识机制</a>  
如何通过共识机制，协调参与各方的利益，促使更多的人参与进来  

### **相关文档**  
项目的相关信息的ppt链接  
<https://zhuanlan.zhihu.com/p/330261216>


  
<a href="./doc/Pki.md"> 安全体系</a>  

<a href="./doc/Applition.md"> 应用体系</a>  

<a href="./doc/Setup.md"> 节点和应用安装</a>  
  
<a href="./opt/dav/"> 示例应用代码</a>  
  
<a href="./api/Api.md"> Api文档</a>  

<a href="./doc/MiMei.md">  弥媒和应用体系</a>  

<a href="./doc/GongShi.md"> 组织和共识机制</a>  
    
<a href="./doc/PaaS.md"> Leither vs 传统的云</a>   
系统可用于云服务，相较于虚拟主机和docker，CPU内存带宽等方面成本有数量级的降低。

