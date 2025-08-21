Leither
========
本文档包括以下内容：系统介绍、下载安装、功能体验、应用示例等

## 系统介绍
Leither是一个数据和应用的容器系统。系统通过容器间的互通互联共同构建出一个去中心化的互联网，应用和数据可以在网络间流动，应用之间，容器之间，可以协作出复杂的生态应用。

### 一个轻量级的容器系统
* **多平台适配**
    系统目前大约20M,支持常见操作系统（windows,linux,darwin等），可以运行在服务器、pc、手机、nas、路由等常见的设备上.
* **功能完备**
    系统实现了认证体系、应用体系、文件和数据库系统，域名和网络系统。
    系统提供200多个api接口，可以实现常见的互联网应用。
* **多种交互**:
    容器内可以运行go/xgo,js,lua等语言
    也可以使用rpc方式开发传统应用，支持常见四十种语言
    也可以使用命令行浏览器等方式访问容器

### 模仿大自然和人类社会
大自然和人类社会是这个世界最复杂的系统。这两个复杂的系统核心的是由基因和模因构成的。Leither参考了这两个系统的设计，提出了弥媒的概念。
* **基因gene**
    大自然由各种生命构成,每个生命都是一系列基因的容器
    生命通过遗传和变异进行演化
    在和环境的交互过程中，适合环境的基因传承下去，不适合的淘汰下去
* **模因meme**
    智人之后，人类进化出一套新的机制。
    人的大脑是一个容器，装载着一系列的信息和知识，我们称之为弥因(meme)
    有些是知识，指导我们改造这个世界
    有些是信仰，构建出宗教民族
    有些是规则，构建出各种组织形态
    有些是价值，构建出不同的人生目标
    不同价值的人，构建出复杂的社会
* **弥媒mimei**
    参考基因和弥因，系统增加了核心模块弥媒
    弥媒是最基本的数据和应用单元
    每个弥媒都包含了一个文件系统或数据库
    弥媒可以指定缺省应用去展现和操作弥媒
    象网页一样，弥媒可以引用其它弥媒或资源去表达复杂的内容
    整个系统是一个弥媒容器，弥媒可以在容器这间流转
    容器之间和弥媒之间都可以互相协作，构建出更复杂的形态
    如同大自然和人类的社会一样

### 全面拥抱AI走向新文明
* **大模型的机会和危机**  
GPT3.5之后，大模型不断有能力涌现，获得了飞速的发展，作为辅助工具，给很多人的工作带来了效率的提升。
但另一方面，越来越多的行业和公司出现了使用大模型替代人的案例，这造成了大量的失业，特别是互联网相关的行业。
分析其原因，大企业掌握完整的生产资料和完整的业务链，员工只是整个链条中的可替代一环。
作为应对，大部分人应该使用大模型去拆除企业的壁垒，去反向争取属于自己的机会。

* **构建自己的生态网络**  
在Leither中，用户可以构建或加入一个去中心化的网络，可以从网络中获取或发布应用服务和信息。
整个网络中，信息（包括应用）的生成、存储、发布、展现、传播都是完全去中心化的。
各个环节的规则都是透明的，用户都可以自建，没有哪一环是不可替代的。  
整个网络的特点是：用户构建、用户控制、用户享有。  

* **给大模型搭建工作台**  
把容器作为工作台提供给大模型，然后然后指定大模型的工作内容
在这种情况下，大模型是你的员工，所有的工作内容和成果，都代表你的意志，产权和收益都是你的。
这些内容可以在网络间流动，认同的人越多，流传的就越广，价值越大，流传的越久。
你提出的规则可以附加在群体应用上，价值越大，认同的人越多，就越能构建出更复杂更大规模的应用.

* **找到属于自己的机会**
一方面，在传统的互联网中，越来越多的机会被少数平台所垄断，这些平台通过规则把属于用户和商家的数据收集起来，用于获利。这些业务都比较适合结合大模型去肢解。
另一方面，传统的互联网都习惯把所有数据都收集到云端统一管理，而基于用户端的多样化需求都没有得到充分满足。有了大模型的加持，以外一些难以实现的需求可以低成本快速实现。

## 下载安装
### 角本安装
支持bash的终端下可以直接执行以下角本安装在当前目录:
```shell
curl -fsSL http://vzhan.cn/install.sh | bash
```

### 手工安装
到以下网址下载程序到本地
<a href="http://vzhan.cn/dl.html"> 各版本的应用</a>  

把指定版本的程序改为Leither  

设置好Leither的运行权限  
```shell
chmod +x ./Leither
```

### 初始化
第一次运行需要初始化  
可以通过-p指定端口，不指定缺省为4800.如果是80端口，需要root权限，命令都要加sudo 
可以通过-b指定网络入口，不指定为：mimei.org vzhan.cn  
生成的信息保存在Systemvar.json中，可手工修改。
```shell
Leither init -p 4800 -b vzhan.cn
```

### 管理服务
后台服务方式启动  
```shell
./Leither run -d
```

关闭服务
```shell
./Leither stop
```

## 功能体验
可以通过命令行，浏览器，api开发应用等多种方式体验功能  
下面通过命令行展示系统功能 
### **网络**
节点启动之后，通过引导节点进入网络  
显示本节点id  
```shell
./Leither swarm id
```
返回：
```shell
9I6JPEqsxWHN7dr2WG9C0CJ-VnN
12D3KooWBiuFhtpQL2fs3CasdDZ2yZsHHGWdbGTEuVM3BHaj4Aj
```
返回有长短两个id
Leither中所有的资源都有一个27长度的id,生成规则类似比特币的钱包地址  
长的id为兼容ipfs网络，生成方式和ipfs网络规则完全相同（multihash）

显示本机地址
```shell
./Leither swarm local
```

返回：  
```shell
/ip4/192.168.3.7/tcp/4800
/ip4/60.186.9.237/tcp/4800
/ip6/240e:390:86b:f630:2900:1e4:50f8:cd6a/tcp/4800
```
Leither网络的地址为multiaddr格式

显示附近的网络节点
```shell
./Leither swarm addrs
```
返回:
```shell
Ngacq50-IRX_DfcndFwZ0c9S0Nh (3)
        /ip4/192.168.3.7/tcp/4800
        /ip4/60.186.9.237/tcp/4800
        /ip6/240e:390:86b:f630:2900:1e4:50f8:cd6a/tcp/4800
l86HuY4FuRDezLEPHOHBjnaQczp (2)
        /ip4/172.31.47.58/tcp/80
        /ip4/18.222.243.60/tcp/80
...
略
....
```


查找节点
```shell
./Leither dht findpeer tNP93yuZhNXd-om4izWQkYHfS50
```
返回：

```shell
/ip4/99.79.46.219/tcp/80
/ip6/2600:1f11:ec1:3001:4d57:55c2:aec6:e279/tcp/80
/ip4/10.0.17.253/tcp/80
```
### **ipfs文件**
ipfs是知名的去中心化文件存储系统  
Leither网络支持ipfs协议，兼容ipfs网络。

**添加ipfs文件到网络**
```shell
./Leither ipfs add test.txt
```
输出结果如下:
```shell
ipfs add ok  /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

通过网址查看文件
```shell
curl 127.0.0.1:4800/ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
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
```shell
./Leither mimei create
```

返回:
```shell
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
```shell
./Leither mimei setcid RXN74QNeiY08LRSaoeQhx3nOLTC QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
```

返回:
```shell
mid= RXN74QNeiY08LRSaoeQhx3nOLTC
cid= QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
MiMeiSetCid ver= 1
```

这时候弥媒的最后版本为1  
版本1指向了上面的ipfs文件


**直接填充文件到弥媒**  
```shell
./Leither mimei add RXN74QNeiY08LRSaoeQhx3nOLTC Leither.txt
```

返回:
```shell
add /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2 50773
MiMeiAdd cid= /ipfs/QmWiEp87XKT5CLfSGiEeGAgMobXuWVz6n5e8dXv82Uu4U2
MiMeiAdd ver= 2
```
这时候弥媒的最后版本为2  
版本2指向了上面的ipfs文件


**弥媒发布**  
发布弥媒信息到网络
```shell
./Leither mimei publish RXN74QNeiY08LRSaoeQhx3nOLTC
```

返回：   
```shell
MiMeiPublish mids= [RXN74QNeiY08LRSaoeQhx3nOLTC] EOL = 168h
MiMeiPublish ok
```
这时候弥媒的信息发布到了网络上  
所有保存弥媒数据的支撑节点会实时更新相应内容。  
  
  
**查看弥媒信息**  
查询本地和网络上的弥媒信息  
```shell
./Leither mimei  show RXN74QNeiY08LRSaoeQhx3nOLTC
```

返回：

```shell
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
```shell
./Leither mimei  get RXN74QNeiY08LRSaoeQhx3nOLTC
```
返回：  
```shell
.......
文件内容略
.......
```


**弥媒同步**  
从网络或指定节点上同步弥媒    
在另一个节点上输入下面指令  
```shell
./Leither mimei sync RXN74QNeiY08LRSaoeQhx3nOLTC
```  
返回:  
```shell
MiMeiSync mid = RXN74QNeiY08LRSaoeQhx3nOLTC
mimei sync ok
```  
弥媒信息和数据就同步到了本地节点  

**弥媒支撑**  
提供弥媒数据的节点我们叫支撑节点，或者叫数据提供者
provide是向网络广播：本节点提供这个弥媒的所有数据  

```shell
./Leither mimei provide RXN74QNeiY08LRSaoeQhx3nOLTC
```  
返回:
```shell
MiMeiProvide cids= [RXN74QNeiY08LRSaoeQhx3nOLTC]
MiMeiProvide ok
```  

查看支撑节点  
```shell
./Leither mimei findprovs RXN74QNeiY08LRSaoeQhx3nOLTC
```
返回:
```shell
MiMeiFindProvide mid= RXN74QNeiY08LRSaoeQhx3nOLTC
mimei findprovs  ver=2
mimei findprovs  addrs=[{Ngacq50-IRX_DfcndFwZ0c9S0Nh [/ip6/240e:390:86b:f630:2900
:1e4:50f8:cd6a/tcp/4800 /ip4/192.168.3.7/tcp/4800 /ip4/60.186.9.237/tcp/4800]}]
mimei findprovs ok
```  
用户发布新的弥媒数据时，所有在线的支撑节点的数据会实时同步
**弥媒使用**  
除了api和命令直接读取弥媒中的数据  
弥媒可以指定应用进行展示    
详细参见后面的“应用”和“域名显示弥媒”  
### **应用**
开发者可以根据api，开始发传统应用，具体可以参考api手册。
下面介绍的是Leither体系的容器应用。

**上传到容器**  
上传一个应用到容器，这个应用的当前版本就更新为当前上传的内容。
```shell
./Leither lapp uploadapp -i dist/dav -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```  

**应用备份**
把应用的当前版本备份到last版本，备份后的版本为只读状态。
```shell
./Leither lapp backup -a dav -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```    

**查看应用信息**  
```shell
./Leither lapp showapp -a dav -v cur -p newuserforlogin.ppt -n http://127.0.0.1:4800/
```  

**发布应用到网络**  
```shell
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
```shell
./Leither mimei sync htpoEXiE6IlCAqVbCvjvkY_XNfu
```  
返回:  
```shell
MiMeiSync mid = htpoEXiE6IlCAqVbCvjvkY_XNfu
mimei sync ok
```  
应用信息和数据就同步到了本地节点  

**指定弥媒缺省应用**  
弥媒创建的时候，可以指定缺省应用
```shell
./Leither mimei create -a <应用id>
```  
打开弥媒的时候，会缺省使用指定的应用  
也可以在打开弥媒的时候再指定应用

### **域名显示弥媒**
通过域名节点，可以把用户家里的节点通过域名展示给其他人。  
整个体验象idc机房里主机一样

以下链接用缺省应用打开一个弥媒，弥媒id为dJM6X7OTmJXbGqPQaFdAZ3kGpBl  
http://www.vzhan.cn/tpt/dJM6X7OTmJXbGqPQaFdAZ3kGpBl/  


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

## 应用示例
