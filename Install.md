
Github不适合存放大的二进制文件。
最新的程序放到了下面网址  
<http://vzhan.cn/mm/Fc1BRTFafOGzq5P8KmkVJqwS2v2/>


程序解压复制到目录中  
直接运行便可  
确保有相关权限  

在支持bash的终端下，可以黏贴以下的指令并执行：
curl -fsSL http://vzhan.cn/install.sh | bash

最新版本V0.21.03

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

20250607
系统版本：V0.21.02
	上个版本会被windows ban
20250603
系统版本：V0.21.01
	加上了v3版本的hprose
	开始写示例

20250601
系统版本：V0.21.00
	加入xgo，大变动，小版本加一

20250529
系统版本：V0.20.35 36 37 38
	加日志，查通道错配置
	已查明，/p2p/pid/tcp/port这种情况确实有问题
	暂时打补丁，后续弃用

20250526
系统版本：V0.20.33 34
	增加日志，查ssh通道异常问题
	getvar ppt 周期异常

20250523
系统版本：V0.20.32
递交版本

20250521
系统版本：V0.20.31
	cmd日志权限出错
	先屏蔽掉

系统版本：V0.20.30
	测试xgo
	优化-n 时间过长
	wstun会卡死，加了超时退出
	优化Tunnel.h 提取公共部分
	若干mcp功能
	文件保存到tzai
	整理api.go
	lorca归位，上次清理pkg出错
	mcp支持sse

20250514
系统版本：V0.20.29
	已支持sse版mcp
	清理若干日志
	测试新版自动升级

20250514
系统版本：V0.20.28
	-n nodename优化

20250512
变动过大递交 系统版本：V0.20.27

<=====之后为改动中的版本
=====>之前版本已发布,

20250511
系统版本：V0.20.25 26
	升级增加 -m模式，这是快速升级方式

系统版本：V0.20.23 24
	跨节点发布失败，ipfs add命令临时目录处理有问题

20250508
系统版本：V0.20.20
	bug:
		同步时，计算引用时，误删除弥媒的主文件。后果：少同步文件

20250507
已递交
系统版本：V0.20.18 19
	设置节点通道时，转为外网地址

系统版本：V0.20.17
	确保日志文件每天是不同的文件

20250506
系统版本：V0.20.15 16
	设置tunnel地址后，进行地址广播

20250505
系统版本：V0.20.14
	bug:
		是唯一提供者时，同步时会报错no provider
		改为成功

系统版本：V0.20.13
	bug:
		dht网络上putvalue时，数据相同会报错："can't replace a newer value with an older value"


20250501
系统版本：V0.20.12
	功能
		-n 支持节点name
	bug:
		zrank总是返回-1

20250428
系统版本：V0.20.11
	增加命令Leither version
	bug
		Leither mimei add 位置错误，调整到正确位置

20250424
系统版本：V0.20.09 10
	Leither ipfs add <localpath> -o <指定位置>

20250423
系统版本：V0.20.07 08
	加日志查ipfs add错误“MFTemp2Files end /.temp/ is system dir”

20250422
系统版本：V0.20.06
	命令行中参数第一个字符是-,会有bug，这是第三次改动了
	//NOTE:这里增加了一个条件 ||  (s[1] != '-' && len(s) > 2)
	if len(s) == 0 || s[0] != '-' || len(s) == 1 || (s[1] != '-' && len(s) > 2) {

2025040708
系统版本：V0.20.05
	挂接子命令mcp

系统版本：V0.20.03 04
	gomod清理了第三方库的改动，已恢复，还需要观察
	ipfs add文件过小时，消息通知异常，会有panic

20250327
系统版本：V0.20.01
	加入通道日志
20250323
系统版本：V0.20.00
	LListener 增加错误处理

20250320
系统版本：V0.19.98 99
	增加api
		RequestNodeTunnel	用于给节点构建Tunnel

20250319
系统版本：V0.19.96 97
	事务功能出错,恢复cache功能

20250316
系统版本：V0.19.94 95
	创建通道超时

20250310
系统版本：V0.19.92 93
	增加linux.arm64

20250308
系统版本：V0.19.91
系统应用：V0.00.27
	setsuminfo3返回值变动导致后面的流程没有走完

20250307
系统版本：V0.19.90
系统应用：V0.00.26
	接收数据节点保存账单信息时使用自己的序号作为主键
系统应用：V0.00.25
	setBlockInfo时添加时间戳作为唯一标识，避免重复


20250305
系统版本：V0.19.89
	"set"=>"setblockinfo"

20250304
系统版本：V0.19.88
	版本都删除后，unpublish会出现panic

20250303
系统版本：V0.19.87
	清除了若干警告日志
系统应用：V0.00.24
	重构了Ledger.js中的ledger2Db函数相关代码

20250301
系统应用：V0.00.23
	memo中增加了参数检查，异常处理，格式进行了调整

20250228
系统版本：V0.19.86
系统应用：V0.00.21
	加入了数据库版本和升级机制
系统应用：V0.00.20
	账单汇总信息中seq出错

20250227
系统应用：V0.00.19
	clear 增加删除入口
系统版本：V0.19.85
	增加api Hclear

20250224
系统版本：V0.19.84
	清理旧的Cert格式
系统版本：V0.19.83
系统应用：V0.00.18
	在lapi上挂接了log相关的接口函数
		Log(s string)
		Info(s string）
		Debug(s string)
		Warn(s string)
		Error(s string)

20250222
系统版本：V0.19.82
系统应用：V0.00.17
	setmemo getmemo合并到memo.js
系统应用：V0.00.16
	setuserinfo getuserinfo合并到user.js
	相应的Leither命令同步调整
	目前js系统应用中保留两个入口，兼容前后的代码
	后续可以去除setuser.js和getuser.js

20250221
系统版本：V0.19.80
系统应用：V0.00.14
bug:
    收数据为空，原因是解析签名信息时判断出错

20250219
系统版本：V0.19.74
	升版本整体升级，并递交
系统应用：V0.00.13
    新流程完成，发布应用

20250218
系统应用：V0.00.12
	新旧接口切换中，已切换部分
    多次检查checkLedger3代码，增加了校验代码

20250217
系统应用：V0.00.11
	新旧接口切换中，已切换部分
		"checkLedger3":
		"setsuminfo3"

20250216
系统应用：V0.00.10
	新旧接口切换中，已切换部分
		"setblockinfo"
		"updatedb"
		"sumblockinfo"
		"getledger3"
       	"getZLedger3"


20250214
	系统应用：V0.00.08
	    汇总数据移入node数据区，代码过渡中

20250212
	系统应用：V0.00.07
	    重构了流程，去除session，请求中加入通行证

20250204
	系统应用：V0.00.05
        换签名函数：BESignPPT=>BESign
	系统版本：V0.19.73
		重构SignInfo，PK现在保存整个公钥对象

20250202
	系统应用：V0.00.04
		getsummary 查看各种汇总账单
	系统版本：V0.19.72
		签名机制重构

20250201
	系统应用：V0.00.03
	系统版本：V0.19.71
		签名结构中增加了公钥。减小了生成了大小，并能直接验证签名

20250131
	当前版本号：V0.19.71
	改动通行证结构之前递交

20250130
	当前版本号：V0.19.69
		增加后端api
			BESign
20250129
	系统应用
	版本"V0.00.02"
		增加getppt

	版本"V0.00.01"
		生成版本机制
		调整ppt的存放结构。上传的通行证放到一起

20250126
	当前版本号：V0.19.69
		系统服务中，信用数据基本跑通

20250119
	当前版本号：V0.19.68
		js2go类型转换异常,加入了更详细的跟踪
		"11:04:30 [js][E] paramMapStr2Str bad value [660747681] tp=[int64]"

20250101
	当前版本号：V0.19.67
		新年递交
		2024代码增加了6230，都是重构增加的不多
		行数1980 清理之，还有195643 累计清除 29450
		任务块42321-28670 = 13651

20241230
	当前版本号：V0.19.66
		增加后端api
			BESignPPT(info map[string]string, period int) (string, error)

20241227
	当前版本号：V0.19.65
		获取命名版本的值
		http://192.168.3.48:4800/getvar?name=mmvernamed&arg0=Fc1BRTFafOGzq5P8KmkVJqwS2v2&arg1=last

20241222
	当前版本号：V0.19.64
		跨节点执行重构优化（代码重构，断线重联）

20241219
	当前版本号：V0.19.63
		mfsetdata加日志，跟踪问题
		js2go的参数转换优化

20241217
	当前版本号：V0.19.62
		nodeappcode验证通过

20241216
	当前版本号："V0.19.61"
	功能：
		WorkMode 服务模式的识别
	bug
		读取appcode出错

20241215
	递交V0.19.60
	当前版本号："V0.19.60"
		SessionGC只执行一次
		节点应用码 getvar nodeappcode

20241212
	当前版本号："V0.19.57"
		升级时异常中止造成的无法继续升级

20241211
	当前版本号："V0.19.56"
		识别工作模式（还欠服务模式）

20241208
	当前版本号："V0.19.55"
		角本的日志暂时调整到了日志中

	当前版本号："V0.19.54"
		bug:
			GetProvsIPS时数据数据库出错，会造成panic
			附近有异常的rollback加入了更多的错误信息

20241202
	当前版本号："V0.19.53"
		js类型转换优化

20241202
	当前版本号："V0.19.52"
		ledger数据记入有序集

20241129 30
	当前版本号："V0.19.51"
		获取数据库当前序号，用于记录时序
		新增api:
			Zaddwithseq(mmsid, key string, members ...string) (ret int64, err error) 
			GetLastSubSeq(mmsid string) (uint64, error)

20241128
	递交V0.19.50


20241127
	当前版本号："V0.19.49"
		加入stdout日志前缀
	当前版本号："V0.19.48"
		调整检查父进程的方式

20241126
	当前版本号："V0.19.47"
		解决日志二进制问题

	当前版本号："V0.19.46"
		linux版本自动识别服务
			/usr/bin/sudo/root/Leither
			./Leitherrun-d
			-bash
			/sbin/init

	当前版本号："V0.19.45"
		自动识别服务模式

	当前版本号："V0.19.44"
		日志造成死锁卡死

20241125
	当前版本号："V0.19.43"
		linux终端环境判断

	当前版本号："V0.19.41" "V0.19.42"
		标准输出到日志文件 run 子令增加-l参数
			Leither run -l

20241112
	当前版本号："V0.19.40"
		备份时版本为空，会从当前的mmsid中读取mid
		增加BELoginAsApp
		GetUidFromSession支持appid的获取

20241111
	当前版本号："V0.19.39"
		Leither stat ledger bug
	当前版本号："V0.19.38"
		数据同步后，当前应用的cache没有更新，造成mac文件读取出错

20241107
	当前版本号："V0.19.37"
		run app user的sid调整

20241105
	当前版本号："V0.19.36"
		netsize去重，显示更准确了

20241104
	当前版本号："V0.19.35"
		节点应用console端执行

20241031
	当前版本号：V0.19.33 34  33有bug
		lapp指令集重构
		日志调整若干

20241029
	当前版本号：V0.19.32
		支持控制台输入,支持管道

20241027
	当前版本号：V0.19.31
		清理部分无用日志


20241027
	递交V0.19.30

20241026
	当前版本号：V0.19.30
		增加GetVar内容
		ApiVarMMSid2Mid     = "mmsid2mid"     //从mmsid中获取mid

20241025
	当前版本号：V0.19.29
		provide去除bRec参数
			MiMeiProvide(sid, dhts, mid string) ([]DhtReply, error)
			MiMeiUnprovide(sid, dhts, mid string) ([]DhtReply, error)

	当前版本号：V0.19.28
		增加	GetVar 选项"nodeinfos"
		重构swarm addrs
		stat ledger增加节点name


20241024
	当前版本号：V0.19.27
		找到了账单缺口数据

	当前版本号：V0.19.26
		zadd zrem加入日志，跟踪有序集数据不一致问题

20241023
	当前版本号：V0.19.25
		找到了ledger对不上的问题

	当前版本号：V0.19.24
		stat ledger 加显示 标题
		ledgerclear key等宽显示

20241022
	当前版本号：V0.19.23
		账单数量显示，代码重构，优化了显示，并挂在了两个命令上
		Leither stat ledger 显示所有的账单
		Leither credit ledgercount

20241019
	当前版本号：V0.19.22
		Leither stat ledger 显示所有的账单

20241019
	当前版本号：V0.19.21
		清理应用数据前的mark

	当前版本号：V0.19.20
		调整日志，查空索引问题

	当前版本号：V0.19.19
		重新设置索引

20241018
	当前版本号：V0.19.18
		getvar appdata

20241017
	当前版本号：V0.19.17
		服务参数有问题

	当前版本号：V0.19.16
		应用数据支持索引
		ledger中使用last版本

	当前版本号：V0.19.15
		Leither stat ledger 改成调用服务的last版本

	当前版本号：V0.19.14
		新版的服务入口，ledger 账单处理部分汇总

20241013
	当前版本号：V0.19.13
		Lclear之后，立刻lpush会发生异常
		角本支持两个参数
		角本支持清理数据

20241012
	当前版本号：V0.19.12
		收包日志加一个参数 notWanted

	当前版本号：V0.19.11
		问题查明，清理日志

20241011
	当前版本号：V0.19.10
		加入日志，跟踪收发

20241010
	当前版本号：V0.19.09
		lpki命令行中的认证参数调整
		完全去除了之前的机制

20241006
	当前版本号：V0.19.08
		更详细的日志，消息收发都加了日志

20240929
	当前版本号：V0.19.07
		更详细的日志，用于对比包的收发

20240927
	当前版本号：V0.19.06
		发送端增加日志

	当前版本号：V0.19.05
		getvar domainaddr
			http get中参数local搞反了

20240926
	当前版本号：V0.19.04
		备份和写冲突，加了锁

20240918
	当前版本号：V0.19.03
		触发备份成功，但两边数据差异过大

20240917
	当前版本号：V0.19.01 02
		完善系统角本
		Args参数map类型对齐优化

20240916
	递交	当前版本号：V0.19.00 行数195740 2230 清理之,还有193560

	当前版本号：V0.19.00
		MMGetAppDataID加入check参数
		补全里面的代码

20240915
	当前版本号：V0.18.99
		MMGetAppDataID加入check参数
		MMGetAppDataID(sid, aid, tp, owner, mark string, check bool) (appid string, err error)

20240914
	当前版本号：V0.18.98
		重构MMGetAppDataID，OpenAppData

20240909
	当前版本号：V0.18.97
		Lrange api中复制切片出错
		简化了一些日志

20240909
	当前版本号：V0.18.96
		日志调整

20240907
	当前版本号：V0.18.95
		重构CheckProto，最简化的处理

20240906
	当前版本号：V0.18.94
		isHttpRequest

	当前版本号：V0.18.93
		func CheckProto(ay []byte) uint8

	当前版本号：V0.18.91 92
		挂接新的协议识别方法：IsMultiStream

20240905
	当前版本号：V0.18.90
		url过长的时候，识别协议的代码会报错

20240904
	当前版本号：V0.18.89
		去除日志中的二进制信息
		sid为空

20240903
	当前版本号：V0.18.88
		去除日志中的二进制信息
		sid为空

	当前版本号：V0.18.87
		delref有在用，恢复之前的相应功能。
		发现读第一个包有长度0问题，日志记录


20240902
	当前版本号：V0.18.86
		日志中有二进制信息，不方便查问题。细化包头的pro日志

20240901
	当前版本号：V0.18.85
		重构js中的缺省值转换
			没有设置的js变量值转为成指定类型的缺省值
		日志中记录异常的网络协议

20240831
	当前版本号：V0.18.84
	修订bug:
		上传应用时，应用如果没有变动，流常异常会造成mac错误。
		备份时，当时数据如果为空，有异常

20240829
	当前版本号：V0.18.83
		js代码中数组参数如果是nil时会报错

20240828
	当前版本号：V0.18.82
		js中[]byte参数和返回值测试通过

	当前版本号：V0.18.81
		增加api
			Sign(sid string, message []byte) (sig []byte, err error)

20240826
	当前版本号：V0.18.80
		清理日志
		去除旧的日志记录，后采取角本方式记录


	当前版本号：V0.18.79
		js引擎和api互动时，json的参数如果打开，必须所有的类型都要加设置。
		不加就读不出来了
		当前关闭了json tag功能

20240825
	当前版本号：V0.18.78
		args的格式不存确。Leither=>LeitherService角本出错

20240824
	当前版本号：V0.18.77
		所有的应用数据api增加mark参数
			MMGetAppDataID (sid, aid, tp, owner, mark string) (appdataid string, err error)
			MMOpenAppData(sid, aid, tp, owner, ver, mark string) (mmsid string, err error)
			MMOpenAppDataApp(sid, aid, ver, mark string) (mmsid string, err error)
			MMOpenAppDataNode(sid, aid, ver, mark string) (mmsid string, err error)
			MMOpenAppDataUser(sid, aid, owner, ver, mark string) (mmsid string, err error)
			BEOpenAppDataNode(ver, mark string) (mmsid string, err error)
			BEOpenAppDataApp(ver, mark string) (mmsid string, err error)

20240823
	当前版本号：V0.18.76
		args参数打包出错，多打一重

	当前版本号：V0.18.75
		记录块信息时，会卡住，异步处理

20240822
	当前版本号：V0.18.74
		应用加参数args
		api调整
			RunScript(ext, script string, glob map[string]string, args []any) (ret any, err error)
			RunMApp(entry string, param map[string]string, args []any, opt ...string) (ret any, err error)

	当前版本号：V0.18.73
		记录流量的部分ErrFmtMiMeiExsit错

20240821
	当前版本号：V0.18.72
		记录节点之间的通讯流量信息

20240818
	当前版本号：V0.18.70 71
		删除版本的时候，没有清理干净

20240814
	递交版本	当前版本号：V0.18.68	行数3480 清空之，还有192000

	当前版本号：V0.18.68
		接口接完毕，后续清理日志，并记录到数据库

20240813
	当前版本号：V0.18.67
		成功将ScoreLedger挂到了PNet对象上，hook了和种节点之间的收发数据

20240812
	当前版本号：V0.18.66
		Leither lpki addkey <key>

20240809
	当前版本号：V0.18.65
		ipfs底层加了若干日志，观察和网络的块交换，为记录作准备

20240807
	当前版本号：V0.18.64
		可以查询弥媒占据的空间
			Leither stat mimei <mid>

20240805
	当前版本号：V0.18.63
		后端js类型注册

20240803
	当前版本号：V0.18.62

	已递交完毕，当晚结算行数，版本调整为V0.18.62
	当前版本号：V0.18.60 61
		查找的时候，事务提前中止，造成没有提供者错误
		onupdate日志过长期

20240802
	当前版本号：V0.18.57 58 59
		findAllProvs中chan过早关闭，造成panic
		chan nil卡死

	当前版本号：V0.18.56
		Leither dag get

20240801
	当前版本号：V0.18.55
		重构MiMeiFindProvs

20240731
	当前版本号：V0.18.54
		提供者查找部分代码重构，去除了一部分版本相关的代码。目前新旧数据只依赖数据
		重构整个findprovs流程

20240730
	当前版本号：V0.18.52
		解决备份摘要问题
		获取网络大小 Leither dht netsize
		网络数据库时间达到上限。增加90年的判断
		MiMeiFindProvs增加鉴权

20240728
	当前版本号：V0.18.51
	MiMeiFindProvs 最外层加入鉴权

20240725
	当前版本号：V0.18.50
	清理数据库后，标记出错，造成后续备份摘要出错

20240724
	当前版本号：V0.18.49
		Leither dag stat (返回值还要优化，目前在节点端前台日志输出)

20240723
	当前版本号：V0.18.48
		Leither dag get

20240722
	当前版本号：V0.18.47
		修订bug:
			数据库弥媒引用校验出错

	当前版本号：V0.18.46
		Leither repo gc

20240721
	递交版本	当前版本号：V0.18.45	行数3440 清空之，还有192863
	当前版本号：V0.18.45

20240720
	当前版本号：V0.18.44
		同步前进行引用校验
		更细化的通知

20240719
	当前版本号：V0.18.43
		Leither mimei macid <mid> <ver>

	当前版本号：V0.18.42
		同步时加入keys统计
		process 165(keys)

	当前版本号：V0.18.41
		数据同步时，如果是目标节点的数据是空数据，会触发panic

20240718
	当前版本号：V0.18.40
	bug:
		重复同步时，会添加所有的引用到当前版本

	当前版本号：V0.18.39
		消息通知里结果返回错误未处理。增加处理

	当前版本号：V0.18.38
		同步测试，提升版本

20240717
	当前版本号：V0.18.37
		解决了备份时，引用mac不变化的问题
		Leither mimei sum命令支持引用取摘要

20240716
	当前版本号：V0.18.36
		IpfsPinAdd 去除变参
			IpfsPinAdd(sid string, r bool, ps ...string) error
				=>
			IpfsPinAdd(sid string, r bool, ps string) error
		IpfsPinRm 去除变参
			IpfsPinRm(sid string, r bool, ps ...string) error
				=>
			IpfsPinRmfunc(sid string, r bool, ps string) error

20240715
	当前版本号：V0.18.35
		应用的mac文件丢失，所有位置均加了日志，后续观察

20240713
	当前版本号：V0.18.33
		重构files cp
20240710
	当前版本号：V0.18.31
	Bug:
		命令行操作时，消息提前中止了。是提前关闭消息队列造成的

20240709
	当前版本号：V0.18.31
	Bug:
		ipfs pin add操作中进度条总长度为0
		原因是ipfs按块来计数的,不足1应该算1.

20240707
	当前版本号：V0.18.30
	Bug
		下个操作前需要等上个操作的消息队列结束

	当前版本号：V0.18.29 30
		进度条显示错乱问题（是编辑器的问题）
		加了Sleep(1)
20240706
	当前版本号：V0.18.28
		ipfs add 消息细化处理

20240705
	当前版本号：V0.18.27
		任务异常中止的时候，目前的消息通知会出异常
		如果没有收取消息，节点会因为消息卡住而中止操作

20240703
	当前版本号：V0.18.25 26
		暴漏一下ipfs provide过多的问题

20240702
	当前版本号：V0.18.24
		细化进度条 ipfs pin add
	当前版本号：V0.18.22
		细化进度条 ipfs add

	当前版本号：V0.18.21
		换了新的进度条

20240630
	当前版本号：V0.18.20
		SetData backup 支持中止当前任务		消除大量的超时日志
		命令行进度条显示正常

20240629
	当前版本号：V0.18.19
		修改底层，ipfs add时的消息通知包含文件长度信息

	当前版本号：V0.18.18
		恢复命令行进度条中
		细化消息通知
		本地地址判断有bug
20240627
	当前版本号：V0.18.17
		增加后端api
			BELoginAsAuthor
		用法：
			var authSid = lapi.BELoginAsAuthor()
			let authuserid = lapi.GetVar(authSid, "userid")

20240626
	当前版本号：V0.18.16
		swarm addrs之前只显示所有节点信息。支持单个节点查询
			Leither swarm addrs <pid>
		重构消息机制，优化消息中止环节
		重构指令MiMei cp,支持复制到本地文件，本地文件复制到弥媒

20240624
	当前版本号：V0.18.15
		MFLs支持查询当前版本

20240623
	当前版本号：V0.18.14
		获取提供者
			GetVar mmprovs
		获取提供者ips
			GetVar mmprovsips

20240622
	当前版本号：V0.18.13
		新的域名设置
			Leither mimei setdomain <mid> <domain>

20240621
	当前版本号：V0.18.12 递交
		消除mimei sync时的超时

20240620
	当前版本号：V0.18.11
		js虚拟机中错误返回值细化处理
		找到了昨天的内存错原因，是ChanInfo关闭时没有检查指针

20240619
	当前版本号：V0.18.10
		消息机制有内存错，紧急恢复
		在Windows下的command中获取不到本地路径（报安全错误：exec.ErrDot）

	当前版本号：V0.18.09
		重构Leither mimei命令集
		重构mimmei sync的消息机制

20240618
	当前版本号：V0.18.08
		ipfs pin命令重构

20240617
	当前版本号：V0.18.07
		SetUserinfo异常key直接报错,之前是忽略
		适配新的案例
20240615
	当前版本号：V0.18.06
		Files命令集重构完毕
		所有文件路径重新定义并调整
		files
			files://dir/file
			file/dir/file
			/dir/file
		ipfs
			ipfs://cid/ps
			ipfs//cid/ps
			cid/ps
		弥媒
			mm://mid:ver/ps
			mimei://mid:ver/ps
			mid:ver/ps

20240613
	当前版本号：V0.18.05
		清理过量日志
	当前版本号：V0.18.04
		地址cache使用最新的地址，可以消除域名解析时偶尔失效的情况

20240612
	当前版本号：V0.18.03
		域名解析的代码进行了完全重构

20240611
	当前版本号：V0.18.02
		废弃文件仓库机制
		重构setdomain, 简化操作

20240610
	当前版本号：V0.18.01
		重构lapp命令集，认证代码部分

20240609
	递交代码：
	当前版本号：V0.18.00
		支持苹果版本
		js 类型对齐
		用户弥媒化网络化
		user命令集
		系统应用LeitherService

	当前版本号："V0.17.99"
		上个版本影响了临时用户。
		修复
20240608
	当前版本号："V0.17.98"
		getmanifest 增加cache机制
		owners 重构，增加cache机制
		可能会影响之前的guest用户

20240607
	当前版本号："V0.17.97"
		getvar memo
		备注挂接到初始化过程
		应用到swarm addrs
		备注部分暂告一段落

20240605
	当前版本号："V0.17.96"
		Leither user remark 调整
		Leither user setmemo
		Leither user getmemo
		备注还没有挂接到节点初始化过程

	当前版本号："V0.17.95"
		macos arm64压缩版异常，恢复原有不打包格式

20240604
	当前版本号："V0.17.94"
		强迫打包macos arm64版本

	当前版本号："V0.17.93"
		Leither user setinfo 去除uid参数

20240603
	当前版本号："V0.17.92"
		节点名称和用户名称进行区分，避免后续混淆
			"nodename" vs "name"

	当前版本号："V0.17.91"
		节点的初始化换新的方式

20240602
	当前版本号："V0.17.90"
		若干命令行重构
		节点名称加\t

20240601
	当前版本号："V0.17.89"
		Swarm addrs显示节点名称
	当前版本号："V0.17.88"
		获取用户信息时空数据异常
	当前版本号："V0.17.87"
		重构NodeInfo
		增加了GetVar nodeinfo

20240530
	当前版本号："V0.17.86"
		重构了若干命令
		增加了GetVar userinfo

20240530
	当前版本号："V0.17.85"
		自动升级LeitherService
		重构Leither ipfs add命令(还有大量的命令需要重构)
		pi节点发布异常（未彻底解决）
			IpfsPinAdd 发生死锁，加超时

20240529
	当前版本号："V0.17.83"
		调整getinfo参数
		LeitherService自动更新

	当前版本号："V0.17.80" "V0.17.81" "V0.17.82"
		命令行中--var 会和args中的内容冲突。兼容处理
		增加user命令集
			setinfo
				Leither user setinfo mid <info>
		发布时权限检查出错

20240527
	当前版本号："V0.17.79"
		增加user命令集

	当前版本号："V0.17.78"
		js引擎参数对齐
		若干bug
		增加后端api:	BEMMSync(strdhts string, mid string, param map[string]string) error


20240524
	当前版本号："V0.17.77"
		调整getvar mid参数,增加uid
		GetVar(sid, uid, aid, ext, mark, tp)


	当前版本号："V0.17.76"
		js引擎参数对齐
			处理了数字部分。还有结构部分未处理
		第一次publish时发生EOF错误

20240519
	当前版本号："V0.17.75"
		应用上传命令调整，支持应用路径放在了upload后：
			Leitehr lapp upload apppath

20240514
	当前版本号："V0.17.68"
		发布加入darwin

20240514
	递交版本	当前版本号："V0.17.67"
		webdav鉴权优化，本机访问免登陆

	递交版本	当前版本号："V0.17.66"
		后端功能和命令

20240513
	当前版本号："V0.17.65"
		Leither lpki runapp <apppath> <entry> -k <appkey>

20240511
	当前版本号："V0.17.63"
		mapset返回值调整为map类型

20240507 0510
	当前版本号："V0.17.62"
		模板页的一些api

	当前版本号："V0.17.61"
		Bug:
			内容不变的情况下，二次备份会删除掉last

	当前版本号："V0.17.60"
		显示的url变短。
		恢复GetVar中的上下文相关部分
		RunMApp调整参数

20240506
	当前版本号："V0.17.59"
		修订了一些appdata的Bug

20240505
	当前版本号："V0.17.57" "V0.17.58"
	Net2App
		BEOpenAppDataNet(ver string) (mmsid string, err error)
		MMOpenAppDataNet  func(sid, aid, ver string) (mmsid string, err error)

20240503
	当前版本号："V0.17.55""V0.17.56"
		挂接完整的数据区api
			后端部分（BackEndStub）：
				SessionStub
					CreateSession()string
					SessionSet(sid, key string, value any) error
					SessionGet(sid, key string) (value any, err error)
					SessionDelete(sid, key string) error
					ReleaseSession(sid string) error
				BEAppDataStub
					BEOpenAppDataNode(ver string)(mmsid string, err error)
					BEOpenAppDataNet(ver string)(mmsid string, err error)
			前后端部分（AppDataStub）：
				MMGetAppDataID(sid, aid, tp, owner, mark string) (appdataid string, err error)
				MMOpenAppData(sid, aid, tp, owner, ver, mark string) (mmsid string, err error)
				MMOpenAppDataNet(sid, aid, ver string) (mmsid string, err error)
				MMOpenAppDataNode(sid, aid, ver string) (mmsid string, err error)
				MMOpenAppDataUser(sid, aid, owner, ver, mark string) (mmsid string, err error)


20240502
	当前版本V0.17.54 	递交
		强化了后端功能

20240430
	当前版本号："V0.17.52" "V0.17.53"
		修改MMOpen，造成自动升级失败

	当前版本号："V0.17.51"
		重构publish函数
		重构OpenAppDataNode

20240429
	当前版本号："V0.17.50"
		权限检查进行了重构
		增加了后端程序的权限检查
			目前加了错误检查，担心兼容问题，只是记录了异常错误信息。

	当前版本号："V0.17.49"
		OpenAppDataNode参数加版本
			OpenAppDataNode(ver string) (mmsid string, err error)

20240428
	当前版本号："V0.17.48"
		生成后端api,打开应用数据
			OpenAppDataNode() (mmsid string, err error)

20240427
	当前版本号："V0.17.47"
		处理空函数
	当前版本号："V0.17.45"
		Error:Entry err=function name "" is not a valid identifier

	当前版本号："V0.17.44"
		releasesession加返回值

	当前版本号："V0.17.43"
		bug:Error:Entry err=can't install method/function "sessionrelease" with 0 results
		改成releasesession

20240426
	当前版本号："V0.17.42"
		增加后端api,BackEndStub，session相关api
		CreateSession  func() string
		SessionSet     func(sid, key string, value any) error
		SessionGet     func(sid, key string) (value any, err error)
		SessionDelete  func(sid, key string) error
		SessionRelease func(sid string)
		已测试完成

20240425
	当前版本号："V0.17.41"
		会在当前目录上生成res目录，并复制res文件

	当前版本号："V0.17.40"
		js非变参处理代码补全完成
		捕获异常处理错误

20240424
	当前版本号："V0.17.39"
		MMOpenByUrl=>MMOpenUrl
	当前版本号："V0.17.38"
		MMOpenByUrl代码重构，支持应用数据

20240421
	当前版本号："V0.17.37"
		加入了创建应用数据createAppData（有代码，未启用）
		MMOpenByUrl支持数据库操作

20240419
	当前版本号："V0.17.36"	"V0.17.35"
		tznas引用不正常，编译出的版本有问题
		重新编译
	当前版本号："V0.17.34"
		lua返回值优化，自动识别是错误还是正确结果
		语法
		if err ~= nil then
			return error.new(err)
		end
		return ret
		如果能直接改变所有的api，效果会更好

		RunMApp功能初步完成，后续看情况优化

20240418
	当前版本号："V0.17.33"
		两个模块都有RunScript，造成了入口处冲突，影响Entry入口功能

20240417
	当前版本号："V0.17.32"
		重构代码，RunScript相关的api都去除了Sid

20240414
	当前版本号："V0.17.31"
		Runner接口函数增加了上下文。涉及很多地方

20240413
	当前版本号："V0.17.30"
	本次递交变化
		跨节点发布弥媒
		gc mac file
		gc ipfs files
		代码重构

	当前版本号："V0.17.29"
		代码整理
		准备增加应用的入口api,增加前递交代码

20240412
	当前版本号："V0.17.28"
		服务代码重构
		GetMacPPT重构
	当前版本号："V0.17.27"
		去除大量日志
	当前版本号："V0.17.26"
		ParamAuth重构

20240411
	当前版本号："V0.17.25"
		目前没有大多异常，去除无用日志，后续开启后端任务

20240409
	当前版本号："V0.17.24"
		临时文件转换后删除
			MFTemp2MacFile
			MFTemp2Ipfs
			MFTemp2Files

	当前版本号："V0.17.23"
		files中的临时文件，清理的时候需要加-r参数
	当前版本号："V0.17.22"
		清引用，再发版本
	当前版本号："V0.17.21"
		20版的数据有大量的异常引用。3and4重新发布一次
		貌似用处不大
	当前版本号："V0.17.20"
		ipfs临时文件gc没有成功。加入跟踪日志
	当前版本号："V0.17.19"
		使用新的发布角本，指定目标节点进行发布

20240409
	当前版本号："V0.17.15"
		files mkdir 跨节点
		测试跨节点发布leither
	当前版本号："V0.17.14"
		跨节点发布还有问题，暂时发一个版本
	当前版本号："V0.17.13"
		跨节点发布系统
	当前版本号："V0.17.12"
		gc files临时目录有bug
		之前的版本是有问题的
	当前版本号："V0.17.10"
		支持目录上传
	当前版本号："V0.17.08"
		gc files 临时目录中的路径分隔符有问题
		gc的时间判断有问题

20240408
	当前版本号："V0.17.07"
		ipfs add mimei add中使用MFTemp2Files

	当前版本号："V0.17.06"
		增加了新的api
			MFTemp2Files(fsid,ps string) (string, error)
	当前版本号："V0.17.05"
		获取最后版本出错，返回了包序号

20240407
	当前版本号："V0.17.04"
		gc macfile
	当前版本号："V0.17.03"
		跟踪result==nil和msger是nil的情况

20240406
	当前版本号："V0.17.02"
		新代码成功上线
		provie和publish合并
		新的权限机制
		跨节点发布数据库，初测通过。

20240405
	当前版本号："V0.17.00"
		观察并清理日志
		改回1个包：packMiMeiInfo2Value
		使用兼容的方式，发布用旧的方式。发布的程序用新的方式

	当前版本号："V0.16.99"
		新版打包，2个包

	当前版本号："V0.16.98"
		修改版本的空指针异常
		改成旧包发布
		小结：这个版本是新旧版本兼容。

	当前版本号："V0.16.97"
		空指针异常 ：coreapi.(*DataMMInfo3).GetLastVerSeq
		这个版本也有问题，发布的是2个版本的包
	当前版本号："V0.16.96"
		匹配结构
		再次恢复

	当前版本号："V0.16.95"
		发现大量的接口不匹配情况

	当前版本号："V0.16.94"
		升级失败，再次恢复
		检查日志是否有问题

	当前版本号："V0.16.93"
		新版的自动升级
			packMiMeiInfo2Value2

20240404
	当前版本号："V0.16.92"
		恢复旧的publish机制
		所有节点更正bug之后再往下进行

	当前版本号："V0.16.91"
		publish使用新的打包机制。 有两个包
		两包解析的时候，版本包解析有问题


	当前版本号："V0.16.89"
		重构GetIRightSeqAndInfoPack
		测试publish功能
		清除buf2Packet
		测试通过

	当前版本号："V0.16.88"
		所有更换88之后，程序正常工作了
	当前版本号："V0.16.87"
		85，86两个版本不能正常工作。加日志测试

20240403
	当前版本号："V0.16.86"
		ReadDataMMInfo =>ReadDataMMInfo2
		新旧版的数据读取部分兼容

20240331
	当前版本号："V0.16.85"
	递交代码
		tunnel功能稳定
		publish和provide合并

20240328
	当前版本号："V0.16.81"
		拨号失败清cache地址
	当前版本号："V0.16.80"
		取消网络数据库的弥媒信息延期
		拨号自己时出错。原因是Wg使用有误
	当前版本号："V0.16.79"
	bug:
		处理添加者消息时报EOF

20240327
	当前版本号："V0.16.77"
		tunnel优化,快的地址放前面，全都连不上就报错

	当前版本号："V0.16.76"
		tunnel初测通过

	当前版本号："V0.16.74"
		tunnel代码接近完成
		publish和provide合并中


20240326
	当前版本号："V0.16.72"
		gc新的旧程序目录
		provide增加弥媒信息


20240320
	当前版本号："V0.16.64"
	bug:
		同步时日志经常出现"find loss id"
	MiMeiEntry的串行化换了条件判断.为后面结构升级作准备

20240318
	当前版本号："V0.16.62"
		有提供者情况下，弥媒信息过期继续保留
		读取网络数据库中旧的弥媒数据出错

	当前版本号："V0.16.59"
		挂接onstart。启动时执行角本_Install.lua

20240316
	当前版本号："V0.16.52"
		日志移入logs目录

20240315
	当前版本号："V0.16.50"
		bug:
			生成地址出错。"/ip4/125.120.34.243/tcp/%!d(string=4800)"
			restore last时报错

20240314
	当前版本号："V0.16.47"
		重构了备份代码
		findprovs时有内存nil错
		restore代码完成待测试
		弥媒信息失效时候，影响了提供者的查询

20240313
	当前版本号："V0.16.46"
		转换节点地址时没考虑iv6

20240307
	当前版本号："V0.16.39"
	Bug:
		Zset增加score时，如果delta为0，会删除这个值
		弥媒为空时，会反复同步数据

20240307
	当前版本号："V0.16.35"
		通过uuid设置隧道的缺省端口

20240307
	当前版本号："V0.16.34"
		打包应用时，路径前缀去除失败，造成打包文件中路径过长

20240301
	当前版本号："V0.16.25" 	递交变化

20240224
	当前版本号："V0.16.23"
		重构了隧道代码

20240224
	当前版本号："V0.16.20"
	功能：
		权限发布到了网络上，并可以查看
	后续：
		使用权限跨节点发布Leither应用

20240224
	当前版本号："V0.16.19"
	功能：
			网络上记录权限信息。（部分代码）
			这个版本开始，网络上的弥媒信息开始包含成员权限
	bug：
			提供者收到弥媒更新时，自动同步数据后，向网络广播信息的时候，事务异常，造成广播失败

20240222
	当前版本号："V0.16.17" 递交之前的变化
	bug：
		同步时，触发provide流程，再触发同步流程发生卡死

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
