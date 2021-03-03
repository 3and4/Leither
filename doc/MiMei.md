弥媒
========  
弥媒是用于描述数据和应用  

<!--传统的中心化互联网，应用和数据是耦合在一起，由单一集团统一规则运行维护。  
在去中心化互联网中，

为了解决节点之间的数据协同问题,节点  
-->

### 一、弥媒基础操作
#### 1.1、创建打开
MMCreate(sid, appid, ext, mark string, tp byte, right uint64) (mid string, err error)  
MMOpen(sid, mid, ver string) (string, error)  

#### 1.2、备份和版本
MMBackup(sid, mid, memo string) (ver string, err error)  
MMRestore(sid, mid, ver string) error  
MMDelVers(sid, mid string, vers ...string) (int64, error)  
MMRelease(sid, mid, ver string) (string, error)   
  
#### 1.3、对象引用
MMAddRef(sid, mid string, fileids ...string) (int, error)  
MMDelRef(sid, mid string, fileids ...string) (int, error)  
MMGetRef(sid, mid string) (ret map[string]int, err error)  

### 1.4、权限  
MMSetRight(sid, mid, member string, right uint64) error  

### 二、弥媒文件
本文描述系统中的文件对象 

#### 2.1、文件对象  
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
  
#### 2.2、打开文件
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


#### 2.3、操作文件  
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



### 三、数据库
数据库的底层有两种，一种是基于LevelDB,一种是基于BoltDB。  
两种数据库都进行过底层改造。  
LevelDb用于当前版本，可写可读，一致性是基于时序。  
写事务时会检查改动的部分是否有第三方同时修改过相应的内容。  
有改动则提示错误，通知调用端重新执行相关操作，避免写冲突。
BoltDb用于历史版本，只用于读。

Api参考Redis  
可以操作字符串，哈希表，列表，集合，有序集五组数据类型  
支持事务    

#### 3.1、事务 
Begin(dbsid string, timeout int) error  
Commit(dbsid string) error  
Rollback(dbsid string) error 

#### 3.2、字符串 
Set(dbsid, key string, value interface{}) error  
Get(dbsid, key string) (interface{}, error)  
Del(dbsid string, key ...string) (int64, error)  
Incr(dbsid, key string) (int64, error)  
IncrBy(dbsid, key string, delta int64) (int64, error)  
Strlen(dbsid, key string) (int64, error)  
  
#### 3.3、哈希表   
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

#### 3.4 列表
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

#### 3.5 集合
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

#### 3.6 有序集
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

