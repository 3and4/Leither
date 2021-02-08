弥媒
========  
弥媒是用于描述数据和应用，是系统中最重要的概念，类似Linux中的文件。      


## 一、基础功能  
弥媒有两种数据类型：文件和数据库。  
它们都有共同的api：创建、打开、权限、版本管理等。  
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
func (wa *WebApi) GetVar(sid, name string,args ...string) (interface{}, error)
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

**1.8.2查询弥媒版本信息**  
```golang
vers, err := api.GetVar(sid, "mmversions", mid) 
```  
返回的信息是版本列表,类型是字符串数组  
包含特殊版本last,release,不包括cur版本

<!--
## 二、弥媒文件  
弥媒的文件系统  
MFOpenByPath(sid, fsid, path string, flag int) (string, error)  
MFOpenMacFile(sid, mmfsid, fileid string) (string, error)  
MFOpenTempFile(sid string) (string, error)  
MFClose(fsid string) error  
MFFind(sid, mmfsid, path string) (*FindResult, error)  
MFTruncate(fsid string, size int64) error  
MFSetObject(fsid string, obj interface{}) error  
MFGetObject(fsid string) (interface{}, error)  
MFSetData(fsid string, data []byte, start int64) (count int, err error)  
MFGetData(fsid string, start int64, count int) ([]byte, error)  
MFCopy(fsid, dst, src, srcVer string) error  
MFGetSize(fsid string) (int64, error)  
MFStat(fsid string) (*FileInfo, error)  
MFIsExist(fsid, fileid string) (bool, error)  
MFReaddir(fsid string, count int) ([]*FileInfo, error)  
MFAddRef(sid, mid string, fileids ...string) (int, error)  
MFDelRef(sid, mid string, fileids ...string) (int, error)  
MFGetRef(sid, mid string) (ret map[string]int, err error)  
MFTemp2MacFile(sid, mid string) (string, error)  
MFGetMimeType(fsid string) (string, error)    

## 三、系统文件  
操作系统的文件系统  
MFOSFSMkDir(sid, fsid, path string, perm os.FileMode) error  
MFOSFSRemoveAll(sid, fsid, path string) error  
MFOSFSStat(sid, fsid, path string) (*FileInfo, error)  
MFOSFSRename(sid, fsid, oldpath, newFullName, newLeft string) error


## 四、数据库  
弥媒中的数据库  
Begin(dbsid string, timeout int) error  
Commit(dbsid string) error  
Rollback(dbsid string) error 
Set(dbsid, key string, value interface{}) error  
Get(dbsid, key string) (interface{}, error)  
Del(dbsid string, key ...string) (int64, error)  
Incr(dbsid, key string) (int64, error)  
IncrBy(dbsid, key string, delta int64) (int64, error)  
Strlen(dbsid, key string) (int64, error)  
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