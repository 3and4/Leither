弥媒
========  
弥媒是用于描述数据和应用  

<!--传统的中心化互联网，应用和数据是耦合在一起，由单一集团统一规则运行维护。  
在去中心化互联网中，

为了解决节点之间的数据协同问题,节点  
-->

### 一、弥媒文件
本文描述系统中的文件对象 

#### 1.1、文件对象  
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
  
#### 1.2、打开文件
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
MFOpenByPath(sid, fsid, path string, flag int) (string, error)
这种方式打开的是弥媒文件系统或操作系统文件系统中的文件  

**打开Mac文件**   
MFOpenMacFile(sid, mmfsid, fileid string) (string, error)
Mac文件必须挂接在某个弥媒之下。  

**打开临时文件**   
MFOpenTempFile(sid string) (string, error)  
操作完成之后转换为Mac文件    
MFTemp2MacFile func(sid, mid string) (string, error)  

**关闭文件**  
除临时文件外，其它文件操作完成后都需要关闭文件  
考虑到远程操作文件经常会有异常掉线情况。   
操作对象的会话id在超时情况下都会自动关闭文件  


#### 1.3、文件读写  
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
MFFind(sid, mmfsid, path string) (*FindResult, error)
<!--
MFOSFSMkDir(sid, fsid, path string, perm os.FileMode) error
MFOSFSRemoveAll(sid, fsid, path string) error
MFOSFSStat(sid, fsid, path string) (*FileInfo, error)
MFOSFSRename(sid, fsid, oldpath, newFullName, newLeft string) error
-->

**其它**  
MFTruncate(fsid string, size int64) error  
MFCopy(fsid, dst, src, srcVer string) error  

#### 1.4、对象引用
MFAddRef(sid, mid string, fileids ...string) (int, error)  
MFDelRef(sid, mid string, fileids ...string) (int, error)  
MFGetRef(sid, mid string) (ret map[string]int, err error)  




### 二、数据库


### 三、弥媒操作

### 四、权限
