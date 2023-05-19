MiMei
========  
MiMei is used to describe data and applications, and is the most important concept in Leither system, similar to File in Linux.  


## I. Fundimentals  
MiMei has two types of data: File and Database  
They share common API: create, open, priviledge and version management.  
<div id="MMCreate"></div>    

### 1.1 Create MiMei   
With necessary permissions, user can create MiMei on the current node.  
If *mark* is "{{auto}}", generate a new MiMei ID each time.  
Otherwise, the same parameters generate the same MiMei ID.
```golang
MMCreate(sid, appid, ext, mark string, tp byte, right uint64) (mid string, err error)
```
  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|appid|App id|The App that creates this MiMei. Null represents Leither
|ext|extension type|Distinguish MiMei type within an App
|mark|special identifier|create MiMei ID according to mark  
|tp|data type|1: File, 2: Database
|right|permissions|

|Return|Note|
|--|--|
|mid|ID of created MiMei

### 1.2 Open MiMei
Open an existing MiMei  
```golang
MMOpen(sid, mid, ver string) (msid string, err error)
```

|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mid|MiMei id|
|ver|MiMei version|

|Return|Note|
|--|--|
|msid|session id, for MiMei operation API

3 special versions：  
|Version|Note|
|--|--|
|cur|current version, writable version must be *cur*
|last|the latest version，Backup command turns *cur* to *last*
|release|stable version, officially released


### 1.3 Setup permissions
Setup MiMei permissions
```golang
MMSetRight(sid, mid, member string, right uint64) error
```  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mid|MiMei id|
|member|Object receiving permissions
|right|same format as in Linux

### 1.4 Backup MiMei
The backup of *cur* generates a new version *last*     
```golang
MMBackup(sid, mid, memo string) (ver string, err error)   
```  

|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mid|MiMei id|
|memo|memo|unused placeholder  
|return|title|note|
|--|--|--|
|ver|new version number|*last* points to this version

### 1.5 Publish MiMei  
Appoint a version as current version
```golang
MMRelease(sid, mid, ver string) (newVer string, err error)
```  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mid|MiMei id|
|ver|appointed version
 

|Return|Note|
|--|--|
|newVer|version number pointed by release

### 1.6 Restore MiMei
Set a version as current version  
```golang
MMRestore(sid, mid, ver string) error
```  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mid|MiMei id|
|ver|MiMei version|

### 1.7 Delete version清除版本  
Delete data of certain versoins  
```golang
MMDelVers(sid, mid string, vers ...string) (int64, error)
```  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mid|MiMei id|
|vers|versions to delete
 
### 1.8 Lookup information  
MiMei information can be queried by GetVar function.  
```golang
GetVar(sid, name string,args ...string) (interface{}, error)
```  
Definition of some contant in API package.  
```golang
//MiMei section
const (
	ApiVarMiMeiInfo     = "mminfo"     //MiMei information
	ApiVarMiMeiVersions = "mmversions" //MiMei version
)
```

**1.8.1 Lookup MiMei information**  
```golang
mminfo, err := api.GetVar(sid, "mminfo", mid) 
```  
Data structure of MiMei information.  
```golang
type MiMeiInfo struct {
	AppType string    //application type
	Author  string    //author
	Ext     string    //extension
	Mark    string    //mark
	Create  time.Time //time of creation
	Right   uint64    //permissions
}
``` 

### 1.9 Reference
**1.9.1 Add reference**  
```golang
MMAddRef(sid, mid, fileids ...string)(count int, err error)
```  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mid|MiMei id|
|fileids|IDs of objects to be referred to    

|Return|Note||
|--|--|--|  
|count|# of succeedly referred IDs|  
 
   
**1.9.2 Delete reference**  
```golang
MMDelRef(sid, mid string, fileids ...string) (count int, err error)  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mid|MiMei id|
|fileids|Reference IDs to be deleted    

|Return|Note||
|--|--|--|  
|count|# of succeedly deleted IDs|  
```  

**1.9.3 Get reference**  
```golang
MMGetRef(sid, mid string) (refs map[string]int, err error)  
```  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mid|MiMei id|

|Return|Note|Detail|
|--|--|--|  
|refs|# of reference|｛File_id->ref#｝  
   

## II. File system   
MiMei file system  
### 2.1 Open and close of MiMei file  
**2.1.1 Open file by path**  
```golang
MFOpenByPath(sid, mmfsid, path string, flag int) (fsid string, err error)  
```  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mmfsid|file system id|
|path|path of file|
|flag|options to open file|

|Return|Title|Note|
|--|--|--|  
|fsid|resource session id|returned by MFOpen  

  
**2.1.2 Open mac file**  
```golang
MFOpenMacFile(sid, mid, fileid string) (fsid string, err error)  
```  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|
|mid|file system id|
|fileid|mac file id|

|Return|Title|Note|
|--|--|--|  
|fsid|resource session id|returned by MFOpen   

**2.1.3 Close file**  
```golang
MFClose(fsid string) error  
```
|Return|Title|Note|
|--|--|--| 
|fsid|session id|Close file opened by MFOpen

**2.1.4 Open temporary file**  
```golang
MFOpenTempFile(sid string) (fsid string, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login|

|Return|Title|Note|
|--|--|--|  
|fsid|resource session id|returned by MFOpen  
  
**2.1.5 Convert temporary file to Mac file**  
```golang
MFTemp2MacFile(fsid, mid string) (macid string, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|Returned from MFOpenTempFile
|mid|file system id|

|Return|Title|Note|
|--|--|--|  
|macid|mac file id|  

### 2.2 Access file  
**2.2.1 Write data as object**  
```golang
MFSetObject(fsid string, obj interface{}) error  
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|returned by MFOpen
|obj|Object to be written into fsid|

**2.2.2 Get data as object**  
```golang
MFGetObject(fsid string) (obj interface{}, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|returned by MFOpen

|Return|Title|Note|
|--|--|--|  
|obj|returned object|  
  
The interface{} is mainly designed for dynamic languages like JavaScript. With static languages, such return values will lose their type information. It is necessary to use an alternative method to wrap the API for these languages.    

**2.2.3 Write data as byte array**  
```golang
MFSetData(fsid string, data []byte, start int64) (count int, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|returned by MFOpen
|data|byte array to be written into fsid|
|start|start position|positive value means an offset from the beginning, negative the end, and -1 indicates the end of the file.  

|Return|Title|Note|
|--|--|--|  
|ret|result returned|  
  

**2.2.4 Get data as byte array**  
```golang
MFGetData(fsid string, start int64, count int) (ret []byte, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|returned by MFOpen
|start|start position|positive value means an offset from the beginning, negative the end, and -1 indicates the end of the file. 
|count|# of bytes|

|Return|Title|Note|
|--|--|--|  
|ret|result returned|   
 
### 2.3 Check status  
**2.3.1 Get file size**  
```golang
MFGetSize(fsid string) (size int64, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|returned by MFOpen

|Return|Title|Note|
|--|--|--|  
|size|file size|  

**2.3.2 Get MiMei type of a file**  
```golang
MFGetMimeType(fsid string) (tp string, err error)    
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|returned by MFOpen

|Return|Title|Note|
|--|--|--|  
|type|MiMe type|  

**2.3.3 Get file information**  
```golang
MFStat(fsid string) (info *FileInfo, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|returned by MFOpen

|Return|Title|Note|
|--|--|--|  
|info|file information|  

**2.3.4 Check file existence**  
```golang
MFIsExist(fsid, fileid string) (bExist bool, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|returned by MFOpen
|fileid|file id|  

|Return|Title|Note|
|--|--|--|  
|bExist|file status|  


### 2.4 Directory operation  
**2.4.1 Get file information under a directory**  
```golang
MFReaddir(fsid string, count int) (infos[]*FileInfo, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|returned by MFOpen
|count|# of files to read|0 or less indicates all files should be read

|Return|Title|Note|
|--|--|--|  
|infos|information of files under the directory|  

### 2.5 Other file operations  
**2.5.1 Truncate file**  
```golang
MFTruncate(fsid string, size int64) error  
```
|Parameter|Title|Note|
|--|--|--|
|fsid|session id|returned by MFOpen
|size|file length|  |  


### 2.5 File system operation  
**2.5.1 Search by path**  
```golang
FSFind(sid, mmfsid, path string) (result *FindResult, err error)  
```  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login
|mmfsid|file system mid|  
|path|relative path|  

|Return|Title|Note|
|--|--|--|  
|result|search result|  

```golang
type FindResult struct {
	FSID     string //file system id  
	Left     string //unmatched path, empty when matched  
	FullName string //full path  
}
``` 

**2.5.2 Create directory**  
```golang
FSMkDir(sid, mmfsid, path string) error  
```  
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login
|mmfsid|file system mid|  
|path|relative path| 

|Return|Title|Note|
|--|--|--|  
|result|result|   

**2.5.3 Lookup file information**  
```golang
FSStat(sid, mmfsid, path string) (info *FileInfo, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login
|mmfsid|file system mid|  
|path|relative path|  

|Return|Title|Note|
|--|--|--|  
|info|file information|  

**2.5.4 Duplicate file**  
```golang
FSCopy(fsid, dst, src, srcVer string) error  
```
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login
|dst|target path|  
|src|source path|  
|srcVer|source version|  

**2.5.5 Rename file**  
```golang
FSRename(sid, mmfsid, oldpath, newFullName string) error
```
|Parameter|Title|Note|
|--|--|--|
|sid|session id|returned by Login
|mmfsid|file system mid|  
|oldpath|old file path|  
|newFullName|new file path|  


## III. Database  
Databse as MiMei  
### 3.1 Transaction  
**3.1.1 Start transaction**  
```golang
Begin(dbsid string, timeout int) error 
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|timeout|delay|unit in second, auto rollback when timeout  

**3.1.2 Commit transaction**  
```golang
Commit(dbsid string) error  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen


**3.1.3 Rollback transaction**  
```golang
Rollback(dbsid string) error 
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen

### 3.2 string  
**3.2.1 Set**  
```golang
Set(dbsid, key string, value interface{}) error  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|
|value|value|
  
**3.2.2 Get**  
```golang
Get(dbsid, key string) (ret interface{}, err error) 
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|
|ret|returned value|

**3.2.2 Delete**  
```golang
Del(dbsid string, key ...string) (int64, error)  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|
|ret|# of successfully deleted|

**3.2.3 Increase**  
```golang
Incr(dbsid, key string) (int64, error)  
IncrBy(dbsid, key string, delta int64) (int64, error)  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|
|delta|value to be increased|

**3.2.3 String length**  
For byte stream, return the actual length  
For string, return the string length (utf8)
```golang
Strlen(dbsid, key string) (int64, error)  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  

### 3.3 Hash table  
**3.3.1 Delete key from hash table**  
Delete key from hash table. Non-exist key is ignored.  
```golang
Hmclear(dbsid string, key ...string) (ret int64, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|   
|ret|# of successes|

**3.3.2 Delete field**  
Given key, delete fields from hash table.  
```golang
Hdel(dbsid, key string, field ...string) (ret int64, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|   
|field|fields|  
|ret|# of successes|

**3.3.3 Hash table size**  
Given key, return number of fields.  
```golang
Hlen(dbsid, key string) (ret int64, err error)    
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  
|ret|# of fields|

**3.3.4 Set**  
Given key, set value of a field.  
```golang
Hset(dbsid, key, field string, value interface{}) (ret int64, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  
|field|field| 
|value|value| 
|ret|# of fields|return 1 if non-exist, 0 if exist

**3.3.5 Get**  
Given key and field, return corresponding value.  
```golang
Hget(dbsid, key, field string) (ret interface{}, err error)    
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  
|field|field| 
|ret|value|

**3.3.6 Batch get**  
Given key and fields, return value set.  
```golang
Hmget(dbsid, key string, fields ...string) (ret []interface{}, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  
|field|field| 
|ret|value|

**3.3.7 Batch set**  
Given key, set multiple field-value pairs in hash table. Existing field will be overwritten.  
If key does not exist, an empty hash table is created to execute the command.  
```golang
Hmset(dbsid, key string, values ...FVPair) error  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  
|field|field|  
|value|value| 

```golang
// FVPair is the pair of field and value.
type FVPair struct {
	Field []byte
	Value []byte
}
```

**3.3.8 Get all field-value pairs**  
Given key, return all field-value pairs.  
```golang
Hgetall(dbsid, key string) (ret []FVPair, err error)    
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  
|ret|field-value pairs|  

**3.3.9 Get all fields**  
Given key, return all fields.  
```golang
Hkeys(dbsid, key string) (ret []string, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  
|ret|fields| 

**3.3.10 Increase field value**  
Given key, increase field by delta.  
```golang
HincrBy(dbsid, key, field string, delta int64) (ret int64, err error)  
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  
|delta|delta|  
|ret|current value| 

**3.3.11 Scan**  
Given key and other conditionss, search for field-value pairs.  
```golang
Hscan(dbsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)   
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  
|beginfield|beginning field value|  
|match|search conditions| 
|count|max # of returned pairs| 
|inclusive|boolean to indicate if inclusive on border| 
|ret|result| 

**3.3.11 Reverse scan**  
Scan in reversed order
```golang
Hrevscan(dbsid, key, beginfield, match string, count int, inclusive bool) (ret []FVPair, err error)   
```
|Parameter|Title|Note|
|--|--|--|
|dbsid|session id|returned by MMOpen
|key|key|  
|beginfield|beginning field value|  
|match|search conditions| 
|count|max # of returned pairs| 
|inclusive|boolean to indicate if inclusive on border| 
|ret|result| 

### 3.4 List 

### 3.5 Set  

### 3.6 Ordered set  
