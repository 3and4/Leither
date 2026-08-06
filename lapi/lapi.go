package lapi

// 节点对外提供的功能通过接口描述
// 功能逐步开放，当前lapi包是已开放对外的功能

// LApi描述节点内使用的api（包括节点内外都可以使用的rpc调用的接口和节点内使用的后端接口）
// 对于Rpc方式，可以通过InitLApiStubByUrl等api以获取stub句柄的方式访问节点功能。
// 对于节点内执行的应用，可以直接通过GetLApi的方式获取api接口
// 所有api最后都使用接口进行了抽象，后续功能介绍都基于接口进行介绍
type LApi interface {
	IBackEnd //节点内使用

	IAuth   //认证部分
	IVarAct //封装了大批Api，GetVar是获取系统变量, Act一类是执行一个动作，这个动作有需求但还没有到需要单独封装为Api的程度。
	IMiMei  //弥媒，封装了文件和数据库操作
	INet    //网络，封装了所有的网络相关操作
}

// 跨节点获取的api,通过rpc方式获取生成
type IRPC interface {
	IAuth   //认证部分
	IVarAct //封装了大批Api
	IMiMei  //弥媒
	INet    //网络部分	TunnelStub，	DNSStub,NodeStub
}

// 目前对外的只有之两个函数
//type IAuth = IAuthStub

// type IVarAct = IVarActStub

// type IMiMei = IMiMeiStub

// type INet = INetStub
