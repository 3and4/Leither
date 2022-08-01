弥媒
========  
弥媒是用于描述数据和应用，是系统中最重要的概念，类似Linux中的文件。      


## 一、基础功能  
弥媒有两种数据类型：文件和数据库。  
它们都有共同的api：创建、打开、权限、版本管理等。  
<div id="MMCreate"></div>    

### 1.1 创建弥媒   
在有弥媒权限的情况下，用户可以在当前节点上创建弥媒。   
mark如果为"{{auto}}" 每次产生新的弥媒id  
其它情况，参数相同生成相同的弥媒id,如果存在指向之前的弥媒对象。  
```golang
MMCreate(sid, appid, ext, mark string, tp byte, right uint64) (mid string, err error)
```
  
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取|
|appid|应用id|创建弥媒的应用，为空表示系统
|ext|扩展类型|同一应用内区分弥媒的类型
|mark|特殊标识|根据mark生成弥媒id,相同  
|tp|数据类型|
|right|权限|

|返回值|描述|
|--|--|
|mid|生成的弥媒 id

### 1.2 打开弥媒
打开一个存在的弥媒  
```golang
MMOpen(sid, mid, ver string) (msid string, err error)
```

|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mid|弥媒id|
|ver|弥媒版本|

|返回值|描述|
|--|--|
|msid|会话id,用于操作弥媒的Api

有三个特殊版本：  
|版本|说明|
|--|--|
|cur|当前版本，可写的版本必须是cur
|last|最后版本，备份操作把当前版本转变为last版本
|release|稳定版本，表示对外的正式版    


### 1.3 设置权限
设置弥媒权限  
```golang
MMSetRight(sid, mid, member string, right uint64) error
```  
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mid|弥媒id|
|member|针对成员
|right|权限

### 1.4 备份弥媒
对弥媒的当前版本进行备份，生成新的版本，并把特殊版本号last指向这个版本  
```golang
MMBackup(sid, mid, memo string) (ver string, err error)   
```  

|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mid|弥媒id|
|memo|备注|保留参数，还未启用
|right|权限|  
|返回值|名称|说明|
|--|--|--|
|ver|生成的版本|last指向这个版本

### 1.5 发布弥媒  
指定某个版本为当前的发布版
```golang
MMRelease(sid, mid, ver string) (newVer string, err error)
```  
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mid|弥媒id|
|ver|指定的版本
 

|返回值|描述|
|--|--|
|newVer|release指向的最新版本号

### 1.6 恢复弥媒
把某个版本复制到当前版本
```golang
MMRestore(sid, mid, ver string) error
```  
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mid|弥媒id|
|ver|弥媒版本|

### 1.7 清除版本  
清除不用的版本数据
```golang
MMDelVers(sid, mid string, vers ...string) (int64, error)
```  
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mid|弥媒id|
|vers|要清除的版本，变参，可指定多个
 

### 1.8 查询信息
弥媒的信息都是通过GetVar函数获取的，  
```golang
GetVar(sid, name string,args ...string) (interface{}, error)
```  
相关的一些变量定义，定义在api包
```golang
//弥媒部分
const (
	ApiVarMiMeiInfo     = "mminfo"     //弥媒信息
	ApiVarMiMeiVersions = "mmversions" //弥媒版本
)
```

**1.8.1查询弥媒信息**  
```golang
mminfo, err := api.GetVar(sid, "mminfo", mid) 
```  
弥媒信息的数据结构
```golang
type MiMeiInfo struct {
	AppType string    //应用类型
	Author  string    //作者
	Ext     string    //扩展信息
	Mark    string    //标识
	Create  time.Time //创建时间
	Right   uint64    //权限
}
``` 

### 1.9 引用对象
**1.9.1 添加引用**  
```golang
MMAddRef(sid, mid, fileids ...string)(count int, err error)
```  
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mid|弥媒id|
|fileids|要引用的对象id，变参，可指定多个    

|返回值|描述|说明|
|--|--|--|  
|count|成功的个数|  
 
   
**1.9.2 删除引用**  
```golang
MMDelRef(sid, mid string, fileids ...string) (count int, err error)  
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mid|弥媒id|
|fileids|要引用的对象id，变参，可指定多个    

|返回值|描述|说明|
|--|--|--|  
|count|成功的个数|  
```  

**1.9.3 获取引用**  
```golang
MMGetRef(sid, mid string) (refs map[string]int, err error)  
```  
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mid|弥媒id|

|返回值|描述|说明|
|--|--|--|  
|refs|引用数|｛文件id->引用数｝  
   

## 二、文件系统   
弥媒的文件系统  
### 2.1 弥媒文件的打开和关闭  
**2.1.1 通过路径打开文件**  
```golang
MFOpenByPath(sid, mmfsid, path string, flag int) (fsid string, err error)  
```  
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mmfsid|文件系统id|
|path|文件系统内的路径|
|flag|打开文件的选项|

|返回值|描述|说明|
|--|--|--|  
|fsid|资源会话id|用于操作资源的api  

  
**2.1.2 打开mac文件**  
```golang
MFOpenMacFile(sid, mid, fileid string) (fsid string, err error)  
```  
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取
|mid|文件系统id|
|fileid|mac文件的id|

|返回值|描述|说明|
|--|--|--|  
|fsid|资源会话id|用于操作资源的api  

**2.1.3 关闭文件**  
```golang
MFClose(fsid string) error  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|上面的MFOpen系列Api打开

**2.1.4 打开临时文件**  
```golang
MFOpenTempFile(sid string) (fsid string, err error)  
```
|参数|名称|说明|
|--|--|--|
|sid|会话id|通过Login获取

|返回值|描述|说明|
|--|--|--|  
|fsid|资源会话id|用于操作资源的api  
  
**2.1.5 临时文件转Mac文件**  
```golang
MFTemp2MacFile(fsid, mid string) (macid string, err error)  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|通过MFOpenTempFile获取
|mid|文件系统id|

|返回值|描述|说明|
|--|--|--|  
|macid|mac文件的id|  


### 2.2 文件读写  
**2.2.1 设置对象**  
```golang
MFSetObject(fsid string, obj interface{}) error  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|MFOpen系列Api获取
|obj|要设置的对象|

**2.2.2 获取对象**  
```golang
MFGetObject(fsid string) (obj interface{}, err error)  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|MFOpen系列Api获取

|返回值|描述|说明|
|--|--|--|  
|obj|返回的对象|  
  
interface{}类型主要是针对js这样的动态语言。  
静态语言这样的返回值会丢失类型。  
需要用另外的方式包装api  

**2.2.3 设置字节数组**  
```golang
MFSetData(fsid string, data []byte, start int64) (count int, err error)  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|MFOpen系列Api获取
|data|要写的字节数组|
|start|起始位置|正表示从头部偏移,负表示从尾部偏移,-1表示文件尾部

|返回值|描述|说明|
|--|--|--|  
|ret|返回的结果|  
  

**2.2.4 读取字节数组**  
```golang
MFGetData(fsid string, start int64, count int) (ret []byte, err error)  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|MFOpen系列Api获取
|start|起始位置|正表示从头部偏移,负表示从尾部偏移,-1表示文件尾部
|count|读取字节数|

|返回值|描述|说明|
|--|--|--|  
|ret|返回的结果|  
  
 
### 2.3 查询状态  
**2.3.1 获取文件长度**  
```golang
MFGetSize(fsid string) (size int64, err error)  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|MFOpen系列Api获取

|返回值|描述|说明|
|--|--|--|  
|size|文件长度|  

**2.3.2 获取文件的Mime类型**  
```golang
MFGetMimeType(fsid string) (tp string, err error)    
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|MFOpen系列Api获取

|返回值|描述|说明|
|--|--|--|  
|type|MiMe类型|  

**2.3.3 获取文件信息**  
```golang
MFStat(fsid string) (info *FileInfo, err error)  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|MFOpen系列Api获取

|返回值|描述|说明|
|--|--|--|  
|info|文件信息|  

**2.3.4 文件是否存在**  
```golang
MFIsExist(fsid, fileid string) (bExist bool, err error)  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|MFOpen系列Api获取
|fileid|文件id|  

|返回值|描述|说明|
|--|--|--|  
|bExist|文件信息|  


### 2.4 目录操作  
**2.4.1 读取目录下的文件信息**  
```golang
MFReaddir(fsid string, count int) (infos[]*FileInfo, err error)  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|MFOpen系列Api获取
|count|读取个数|不大于0表示读取全部  

|返回值|描述|说明|
|--|--|--|  
|infos|目录中的文件信息|  

### 2.5 其它文件操作  
**2.5.1 截断文件**  
```golang
MFTruncate(fsid string, size int64) error  
```
|参数|名称|说明|
|--|--|--|
|fsid|会话id|MFOpen系列Api获取
|size|文件长度|  |  


### 2.5 文件系统操作  
**2.5.1 根据路径查找**  
```golang
FSFind(sid, mmfsid, path string) (result *FindResult, err error)  
```  
|参数|名称|说明|
|--|--|--|
|sid|会话id|Login获取
|mmfsid|文件系统mid|  
|path|相对路径|  

|返回值|描述|说明|
|--|--|--|  
|result|查询结果|  

```golang
type FindResult struct {
	FSID     string //文件系统的id
	Left     string //未匹配的路径，匹配时为空
	FullName string //完整路径
}
``` 

**2.5.2 创建目录**  
```golang
FSMkDir(sid, mmfsid, path string) error  
```  
|参数|名称|说明|
|--|--|--|
|sid|会话id|Login获取
|mmfsid|文件系统mid|  
|path|相对路径|  

|返回值|描述|说明|
|--|--|--|  
|result|查询结果|  


**2.5.3 查询文件信息**  
```golang
FSStat(sid, mmfsid, path string) (info *FileInfo, err error)  
```
|参数|名称|说明|
|--|--|--|
|sid|会话id|Login获取
|mmfsid|文件系统mid|  
|path|相对路径|  

|返回值|描述|说明|
|--|--|--|  
|info|文件信息|  

**2.5.4 复制文件**  
```golang
FSCopy(fsid, dst, src, srcVer string) error  
```
|参数|名称|说明|
|--|--|--|
|sid|会话id|Login获取
|dst|目标路径|  
|src|源路径|  
|srcVer|源版本|  

**2.5.5 文件改名**  
```golang
FSRename(sid, mmfsid, oldpath, newFullName string) error
```
|参数|名称|说明|
|--|--|--|
|sid|会话id|Login获取
|mmfsid|文件系统mid|  
|oldpath|旧的文件路径|  
|newFullName|新的文件路径|  




## 三、数据库  
弥媒中的数据库  
### 3.1 事务  
**3.1.1 开始事务**  
```golang
Begin(dbsid string, timeout int) error 
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|timeout|延时|单位秒，超时自动Rollback

**3.1.2 事务递交**  
```golang
Commit(dbsid string) error  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取


**3.1.3 事务回滚**  
```golang
Rollback(dbsid string) error 
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取

### 3.2 字符串  
**3.2.1 Set**  
```golang
Set(dbsid, key string, value interface{}) error  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|
|value|值|
  
**3.2.2 Get**  
```golang
Get(dbsid, key string) (ret interface{}, err error) 
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|
|ret|返回值|

**3.2.2 删除**  
```golang
Del(dbsid string, key ...string) (int64, error)  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|
|ret|成功的个数|

**3.2.3 增加**  
```golang
Incr(dbsid, key string) (int64, error)  
IncrBy(dbsid, key string, delta int64) (int64, error)  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|
|delta|要增加的值|

**3.2.3 取值的长度**  
值的类型是字节流返回实际长度  
值的类型是字符串，返回字符串的长度。（utf8格式）
```golang
Strlen(dbsid, key string) (int64, error)  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
<!--
**3.2.4查询**  
Scan 根据参数查找符合条件的域名值
```golang
Scan(dasid string, begin, match string, count int, inclusive bool, tp byte) (keys []string, err error)
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|begin|键的起始值|  
|match|匹配条件| 
|count|返回的最大个数| 
|inclusive|包含边界值| 
|tp|类型| 

**3.2.5逆序查询**  
scan 根据参数查找符合条件的域名值
```golang
RevScan(dasid string, begin, match string, count int, inclusive bool, tp byte) (keys []string, err error)
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|begin|键的起始值|  
|match|匹配条件| 
|count|返回的最大个数| 
|inclusive|包含边界值| 
|tp|类型| 
 -->
### 3.3 哈希表  
**3.3.1 删除哈希表**  
删除哈希表 不存在的key将被忽略。  
```golang
Hmclear(dbsid string, key ...string) (ret int64, err error)  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|ret|成功的个数|

**3.3.2 删除域**  
删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略。  
```golang
Hdel(dbsid, key string, field ...string) (ret int64, err error)  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|field|域|  
|ret|成功的个数|

**3.3.3 哈希表大小**  
返回哈希表 key 中域的数量。
```golang
Hlen(dbsid, key string) (ret int64, err error)    
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|ret|域的数量|

**3.3.4 写**  
将哈希表 hash 中域 field 的值设置为 value 。
```golang
Hset(dbsid, key, field string, value interface{}) (ret int64, err error)  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|field|域| 
|value|值| 
|ret|域的数量|不存在返回1，存在返回0

**3.3.5 读**  
返回哈希表中给定域的值。
```golang
Hget(dbsid, key, field string) (ret interface{}, err error)    
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|field|域| 
|ret|值|


**3.3.6 批量读**  
返回哈希表 key中，一个或多个给定域的值。
```golang
Hmget(dbsid, key string, fields ...string) (ret []interface{}, err error)  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|field|域| 
|ret|值|


**3.3.7 批量写**  
同时将多个 field-value (域-值)对设置到哈希表 key 中。  
此命令会覆盖哈希表中已存在的域。
如果 key 不存在，一个空哈希表被创建并执行 Hmset 操作。
```golang
Hmset(dbsid, key string, values ...FVPair) error  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|field|域| 
|value|域-值| 

```golang
// FVPair is the pair of field and value.
type FVPair struct {
	Field []byte
	Value []byte
}
```

**3.3.8 取所有域和值**  
返回哈希表 key 中，所有的域和值。
```golang
Hgetall(dbsid, key string) (ret []FVPair, err error)    
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|ret|域-值|  

**3.3.9 取所有域**  
返回哈希表 key 中，所有的域。
```golang
Hkeys(dbsid, key string) (ret []string, err error)  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|ret|域| 

**3.3.10 增加值**  
为哈希表 key 中的域 field 的值加上增量 delta 。
```golang
HincrBy(dbsid, key, field string, delta int64) (ret int64, err error)  
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|delta|增加值|  
|ret|当前值| 

**3.3.11查询**  
Hscan 根据参数查找符合条件的域名值
```golang
Hscan(dbsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)   
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|beginfield|域的起始值|  
|match|匹配条件| 
|count|返回的最大个数| 
|inclusive|包含边界值| 
|ret|符合条件的域值| 

**3.3.11逆序查询**  
Hscan 根据参数查找符合条件的域名值
```golang
Hrevscan(dbsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)   
```
|参数|名称|说明|
|--|--|--|
|dbsid|会话id|MMOpen获取
|key|键|  
|beginfield|域的起始值|  
|match|匹配条件| 
|count|返回的最大个数| 
|inclusive|包含边界值| 
|ret|符合条件的域值| 

### 3.4 列表 
<!--
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
 -->
### 3.5 集合  
<!--
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
 -->
### 3.6 有序集  
<!--
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
 -->