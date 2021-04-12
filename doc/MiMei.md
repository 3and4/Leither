Leither&弥媒
========  
Leither是一个去中心化的云操作系统  
弥媒是系统中的数据和应用单元  
弥媒通过应用生成、管理、存取、展现、传播  
弥媒基于内容主题创建，每一个弥媒都有一个唯一ID  
用户可以设置弥媒的各种权限  
弥媒和弥媒之间通过引用进行关联  
弥媒可以运行保存在一个或多个节点上  
象Git一样，弥媒有版本的概念  
版本主要是解决数据的一致性问题  
节点之间的数据库可以实时同步变化  
也可以在生成确定版本的时候再进行同步   
弥媒目前描述的数据类型有数据库和文件，后续会支持流
文件系统是特殊的类型，基于数据库和文件完成  
每个弥媒都可以有独立的数据库和文件系统空间  
弥媒通过不断的修改备份完成内容的继承进化  
用户也可以用分支的方式完成弥媒内容的变异  

### 一、背景
#### 1.1、互联网现状
**1.1.1 大数据控制了整个社会**  
+ **互联网正在逐步成为社会的核心**  
互联网在当今社会越来越重要  
互联网一步步重构了整个社会  
2019年，中国线上交易额250万亿  
  
+ **互联网通过大数据影响社会**  
大数据的本质是收集数据和控制规则  
进而影响和控制个人行为，并因此创造暴利的商业模式  
  
**1.1.2 寡头控制了大数据**  
+ **最核心的数据已属于寡头**  
社交相关的数据归腾讯， 电商相关的数据归阿里  
其他重要的数据也分别落入两者相关的系统内  

+ **通过规则制造了层层壁垒**  
通过系统应用规则进行保护; 通过现行的法律法规保护  
通过生态潜规则进行保护  

+ **利益和机会属于寡头**  
90%的利润被AT系所获取，财富集中速度越来越快  
稍大一些的创业或投资机会，几无例外都落入了寡头手中  
阿里和腾讯市值超过10万亿人民币  

**1.1.3 绝大多数人被困数据牢笼**
+ **普通用户，任人宰割**  
没有隐私，是一系列的数据标签  
没有自主，是可以被引导控制的数学模型  

+ **中小企业，进退失据**  
在大数据面前，每个企业都变成了透明体，机会和利润被压到了边缘  
腾讯平台上的游戏厂商分成10%都不可得；大部分电商处在平衡线上  
 
+ **投资环境，机会丧失**   
新投资机会出现的时候，最终都不可避免的落入阿里和腾讯这样的寡头手中  
社会内卷化，我们很久没有听过普通人逆袭的故事了  

**全人类的智慧汗水结晶，被少数公司攫取果实**  
+ **美国军方发明了互联网，没有控制获利**  
互联网源自美国军方的阿帕网（国防部高级计划局网络）   
互联网的架构从创建那天就是开放的，去中心化的   

+ **运营商构建了国内网络，没有控制获利**   
国家通过国企业投入巨资，在全国范围通水通电通路通网  
国企运营商不计回报构建了世界上最大的互联网  
投入的4G、5G基站都是占世界一半以上  

+ **少数公司通过几个应用控制了互联网**    
国内互联网被阿里、腾讯系控制  
海外互联网被脸书、谷歌、亚马逊、微软等公司控制  

**1.1.4 构建去中心化互联网是当前时代的需求**

#### 1.2、计算机的发展历史
**文件——数据的模型**  
计算机诞生初期，最早的数据需求是应用  
后续随着应用功能复杂，应用也产生了大量的相关数据    
在unix中，设计了文件这样的数据载体，通过文件系统管理数据。  
一切皆文件，所有信息相关的设备都通过文件表达  
文件接口包括打开，读，写，关闭这样通过的数据接口。  
同时为了管理文件和存储介质，设计了文件系统

**MIME——数据类型**  
随着应用的增加，需要表达文件和应用之间的关联。  
在dos中有了扩展名的概念，这时候还是比较粗的文件类型表达  
到了windows中，扩展名已经可以和应用进行绑定  
在电子邮件的MIME中，对文件（附件）类型有了严格的定义    
目前的浏览器都对这个定义进行了支持

**数据库——数据间关联**  
为了处理复杂的数据关系，发明了数据库的概念。  
数据库还是封闭的数据集合，是一个体系内的关联关系   
缺点，数据是私产，不能自由流动，产生最大的价值  

**HTML——数据跨点关联**  
www的出现，从规则上定义了数据和数据之间跨文件跨主机的关联  
互联网诞生，信息极大爆炸  
数据可以跨节点关联，突破了主机的限制  
门户网站充分利用了这一特征  

**google——数据检索**  
信息大爆炸，产生了检索的需求  
Goolge提供了一种根据关键字进行内容检索的服务  
给数据附加了标签的属性  

**facebook——社交模型**  
最早的社交通讯工具是icp、msn、qq  
解决了人和人之间的实时消息通讯问题  
twitter提供了一个更快速简短的信息分享平台
myspace到facebook一个更系统的社交模型诞生了  
这个模型分四部分：
1. 我  
主要是管理用户个人相关的的信息
2. 朋友
用户的关系链  
3. 应用和服务  
用户使用的应用和服务
4. 消息流  
用户和外界的互动区，互动对象是朋友、应用、服务等

这个系统模型非常成功，后续出的互联网平台产品都在向这个架构靠拢  
典型代表是微信，微博，阿里来往，支付宝，一些企业的OA系统，都或多或少有这个框架的影子。


### 二、Leither & 弥媒  
信息熵是宇宙和人类发展过程中重要的一个物理量   
熵增定律为宇宙演变指明了一个方向。熵也被称为时间之箭    
薛定谔在《生命是什么》中提到：人活着就是在对抗熵增定律，生命以负熵为生    
生命把对环境的信息记录在了基因（gene）里
通过遗传适应之前的环境，通过变异筛选适应环境的变化  
人类有了意识之后，可以在大脑中对环境进行模拟调整，在大脑中处理相应的信息    
这些信息的基础单位被查德-道金斯称为弥母（meme）  
计算机诞生以后，信息的处理方式发性变化，需要一种新的概念来表达  
这种新的信息单位，我们称之为弥媒（mimei）  

#### 2.1、需求目标

构建基于用户的去中心互联网，设计原则如下：  
+ **一切以用户为中心**  
用户构建、用户控制、用户享有  
a new birth of freedom and that  internet of the people, by the people, for the people 

+ **规则自由开放**  
网络底层协议是开放的；  
底层操作系统规则是开放的  
规则开放，用户才有选择权，用户才有控制权

+ **数据自由流动**  
在用户许可、安全可靠的情况下，用户数据可以从一个平台流动到另一个平台，可以从一种形式转换成另一种形式。  
数据能自由流动，第三方才可以为用户提供丰富多采的应用和服务  


结合计算机和互联网的历史，我们可以列出更细的设计需求：
+ 支持文件，文件系统和数据库  
+ 数据和应用的关联  
+ 跨节点的内容关联  
+ 跨节点的检索  
+ 跨节点数据流动  
+ 去中心化的社交模型  
+ 轻量级的容器系统
+ 简洁的应用开发方式
+ 体验要和当前互联网兼容
+ 合法合规兼容当下的社会制度  


#### 2.2、数据和规则解耦
数据结构和算法，是计算机科学中重要的两个概念;  
数据结构描述了数据存放的方式，通常描述内存，数据库或文件中的对象
算法描述了怎么处理这些数据对象，打开，生成，存取，保存，检索，展现等  
下文中的**数据**是遵守同一数据结构的数据对象， **规则**指的是使用这些数据的功能模块或应用。

在中心化的互联网中，寡头通过规则把用户数据绑定在了寡头的空间。 数据和规则耦合在了一起。要解放数据，首先要把数据和规则进行解耦。  
我们先分析一下规则和数据的关系和常见的应用模式

规则->数据
所有的互联网平台都是通过


数据->规则


我们观察现行的互联网案例，发现从数据->规则


**需求**  
背景中描述了的各元素要都能进行表达，需要支持以下特性：  
  

**数据和规则**  
在中心化的互联网中，寡头通过规则把数据绑定在了自己的数据空间。 数据和规则耦合在了一起
第一个需求就是解耦

我们观察现行的互联网案例，发现从数据->规则


用户在访问数据的时候


在之前的互联网案例中，windows中的文件类型



**需要对中心化的技术架构进行解耦**  
在中心化互联网中，所有的数据和应用规则耦合在一起，最终数据和规则都被少数寡头控制  
现有的互联网从下到上大多都是为中心化互联网设计的，功能强大而厚重  
从硬件设备到底层系统到应用体系都需要重构

+ **需要一个简洁的容器存放用户数据**  
数据要回归用户，首先需要一个容器   
这个容器必须是轻量级免维护的  
最好是存放在用户的个人空间  
这一点类似Solid中推荐的Pod  

+ **要能完成当前互联网的大部分应用**  


#### 2.3、弥媒



### 二、弥媒基础操作
#### 2.1、创建打开
MMCreate(sid, appid, ext, mark string, tp byte, right uint64) (mid string, err error)  
MMOpen(sid, mid, ver string) (string, error)  

#### 2.2、备份和版本
MMBackup(sid, mid, memo string) (ver string, err error)  
MMRestore(sid, mid, ver string) error  
MMDelVers(sid, mid string, vers ...string) (int64, error)  
MMRelease(sid, mid, ver string) (string, error)   
  
#### 2.3、对象引用
MMAddRef(sid, mid string, fileids ...string) (int, error)  
MMDelRef(sid, mid string, fileids ...string) (int, error)  
MMGetRef(sid, mid string) (ret map[string]int, err error)  

### 2.4、权限  
MMSetRight(sid, mid, member string, right uint64) error  

### 三、弥媒文件
本文描述系统中的文件对象 

#### 3.1、文件对象  
**弥媒当前文件**  
指的是弥媒文件的当前版本，当前文件是可读可写的  

**弥媒版本文件**  
指的是弥媒的备份文件，备份文件是只读的    
由当前版本文件备份生成， 同系统动态分配一个版本号指向这个备份文件。  
  
**Mac文件**  
以文件内容的mac id为标识的文件，Mac文件是只读的  
Mac文件可以被弥媒对象引用  
弥媒的版本文件就是一个版本号指向一个mac文件
Mac文件也可以由临时文件转换生成

**临时文件**  
由MFOpenTempFile创建生成，进行数据写操作之后。
由MFTemp2MacFile转换为mac文件  

**弥媒文件系统**    
系统内置的一个弥媒类型,描述一个弥媒的文件系统  
创建时系统代码如下：  
```golang
	fs.MTreeID, _, err = mc.CreateMiMei(
		dir.OwnerID,
		"",                 //需要一个系统app的id,目前用空表示
		api.MM_EXT_TREE,    //		"mmtree",
		api.MM_Auto,        //自动生成唯一id
		api.MM_File,        //内部保存为一个文件
		api.DataRight_Default_Group,
		false,
	)
```  
在webdav目录下，可以通过设置扩展名为.mmfs的配置文件配置或自动生成弥媒文件系统。  
可以通过MFOpenByPath打开文件系统中的对象  
注：目前是通过json文件的方式实现的，在文件比较多的时候性能较差，正在规划变成数据库方式实现  

**操作系统文件和目录**  
可以把操作系统中的文件目录link到webdav目录下
这时候可以通过MFOpenByPath函数打开操作系统中的文件和目录

**弥媒根目录**  
webdav目录是节点对外展示的弥媒总入口。  
可以把操作系统中的文件或目录link到这个目录
也可以生成一个配置文件指向节点内的一个弥媒对象  
  
#### 3.2、打开文件
**弥媒文件**
创建弥媒文件  
接口函数MMCreate，参数类型为api.MM_File
返回值为弥媒id  
打开弥媒文件  
接口函数为MMOpen(sid, mid, ver string) (string, error)
返回值为会话id,通过会话id可以操作文件内容
ver如果是"cur",可以进行写操作。  
其它版本文件为备份过的文件，只能进行读操作。  

**通过路径打开**    
这种方式打开的是弥媒文件系统或操作系统文件系统中的文件  
相关API：  
MFOpenByPath(sid, fsid, path string, flag int) (string, error)

**打开Mac文件**   
打开挂在某个弥媒下的mac文件  
相关API：  
MFOpenMacFile(sid, mid, fileid string) (string, error)

**打开临时文件**   
打开一个临时文件，用于读写操作
MFOpenTempFile(sid string) (string, error)  
读写操作完成之后转换为Mac文件    
MFTemp2MacFile func(sid, mid string) (string, error)  

**关闭文件**  
除临时文件外，其它文件操作完成后都需要关闭文件  
考虑到远程操作文件经常会有异常掉线情况。   
操作对象的会话id在超时情况下都会自动关闭文件  


#### 3.3、操作文件  
**对象方式读写**  
MFSetObject(fsid string, obj interface{}) error  
MFGetObject(fsid string) (interface{}, error)

**字节数组读写**  
MFSetData(fsid string, data []byte, start int64) (count int, err error)  
MFGetData(fsid string, start int64, count int) ([]byte, error)  

**查询状态**   
MFGetSize(fsid string) (int64, error)
MFStat(fsid string) (*FileInfo, error)
MFIsExist(fsid, fileid string) (bool, error)  
MFGetMimeType(fsid string) (string, error)

**目录操作**  
MFReaddir(fsid string, count int) ([]*FileInfo, error)
  
**文件系统**  
FSFind(sid, mmfsid, path string) (*FindResult, error)  
FSMkDir(sid, mmfsid, path string) error  
FSRemoveAll(sid, mmfsid, path string) error  
FSStat(sid, mmfsid, path string) (*FileInfo, error)  
FSRename(sid, mmfsid, oldpath, newFullName string) error

**其它**  
MFTruncate(fsid string, size int64) error  
MFCopy(fsid, dst, src, srcVer string) error  



### 四、数据库
数据库的底层有两种，一种是基于LevelDB,一种是基于BoltDB。  
两种数据库都进行过底层改造。  
LevelDb用于当前版本，可写可读，一致性是基于时序。  
写事务时会检查改动的部分是否有第三方同时修改过相应的内容。  
有改动则提示错误，通知调用端重新执行相关操作，避免写冲突。
BoltDb用于历史版本，只用于读。

Api参考Redis  
可以操作字符串，哈希表，列表，集合，有序集五组数据类型  
支持事务    

#### 4.1、事务 
Begin(dbsid string, timeout int) error  
Commit(dbsid string) error  
Rollback(dbsid string) error 

#### 4.2、字符串 
Set(dbsid, key string, value interface{}) error  
Get(dbsid, key string) (interface{}, error)  
Del(dbsid string, key ...string) (int64, error)  
Incr(dbsid, key string) (int64, error)  
IncrBy(dbsid, key string, delta int64) (int64, error)  
Strlen(dbsid, key string) (int64, error)  
  
#### 4.3、哈希表   
Hmclear(dbsid string, key ...string) (int64, error)  
Hdel(dbsid, key string, field ...string) (int64, error)  
Hlen(dbsid, key string) (int64, error)  
Hset(dbsid, key, field string, value interface{}) (int64, error)  
Hget(dbsid, key, field string) (interface{}, error)  
Hmget(dbsid, key string, fields ...string) ([]interface{}, error)  
Hmset(dbsid, key string, values ...FVPair) error  
Hgetall(dbsid, key string) ([]FVPair, error)  
Hkeys(dbsid, key string) ([]string, error)  
Hscan(dbsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)  
Hrevscan(dbsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)  
HincrBy(sid, key, field string, delta int64) (ret int64, err error)  

#### 4.4 列表
Lpush(dbsid, key string, value ...interface{}) (int64, error)  
Lpop(dbsid, key string) (interface{}, error)  
Rpush(dbsid, key string, value ...interface{}) (int64, error)  
Rpop(dbsid, key string) (interface{}, error)  
Lrange(dbsid, key string, start, stop int32) ([]interface{}, error)  
Lclear(dbsid, key string) (int64, error)  
Lmclear(dbsid string, keys ...string) (int64, error)  
Lindex(dbsid, k string, index int32) (interface{}, error)  
Llen(dbsid, k string) (int64, error)  
Lset(dbsid, k string, index int32, value interface{}) error  

#### 4.5 集合
Sadd(dbsid, key string, args ...string) (int64, error)  
Scard(dbsid, key string) (int64, error)  
Sclear(dbsid, key string) (int64, error)  
Sdiff(dbsid string, keys ...string) ([]string, error)  
Sinter(dbsid string, keys ...string) ([]string, error)  
Smclear(dbsid string, key ...string) (int64, error)  
Smembers(dbsid, key string) ([]string, error)  
Srem(dbsid, key string, m string) (int64, error)  
Sunion(dbsid string, keys ...string) ([]string, error)  
Scan(dbsid string, begin, match string, count int, inclusive bool, tp byte) (keys []string, err error)

#### 4.6 有序集
Zadd(dbsid, key string, args ...ScorePair) (int64, error)  
Zcard(dbsid, key string) (int64, error)  
Zcount(dbsid, key string, mins, maxs int64) (int64, error)  
Zrem(dbsid, key string, members ...string) (int64, error)  
Zscore(dbsid, key, member string) (int64, error)  
Zrank(dbsid, key, member string) (int64, error)  
Zrange(dbsid, key string, mins, maxs int) ([]ScorePair, error)  
Zrangebyscore(dbsid, key string, mins, maxs int64, offset int, count int) ([]ScorePair, error)  
Zremrangebyscore (dbsid, key string, mins, maxs int64) (int64, error)
Zrevrange(dbsid, key string, start, stop int) (ret []ScorePair, err error)  
Zrevrangebyscore (dbsid, key string, mins, maxs int64, offset int, count int) (ret []ScorePair, err error)  
Zmclear(dbsid string, key ...string) (int64, error)  
Zclear(dbsid, key string) (int64, error)  
ZincrBy(dbsid, key string, delta int64, member string) (ret int64, err error)  

