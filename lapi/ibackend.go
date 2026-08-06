// ibackend.go
package lapi

/*
一个应用执行的时候有前端模式和后端模式

前端模式通过rpc操作执行api，运行在节点之外，浏览器，控制台，app等

后端模式执行在节点内，通过一个全局变量lapi来操作Api

后端除了拥有所有的前端功能（ApiStub）外，还有独有的后端api

后端Api目前有两部分，后端Session（SessionStub）和应用数据操作（BEAppDataStub）
*/

// 用于mapp,调用前会替换
func GetLApi() LApi {
	panic("GetLApi() 用于mapp,调用前会替换")
}

func GetArgs() []any {
	panic("GetArgs() 用于mapp,调用前会替换")
}

// 后端专用
type IBackEnd interface {
	ISession
	IBEAppData
	ILog
}

/*

Session功能类似于cookie，区别点在于前者作用于后端，后者作用于前端

Session可以保存用户的状态和上下文信息。

CreateSession创建一个session并返回id.

通过这个id可以在后端程序中读写删除状态，不同的后端模块之间也可以共享信息

这个状态是存放在内存中的，超过1小时不访问就会被GC清作
*/

// ISession 接口定义了Session操作的所有方法
type ISession interface {
	// CreateSession 创建一个session并返回id
	CreateSession() (sid string)

	// SessionSet 设置session中指定key的值
	SessionSet(sid, key string, value any) error

	// SessionGet 获取session中指定key的值
	SessionGet(sid, key string) (value any, err error)

	// SessionDelete 删除session中指定key的值
	SessionDelete(sid, key string) error

	// ReleaseSession 释放指定的session
	ReleaseSession(sid string) error
}

/*
	IBEAppData 接口定义了应用数据操作的所有方法

这里面放的是应用中后端数据操作的api接口。

一、BEOpenAppDataNode(ver, mark string) (mmsid string, err error)
简介：

	以节点身份打开应用数据区域，具有可读可写权限。
	节点身份意味着以当前运行节点的系统身份访问数据，通常用于节点级别的数据管理。
	如果指定的应用数据区域不存在，会自动创建新的数据区域。

参数：

	ver: 数据版本，"cur"表示当前的可写版本，"last"表示最后一次备份的版本
	mark: 标记字符串，用于标识数据区域的用途或上下文

返回值：

	mmsid: 弥媒的数据ID，可用于后续的数据库操作API（如Set、Get、Hset、Hget等）
	err: 错误信息，如果操作成功则为nil

二、BEOpenAppDataApp(ver, mark string) (mmsid string, err error)
简介：

	以应用作者身份打开应用数据区域，具有可读可写权限。
	应用作者身份意味着以应用创建者的身份访问数据，通常用于应用级别的数据管理。
	如果指定的应用数据区域不存在，会自动创建新的数据区域。

参数：

	ver: 数据版本，"cur"表示当前的可写版本，"last"表示最后一次备份的版本
	mark: 标记字符串，用于标识数据区域的用途或上下文

返回值：

	mmsid: 弥媒的数据ID，可用于后续的数据库操作API（如Set、Get、Hset、Hget等）
	err: 错误信息，如果操作成功则为nil

使用场景：
- BEOpenAppDataNode: 适用于节点级别的配置数据、系统状态数据等
- BEOpenAppDataApp: 适用于应用级别的用户数据、业务数据等

BELoginAsAuthor()(sid string, err error)

简介：

	以作者身份登陆

返回值:

	Sessionid，可用于所有需要会话id的api，执行的时候代表作者身份
*/

type IBEAppData interface {
	// BEOpenAppDataNode 以节点身份打开应用数据区域
	// 节点身份具有可读可写权限，适用于节点级别的数据管理
	// 如果数据区域不存在会自动创建
	// ver: 数据版本（"cur"当前版本，"last"备份版本）
	// mark: 数据区域标记
	// 返回: 弥媒数据ID，可用于后续数据库操作
	BEOpenAppDataNode(ver, mark string) (mmsid string, err error)

	// BEOpenAppDataApp 以应用作者身份打开应用数据区域
	// 应用作者身份具有可读可写权限，适用于应用级别的数据管理
	// 如果数据区域不存在会自动创建
	// ver: 数据版本（"cur"当前版本，"last"备份版本）
	// mark: 数据区域标记
	// 返回: 弥媒数据ID，可用于后续数据库操作
	BEOpenAppDataApp(ver, mark string) (mmsid string, err error)

	// BEMMSync 同步弥媒数据到DHT网络
	// strdhts: DHT网络标识
	// mid: 弥媒ID
	// param: 同步参数
	BEMMSync(strdhts string, mid string, param map[string]string) error

	// BELoginAsApp 以应用身份登录
	// 返回会话ID，可用于需要会话的API操作
	BELoginAsApp() (sid string, err error)

	// BESignPPT 签名PPT文档
	// info: 签名信息
	// period: 有效期（小时）
	// 返回: 签名的PPT字符串
	BESignPPT(info map[string]string, period int) (string, error)

	// BESign 签名操作
	// info: 签名信息
	// 返回: 签名结果字符串
	BESign(info map[string]string) (string, error)

	// BELoginAsAuthor 以作者身份登录
	// 返回会话ID，可用于需要会话的API操作
	BELoginAsAuthor() (sid string, err error)
}

// ILogStub 接口定义了日志操作的所有方法
type ILog interface {
	// Trace 记录跟踪级别日志
	Trace(format string, v ...interface{})

	// Debug 记录调试级别日志
	Debug(format string, v ...interface{})

	// Info 记录信息级别日志
	Info(format string, v ...interface{})

	// Warn 记录警告级别日志
	Warn(format string, v ...interface{})

	// Error 记录错误级别日志
	Error(format string, v ...interface{})

	// Critical 记录严重级别日志
	Critical(format string, v ...interface{})
}
