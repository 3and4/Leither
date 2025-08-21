// backend.go
package lapi

/*
BackEndStub

一个应用执行的时候有前端模式和后端模式

前端模式通过rpc操作执行api，运行在节点之外，浏览器，控制台，app等

后端模式执行在节点内，通过一个全局变量lapi来操作Api

后端除了拥有所有的前端功能（ApiStub）外，还有独有的后端api

后端Api目前有两部分，后端Session（SessionStub）和应用数据操作（BEAppDataStub）
*/

type BackEndStub struct {
	//	*ApiStub
	*LApiStub
	*SessionStub
	*BEAppDataStub
}

// 用于xgo,调用前会替换
func GetLApi() *BackEndStub {
	return nil
}

func GetArgs() []any {
	return nil
}

const (
	Request_AppVer  = "ver"
	Request_Author  = "author"
	Request_AppName = "app"
	Request_AppID   = "aid"
	Request_Entry   = "entry"
	Request_MID     = "mid" //内部必然有一个
	Request_Sid     = "sid" //内部必然有一个
	Request_NodeID  = "nid"
	Request_SidApp  = "sidapp" //应用sid,作者身份

	ApiVarNodeAppCode = "nodeappcode"
)

func GetRequest() map[string]string {
	return nil
}

/*
SessionStub

# Session功能类似于cookie，区别点在于前者作用于后端，后者作用于前端

Session可以保存用户的状态和上下文信息。

CreateSession创建一个session并返回id.

通过这个id可以在后端程序中读写删除状态，不同的后端模块之间也可以共享信息

这个状态是存放在内存中的，超过1小时不访问就会被GC清作
*/
type SessionStub struct {
	CreateSession  func() (sid string)
	SessionSet     func(sid, key string, value any) error
	SessionGet     func(sid, key string) (value any, err error)
	SessionDelete  func(sid, key string) error
	ReleaseSession func(sid string) error
}

/*
BEAppDataStub

这里面放的是应用中后端数据操作的api接口。

一、BEOpenAppDataNode(ver string) (mmsid string, err error)
简介：

	这是应用在节点上的打开节点的弥媒

参数：

	ver,弥媒的版本，cur表示当前的可写版本，last表示最后一次备份的版本

返回值：

	mmsid 弥媒的数据id,可用于数据库相关的api操作

二、BEOpenAppDataApp(ver, mark string) (mmsid string, err error)
简介：

	打开节点的弥媒

参数：

	ver,弥媒的版本，cur表示当前的可写版本，last表示最后一次备份的版本

返回值：

	mmsid 弥媒的数据id,可用于数据库相关的api操作

BELoginAsAuthor()(sid string, err error)

简介：

	以作者作者身份登陆

返回值:

	Sessionid，可用于所有需要会话id的api，执行的时候代表作者身份
*/
type BEAppDataStub struct {
	BEOpenAppDataNode func(string, string) (mmsid string, err error)
	BEOpenAppDataApp  func(string, string) (mmsid string, err error)
	BEMMSync          func(strdhts string, mid string, param map[string]string) error
	BELoginAsAuthor   func() (sid string, err error)
	BELoginAsApp      func() (sid string, err error)
	BESignPPT         func(info map[string]string, period int) (string, error)
	BESign            func(info map[string]string) (string, error)
}
