
Github不适合存放大的二进制文件。
最新的程序放到了下面网址  
<http://leither.cn/dl/>


程序解压复制到目录中
直接运行便可
确保有相关权限
最新版本V0.11.29


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
20220731  
最新版本号："V0.11.29"  
161272 nouse3735  
命令增加  
	删除弥媒版本 	Leither mimei delvers  
	取macid		Leither mimei macid  
	显示弥媒		Leither mimei show
	同步弥媒到本地	Leither mimei sync  
	弥媒发布		Leither mimei publish  
	成为数据节点	Leither mimei provide  
	添加密钥		Leither lpki addkey  
	删除密钥		Leither lpki delkey  
	pin删除	 	Leither ipfs pin rm  
命令优化  
	所有命令缺省值都进行了优化  
Api  
	生成一个临时id, 供外部节点操作本地文件  
	MFPath2TempID(sid, localpath string) (string, error)  
	本地文件上传为mac格式文件  
	MFLocal2MacFile func(sid, mid, ps string) (string, error)  
  
bug  
	生成证书的缺省值问题  



20220608  
最新版本号："V0.11.21"  
前接  
	系统的业务功能基本具备，大量细节需要优化。  
本次    
	优化命令行，优化跨节点操作，跨目录操作支持    
	命令行  
		生成一个会话id  
			Leither lpki gensid  
		设置弥媒的内容  
			Leither mimei setid <cid>  
		支持通过指定节点id访问节点,不指定情况下，缺省当前节点和密钥进行命令行操作  
			-n 节点地址或id  
		支持了跨节点上传数据  
			Leither ipfs add testvideo -n Ngacq50-IRX_DfcndFwZ0c9S0Nh  
	其它  
		取消了旧的网关机制，后续直接从网络上查找节点  
		配置中的网关作为网络的接入点处理  
		节点外用户登录支持消息  
		工作目录调整  

20220519  
最新版本号："V0.11.18"  
前接  
	上次发布不支持文件系统，上传大文件时间过长，容易卡死  
	发布在网络上的内容会被定期清理，应用经常白屏  

  
本次  
	弥媒支持dag文件系统,重构文件系统(dag部分),支持类unix命令行及相应的api  
	命令集和相应的api  
		增加Files命令集和对应的api,包括ls,cp,mkdir,mv,rm,stat,flush  
		增加pin命令集和对应的api, 包括add  
		增加mimei命令集和对应的api,包括ls  
	改进  
		publish支持时效（EOL）  
		publish支持自动发布（之前1天自动失效，现在和核酸检查续期差不多）  
		应用支持文件系统展示  
		ipfs add支持进度显示  
	Bug  
		修改了地址报备出错造成的页面白屏现象  
		修改了发布应用过程中的bug  
		webpack打包时的全局变量出错bug  
  
20220425    
最新版本号："V0.11.15"    
	完全去中心化的域名解析。  
		域名解析去中心化  
			网络内任意节点都可以解析弥媒应用和内容，以兼容传统的互联网生态  
		弥媒信息去中心化  
			信息包括作者关联应用权限版本等，对应publish机制  
		弥媒内容支撑节点去中心化  
			内容和展示应用都有支撑节点提供，对应provide机制  
			支撑节点可以动态加入退出，支持负载均衡和容错  
		节点之间路由查询去中心化  
			由dht机制实现完成  
	数据的上传和备份机制  
		进行了异步处理，命令行加了进度条提示  
	增强了命令行和后台的交互  
		通过消息通讯机制把内部执行过程同步到前台  

20220407  
最新版本号："V0.11.14"  

近期进展  
	弥媒数据发布到网络  
	弥媒信息发布到网络（还少版本部分）  
	弥媒支持节点信息发布到网络  

	webapi   
		ipfs域名解析展示本地的弥媒   

	增加api  
		IpfsAdd(sid, ps string) (string, error)  
		IpfsProvide(sid string, rec bool, keys ...string) error  
		IpfsFindProvs(sid, cid string, n int) ([]AddrInfo, error)  
		IpfsPublish(sid string, mids ...string) error  
		MiMeiPublish(sid string, mids ...string) error  
		MiMeiProvide(sid, dhts string, rec bool, keys ...string) error  
		MiMeiFindProvs(sid, dhts, strmid string, count int) ([]AddrInfo, error)  

		HExists(key []byte, field []byte) (int64, error)  
		HasMiMei(mid string) (bool, error)  

	增加命令  
		Leither dht provide  

		Leitehr mimei create  
		Leitehr mimei setdata <mid> <filename>  
		Leitehr mimei publish <mids>  
		Leitehr mimei findprovs <mid>  


20220303  
最新版本号："V0.11.13"  
近期进展  
		阅读了ipfs上层数据存储相关的代码  
系统中增加了Leither ipfs add命令行，功能是添加数据到本地的ipfs节点  
增加了/ipfs/{cid}的web服务，用于通过url查看ipfs网络中的数据内容  
当前目标  
		给ipfs资源增加leither的域名访问功能有。  
可以在微信里传播，支持负载均衡和容错。  
	当前工作  
		融合ipfs节点间的数据流转功能  
		学习bitswap协议相应知识  


20220207
最新版本号："V0.11.12"  
	当前目标是把弥媒的数据和应用发布在ipfs网络上  
	目前dhts增加了ipfs支持。  
	ipfsnode跑通了测试案例  
	后续ipfsnode增加到pnet中  
	重构了nodeid部分  
	修正了peerstore中的bug(之前未完成的record部分)  
	修订了memcache的map冲突bug  

20220112  
最新版本号："V0.11.11"  
去除引入的libp2pcore  

20220112  
最新版本号："V0.11.10"  
长短id兼容  
	获取节点id  
	节点预处理  
	dht中内核使用短id  
		self  
		selfKey  

20211221  
最新版本号："V0.11.09"  
	节点管理全部使用PeerStore  

20211221  
最新版本号："V0.11.08"  
调整  
	使用上下文控制关闭任务  
		TaskSchedule  
		SessionManager  
		MiMeiCenter  
	根据密钥生成keypair  
	域名和相关getvar使用peerstore中的地址  
	地址检查流程使用peersotre机制  
废弃  
	原来的地址广播   
bug  
	api句柄超时问题  
	webfunc getipfs 不支持pid  
	地址检查ipv6漏查  

20211219  
最新版本号："V0.11.07"  
手机版  
	消除配置  
	应用变为后台服务  
	优化休眠问题  
	支持sd卡  
	支持安装应用  
	解决链接数过多死机问题  
	支持ipv6  
	支持mac身份识别  
	重构鉴权流程  
pc版  
	我的节点  
	节点网络  
	我的应用  
操作网内节点  
	过滤掉ipfs网络的节点    

20211109  
最新版本号："V0.11.06"  
	Bug:  
		ips2Multiaddr地址解析出问题  

20211107  
最新版本号："V0.11.05"  
	手机版增加了框架。  
		我的节点，朋友节点，应用服务，消息交互  
	我的节点启动关闭服务  
	朋友节点增加了list的部分代码  
  
20211103  
最新版本号："V0.11.04"  
	功能  
		新添加了手机版，并且稳定运行  
	修订了一些Bug  
		密钥id的生成bug  

20211021   
最新版本号："V0.11.03"  
	增加api  
		DHTStub  
			DhtFindPeer func(sid, pid string) (*AddrInfo, error)  
		SwarmStub  
			SwarmAddrs      func(sid string) (map[string][]string, error)  
			SwarmLocal      func(sid string) ([]string, error)  
			SwarmListen     func(sid string) ([]string, error)  
			SwarmPeers      func(sid string) ([]string, error)  
			SwarmFilters    func(sid string) ([]string, error)  
			SwarmConnect    func(sid, addr string) error  
			SwarmDisconnect func(sid, addr string) error  
			FiltersAdd      func(sid string, cidrs []string) error  
			FiltersRm       func(sid string, cidrs []string) error  
	功能
		防止ddos攻击选成的死机
			Stream打开过多就关闭链接
		节点黑名单
			可以配置和命令行两种方式操作
		增加10个命令
			Leither dht findpeer
			Leither swarm addrs
			Leither swarm addrs local
			Leither swarm addrs listen
			Leither swarm peers
			Leither swarm filters
			Leither swarm filters add
			Leither swarm filters rm
			Leither swarm connect
			Leither swarm disconnect
20211013  
最新版本号："V0.11.02"  
	变动：  
		整理日志中的目标当下情况  
		支持ip过滤节点  
		增加了ipfs stream式stub  
		节点之间互联通过流式stub完成  
		兼容了ipfs秘钥接口  
		返向rpc功能屏蔽  
		入口增加跨域支持  
		Node.Addr调用处设置Panic，看后续问题调整  
	Bug:  
		地址更新出错  
		新版chrome和Edge不支持局域网跨域  

20211010  
最新版本号："V0.11.01"  
	bug:  
		链接不断增加，超出限制。 增加了链接管理，问题解决  
		经常有攻击。增加了门控，问题还在，在持续跟踪  
	优化：  
		解决了id冲突问题。  
	功能  
		通过ipfs寻找失联代码  
	netpub包移入pnet  
	pub包go mod init  

20210930  
最新版本号："V0.11.00"  
	这是发布前最后一个子版号，后续版本号从13升为1  
	RedisTxer=>SystemDB  
	增加公钥支持，这是为了兼容ipfs和leither之间的id兼容  
		GetVar hostpk  
		节点增加公钥  
	调整最小版本，如果对方版本过低直接logout  
	getvar act恢复 ActWarn，SupportTunnel  
	接口调整  
	返向Rpc从fram=>pnet  
	配置文件中增加ipfs引导入口  
	WaitGroup, Context放入BUS  
	修复pnet conn不兼容ipfs的低层接口问题  
	服务没有启动的时候地址更新不工作（可能有问题）  
	basichost代码移入leither  
	重构host的地址部分。ipfs的host增加原来的leither地址功能  
	补齐代码AddNodeByWebApiHandler  
	增加findnode go程，功能还没有完全实现  

20210922  
最新版本号："V0.10.70"  
	把lnet包中的部分代码变成公共代码  
	pnet包增加api支持  
	pnet已替换lnet  
	lnet从项目中移除  

20210831  
最新版本号："V0.10.65"  
	清除不用的定义和代码：  
		FDOSFile File_Type_OSFile  
		File_Type_OSFS File_Type_OSFS  
	增加pnet包  
		后续替换lnet包  
		增加了libp2p路由功能  
		lnet功能尚未加入本包。  
		ipfs和leither密钥和id兼容  


20210316
最新版本号："V0.10.64"
	lapp命令集
		uninstall=>delvers
	mmfs重构前备份递交

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
