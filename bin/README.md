
Github不适合存放大的二进制文件。
最新的程序放到了下面网址  
<http://vzhan.cn/mm/Fc1BRTFafOGzq5P8KmkVJqwS2v2/>


程序解压复制到目录中  
直接运行便可  
确保有相关权限  
最新版本V0.11.39  


**Leither.darwin64**  
MacOS amd64芯片，通常是苹果电脑或笔记本 

**Leither.amd**  
Linux amd64芯片，通常是服务器  

**Leither.arm6**  
linux arm6芯片，通常是nas和树莓派  

**Leither.arm7**  
linux arm7芯片，通常是nas和树莓派  


**Leither.windows.amd64**  
windows x86 64位芯片 适合常见的pc和笔记本  

**近期的一些版本变动**  

20240219
	当前版本号："V0.16.16" 递交之前的变化
	bug：
		wasm引擎标准输出异常
		正常关闭时，显示listener错误
		后台同步时，发用户消息超时错

20240214
	当前版本号："V0.16.13"
		js引擎初步通过

20240214
	当前版本号："V0.16.12"
		ipfs查询的时候标识错误，造成查询不存在。Direct => any

20240210
	当前版本号："V0.16.11"
	Bug:
		命令行的参数中如果有-开头，会误触发-flag流程
		Leither stop中没找到密钥会有异常错误

20240207
	当前版本号："V0.16.09"
	gc时重新provide时没有使用最新的时间
	开始加入js引擎代码大小变化如下：
				linux.amd64		linux.arm7		windows.amd64
swarm重构		16787252		14552108		17310208
加入js引擎		17371328		15053912		17902080


20240206
	当前版本号："V0.16.07"
	重新proivde失败，造成provide信息丢失

20240201
	当前版本号："V0.16.04"
	restore之后，最后序号处理错误，造成无法再备份

20240201
	当前版本号："V0.16.01"
	pinadd之后，加入了检查代码。

20240131
	递交
	日志同步时，刷新信息出错
	同步ipfs文件时，引用错误添加
	CheckIntegrity
	时区问题
	备份时删除引用
	退出时异常
	调整参数源节点，指定节点升级 应用
	增加update状态
	升级时出错恢复

20240126
当前版本号："V0.15.99"
	恢复restore函数

20240117
当前版本号："V0.15.94"
	AcceleratedDHTClient: true
	自动升级时过多的消息超时。
	下载更新时，程序异常关闭
	同步和下载时当前版本引用异常增加

20240111
当前版本号："V0.15.73"
	弥媒同步时，同步应用时空应用bug

	20240111
当前版本号："V0.15.72"
	弥媒同步时，同步应用

20240101
当前版本号："V0.15.72"
	检查数据库的一致性CheckIntegrity
	重构Travel函数
	增加AddRefs函数
	生成系统数据库mid
	增加选项，备份后清除引用

20231224
当前版本号："V0.15.66"
	引用出错造成的资源误删除
	程序关闭的时候最多等待1分钟。避免卡死情况
	命令 GetRight SetRight
	开发应用mExplorer
		支持md类型

20231212
191660	2120
当前版本号："V0.15.50"
	清理不用的api
		UserGroupStub //用户组
		DataRightStub //数据权限
		FileSystemStub
		//文件系统，描述磁盘上一个真实的文件
	移除旧的文件系统
	新的应用入口
	新的命令
		swarm id
	检查引用数
	自动升级bug
		更新当前版本
		重复下载当前版本
	重构命令行，认证参数部分
	增加ctrl-c
		SIGInterrupt
	tcp close优化
	日志删除bug
	引用出错造成应用发布失败
20231201
当前版本号："V0.15.48"
	域名解析，确保指定的节点在列表内。
	支持没有备案的域名

20231130
当前版本号："V0.15.47"
	重构：
		应用的入口部分进行了重构。去除了旧的代码实现。和Runscript等其它处的应用使用相同代码
	Bug:设置系统变量时经常有重入冲突
	解决：写操作时加锁

20231127
当前版本号："V0.15.45"
	重构wasm输出机制
	runscript支持运行弥媒本地应用
	Wasm挂接到了web服务上
	确本文本

20231120
当前版本号："V0.15.40"
	加入gc流程，清理过期的应用备份文件
		自动升级时会备份之前的可执行程序

20231119
当前版本号："V0.15.39"
	操作系统重启后，pid失效
		程序启动时，会记录当前程序的pid,用于后续操作。
		系统重启后，这个pid会失效，造成异常
	wasm leither api跑通
		重构了runner机制
		之前的代码重构优化
20231114
当前版本号："V0.15.37"
	wasm代码初步完成
		正式挂接前递交
	程序关闭和重启
		程序启动时记录pid,根据pid了解当前进程情况
	备份时取摘要出错
	查找提供者bug

20231023
当前版本号："V0.15.25.20"
	ipfs中ds用的数据库和弥媒数据库合并
	重新写了一下发布角本
	cpu指令对齐造成的不断重启
	hprose bug造成的同步数据出错。换hprose代价大，通过新增加一个api解决

程序大小变化
			linux.amd64		linux.arm7		windows.amd64
改动前		13300748		11350076		13660160
ipfs升级		16037716		13924716		16554496
数据库合并	15887100		13808592		16400896


20231013
当前版本号："V0.15.14.20"
	消息同步bug,特殊消息处理出错
	整理代码

20231011
当前版本号："V0.15.12"
	编译已通过，备份

20231002
当前版本号："V0.15.11"
	ipfs升级改动前备份
	变动尽可能放入iipfs包
	耦合的包进行解耦

20230922
当前版本号："V0.15.10"
	取消了旧的节点添加节点
	libp2p相关内容移入nodeid包，为升级网络代码作准备
	修订了第一次自动更新要多等 一天的bug
	取消了外网地址存取接口

20230831
当前版本号："V0.15.02"
	优化自动升级，支持程序外的其它配置文件
	重构优化配置文件中引导节点部分
	图标文件找入应用


20230814
当前版本号："V0.14.90"
	完成了自动升级
	配置
		Autoupdate：缺省是4点种检查重新，manual表示手动，不自动更新。hh:mm表示每天hh:mm更新
	命令行
		Leither update

20230801
当前版本号："V0.14.70"
	完成了应用的发布流程
		获取版本=》编译=》复制到弥媒=》应用发布流程
		自动升级的工作刚搭好框架
	重构命令:
		mimei 命令集若干命令
		ipfs  命令集若干命令
		files 命令集若干命令

20230701
当前版本号："V0.14.65"
	Bug:
		查找地址为空

20230623
当前版本号："V0.14.50"
	Api
		增加了mimei unpublish <mid>
		增加了mimei unprivode <mid>
	bug
		mimei get 卡死。
		mimei findprovs 超时
	调整
		lpki runscript 命令增加-k keyfile
		lpki requestservice 命令行增加 -k keyfile

20230524
当前版本号："V0.13.99"
	重构
		网络信息的广播机制(provide部分)
		网络操作时的去重机制
	增加
		ipfs的gc功能
		lpki id命令支持对输入文件取macid
		mimei ref get增加缺省时的当前版本获取
		同步流程加入了异常保护。出错时报错，不异常中止
		增加备份选项：备份时删除最后版本，缺省是True，格式 "rmlastver=false"

	调整
		移除了ipfs网络，保留了wan lan两个网络
		移除usergroup功能
		ipfs增加超时，3分钟没有收到包就结束
		值变动后，原有的叶摘要都删除，不再disable
		ipfs findprovs  缺省返回长id
		mimei findprovs 缺省返回短id
		引用查询ref get格式优化
	bug
		异步关闭文件出错
		同步时消息会卡住
		数据库回滚出错
		同步时，同步过多的引用文件
		ipfs同步时卡死
		删除版本时造成叶摘要错误，表现为备份出错
		去引用时数据库类型的弥媒报错

20230421
当前版本号："V0.13.45"
	重构
		网络信息的广播机制（部分替换）
	bug:
		macfile同步问题
		部分引用相关的问题
		去除日志中的二进制信息
		某种情况下，子库序号备份出错，造成数据库recover时异常
		子库序号解析出错，造成节点摘要被异常删除


20230412
当前版本号："V0.13.25"
	调整异步机制
		重构了异步的api和命令行
	重构了节点间数据同步的代码

	增加
		Leither ipfs pin ls
	重构
		ipfs pin命令集的消息机制及相关的api

20230324
当前版本号："V0.12.92"
	功能
		同步被引用的弥媒
	优化
		备份时，发布的版本不能删除
		发布时删除不用的旧发布版本

20230313
当前版本号："V0.12.84"
	180010	11340	240		50	//删除版本时保留删除信息,同步端代码
	异步函数变成同步方式
	备份同步流程
		备份后清除旧的last
		设置最后有效版本序号，指定序号之后才能同步差额变动数据
		同步取摘要失败之后，同步全部数据

20230222
当前版本号："V0.12.68"
	调整异步机制
	相关的命令
		Leither mimei sync
		Leither mimei sum
		Leither mimei setdata
		Leither mimei backup
	相关的api
		MMBackup
		MMSum
		MiMeiSync
		IpfsAdd
		MFTemp2Ipfs


20230213
当前版本号："V0.12.63"
	bug:
		Leither ipfs add命令权限

20230212
当前版本号："V0.12.58"
	重构命令 Leither ipfs add

20230211
当前版本号："V0.12.57"
	删除版本时，引用没有清理干净

20230208
当前版本号："V0.12.56"
	弥媒的缺省文件格式转为ipfs格式

20230203
当前版本号："V0.12.54"
	增加守护模式
		Leither run -d

20230203
当前版本号："V0.12.53"
	bug：
		最后的操作是删除操作时，备份生成的摘要有错

20230130
当前版本号："V0.12.43"
	bug：
		事务的迭代器未去重
		根据name取数据库索引出错
	其它若干备份和同步过程中的问题

20230113
当前版本号："V0.12.33"
	备份加入检查摘要操作
	用两种不同方法取摘要，如果失败直接报错

20230109
最新版本号："V0.12.32"
	bug:
		取摘要时，有大量删除信息时，特殊情况下处理有误

20230103
最新版本号："V0.12.31"
	bug:
		清版本的时候，没有清对应版本的摘要数据

20221231
最新版本号："V0.12.27"
	增加重取数据摘要模块，查摘要不等情况


20221227
最新版本号："V0.12.23"
	bug:
		消息过多时发生死锁问题

20221222  
最新版本号："V0.12.20"  
	校验过严，没考虑数据库情况,放松校验  
	数据库弥媒缺省了引用资源的同步    

20221214
最新版本号："V0.12.08"  
	bug:  
		同步的时候，提前写把弥媒版本信息写入数据库，造成复制文件出错时造成信息异常。  

20221213  
最新版本号："V0.12.07"  
	如果应用里的代码缺失，错误返回到最外层  

20221212  
最新版本号："V0.12.06"  
	增加mimei sum命令。  
		按key排序，对数据库每个位置取摘要，方便内容比对  
	消息通过关闭从超出数量关闭改为超时关闭  

20221208  
最新版本号："V0.11.99"  
	Bug:  
		同步任务一直在循环执行  
		原因是之前有序集中的垃圾数据造成任务删除不干净  

20221204  
最新版本号："V0.11.97"  
	事务在递交的时候，删除信息丢失。造成所有相关函数删除失败  

20221204  
最新版本号："V0.11.83"  
	bug:  
		zrem删除之后，再次zadd失败.调整空值判断解决  

20221203  
最新版本号："V0.11.82"  
	数据库加锁  
	事务加锁。有两种，一种是整个周期加锁，一个是递交时加锁   
	恢复数据库restore功能  
		提取指定版本到当前  
	bug  
		ipfs文件系统web方式浏览错误显示  

20221130  
最新版本号："V0.11.80"  
	数据库加锁  
	事务加锁。有两种，一种是整个周期加锁，一个是递交时加锁  
bug:  
	事务的变动检查，之前是所有数据库变动都返回错，现在是检查当前弥媒的数据库  
	删除信息不再记录系统库，并忽略restore时的删除信息  

20221129  
最新版本号："V0.11.78"  
	bug:  
		事务递交时，有情况下没有解锁，造成数据库卡死  

20221125  
最新版本号："V0.11.77"  
	使用模板功能的时候，如果模板异常报错  
	处理同步时的异步任务数据造成同步卡死的问题  
	以下两个命令中多个mid参数调整为单mid  
	mimei publish; mimei provide  

20221124  
最新版本号："V0.11.74"  
	findprovs时触发旧的数据节点更新数据流程跑通  
	bug:  
		同步一直没成功，数据gc清理后，重复执行时有nil错  
		网络用数据库gc时会卡死，相关操作异步处理  
  
  
20221123  
最新版本号："V0.11.68"  
	findprovs时触发旧的数据节点更新数据  

20221122  
最新版本号："V0.11.67"  
	重构publish  
		同步发送信息到provider  
	bug  
		ipfs里peer.ID和golang的gob包有冲突，造成provider读取的时候数据丢失  

20221121  
最新版本号："V0.11.66"  
	publish时同步更新provider上的信息，并且触发数据同步  

	bug:  
		命令行中的参数首字母不能为-，  
		弥媒续期过程卡死bug  

20221120  
最新版本号："V0.11.62"  
	bug：弥媒续期时，记录到provide数据库，没有记录到dht网络数据库  
  
20221119  
最新版本号："V0.11.61"  
	重构弥媒的续期函数。两处变动：  
	1、依据发布时的版本进行续期  
	2、不再单独启动go程，依托在gc中  
  
20221117  
最新版本号："V0.11.59"  
	增加了数据库检索命令  
		Leither mimei grep <mid> <ver> -b <begin> -m <match>  
	缺省显示指定数据库的所有内容，可以通过-b指定起始值,通过-m指定筛选条件  
	后续可以根据需求加上其它条件。包括不限于反向查找，指定数据类型，数量，结束key等  
	修订了一个迭代器读取的bug  
	修订了同步时漏掉一个版本的bug  
  
20221116  
最新版本号："V0.11.55"  
	重构事务代码，支持删除信息  
	publish去除eol参数  
	加强节点用户的权限  
	增加同步过程中的提示信息  
	解决之前备份过程中丢失的删除信息  

最新版本号："V0.11.53"  
	调整MiMeiPublish参数，去除EOL参数  
	准备应用案例dht.lua  

20221109
最新版本号："V0.11.52"  
Bug:  
	修改了删除节点密钥时造成两个密钥的情况（网络上会出现空地址节点）  
	备份和同步时没有正确处理删除信息  

20221105  
最新版本号："V0.11.46"  
	修改了兼容早期域名时的一个bug  
	provide版本解析错  
	删除数据库当前版本  

20221103  
最新版本号："V0.11.45"  
	修改了模板页在无弥媒时的提示瑕疵  
	修改了两个模板地址的bug,一个是空的时候返回值错，一个是多节点的时候，点错了  

20221102
最新版本号："V0.11.44"  
重构备份数据库的迭代器  
重构模板页  
	节点地址增加距离概念，优先均衡到最近节点  
	识别不同类型地址  
		当前节点公网地址>其它节点的公网地址（同运营商同地区优先）>本节点的内网地址  
	识别浏览器的安全设置  
	没有外网地下，同时cors限制的情况下  
	提示用户手工跳转，或打开浏览器选项  
bug  
	消息为空时异常错  
	查找地址为空。  
	应用错误无法显示在均衡的模板页上  
	数据备份状态异常。stack的每层status复用的时候没有初始化好  
	hmget last版本数据无法读取  
其它  
	若干测试案例  
  
20221027  
最新版本号："V0.11.43"  
	备份时卡死和泄漏bug  
	lua中table的映射类型json不支持bug  

20221026
最新版本号："V0.11.42"  
html扩展名兼容bug  
应用发布不能立刻生效bug  


20221025  
最新版本号："V0.11.41"  
lua角本中table类型串行化类型错  
lua角本中table的key首字母被转换成了大写  

20221024  
最新版本号："V0.11.40"  
重构域名设置，为均衡作准备  
重构应用框架，为后续的测试案例作准备  
重构角本应用中的log包  
修改了包名提取的bug  
增加了可以外部运行的测试案例  
都放在了TestCase目录  

20221019  
最新版本号："V0.11.39"  
	重构旧的域名setdomain  
	原来的域名只支持解析到固定节点    
	新域名增加支持网络上的弥媒和应用  

20221018  
最新版本号："V0.11.38"  
	优化findprovs  
	修订provide bug  
	优化了架构上的总线，避免循环引用  

20221017  
最新版本号："V0.11.37"  
	支持guest用户  

20221016  
最新版本号："V0.11.36"  
	整理了弥媒引用相关的代码  
		修正了之前备份时的引用bug  
		增加命令 getref delref addref  
		网络发布，查询，同步支持资源引用  
	增加了命令 id 显示当前节点用户id  
	删除版本的命令delvers改为delver  
	启动的时候显示启动时间  
	调整了cache机制，支持访问增加cahche有效时间  

20221004  
最新版本号："V0.11.32"  
数据库支持去中心化网络，主要是四个操作（命令或api）  
数据所有者：  
	备份（backup）操作可以对数据库作一个镜像  
  		发布(publish)操作可以把这个版本的数据库发布到网络  
	数据使用者  
  		同步(sync)操作可以从网络或指定节点同步数据库内容到本地  
	数据支撑者  
  		支撑(provide)操作表示本节点愿意对这个数据库提供支撑  
  		数据变动的时候，支撑节点的数据库会实时变动。  
	数据库的一致性通过版本（序号）进行保证  
其它  
重构了版本删除功能  
迭代器支持子库序号  
重构了shell命令，增加了数据库操作，打开数据库，get,put  
数据库事务的begin函数参数定义改变,和leveldb保持了一致  
mimei create 支持数据库类型


20220831
最新版本号："V0.11.30"  
	重构了数据库的备份，单元测试通过，还未挂接到弥媒备份  
	重构了shell命令，增加了数据库操作，打开数据库，get,put  
	数据库事务的begin函数参数定义改变,  
		Begin(bWrite bool) error  
		和leveldb保持了一致  
	mimei create 支持数据库类型  
	bug:被provide go程卡死，退出异常  

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
