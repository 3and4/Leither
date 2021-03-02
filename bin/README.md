程序解压复制到目录中
直接运行便可
确保有相关权限
最新版本V0.10.63

Github不适合存放大的二进制文件。
新的程序放到了下面网址
<http://leither.cn/dl/>

**Leither.darwin64**  
MacOS amd64芯片，通常是苹果电脑或笔记本 

**Leither.amd**  
Linux amd64芯片，通常是服务器  

**Leither.arm6**  
linux arm6芯片，通常是nas和树莓派  

**Leither.arm7**  
linux arm7芯片，通常是nas和树莓派  


**扩展名.exe**  
windows x86 64位芯片 适合常见的pc和笔记本  

**Leitherdev.exe**  
windows x86 64位芯片 开发者版本
和普通版本的差别是有控制台，可以命令行进行交互



**近期的一些版本变动**
20210302
最新版本号："V0.10.63"
	整理mkdir代码，	去除dav.ProsessRequest
	文件系统的错误信息标准化，兼容http和资源管理器
	修订了Golang webdav 系统包的bug
	FSName函数去除Left参数

20210222
最新版本号："V0.10.62"
	整理弥媒Api：
		引用部分Api
			从MFileStub=>MiMeiStub
		文件系统Api
			OSFileStub=>MFSStub
			MFFind=>FSFind
			MFCopy=>FSCopyqw
			MFOSFSMkDir=>FSMkDir 同时去除perm参数
			MFOSFSRemoveAll=>FSRemoveAll
			MFOSFSStat=>FSStat
			MFOSFSRename=>FSRename
	去除GetVar中的：
		"mmfsid"
		ApiVarMMRoot："mmroot"
	重构
		引用函数，出错中止，返回成功的个数，出错id和出错信息

20210208  
最新版本号："V0.10.61"  
	整理查询弥媒信息的相关代码
  
20210202  
最新版本号："V0.10.60"  
	MFReaddir中如果有符号链接，返回指向的文件信息  
  
20210201  
最新版本号："V0.10.59"  
	变动：  
		把设置工作目录的功能限定在run指令集  
		去降lmac中的多余日志  
		获取ipv4的代码合并  
		增加节点管理员删除节点上的他人域名功能  
	bug:  
		ip检查时导致失败的bug  
		restore时Cache清除出错  
    
20210127  
最新版本号："V0.10.58"  
	变动：   
		deploy命令行变为lapp  
		SignPPT重构，参数变动  
		MMVersion取消，用MMBackup,MMRestore替代其功能  
		lapp version命令行取消  
		工作目录的运行机制变动  
		签发通行证的机制调整  
		命令行的框架调整  
	增加  
		run命令行，可以设置工作目录  
		增加lapp release功能  
		删除版本时增加去重  
	Bug:  
		upload时，last版本中的应用丢失文件  
		数据序有序集Zrange数据不全  
		lpki shell命令中恢复数据库对象查询  
  
20210119  
最新版本号："V0.10.57"  
	SignPPT重构  
