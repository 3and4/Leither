// backend.go
package lapi

import (
	"fmt"
	"io"
)

var _ IRPC = (*LApiStub)(nil)

var _ LApi = (*BackEndStub)(nil)

var _ IBackEnd = (*BackEndStub)(nil)

// errUnwired 未接线错误（D1）：后端方法函数字段为 nil（宿主未接线）时返回，
// 把"未模拟/未接线"显式报错，不允许静默零值成功——排障时不再伪装成成功。
func errUnwired(method string) error {
	return fmt.Errorf("%s: 未接线", method)
}

type BackEndStub struct {
	*LApiStub
	*SessionStub
	*BEAppDataStub
	*LogStub
	*BEWebStub
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
	Request_MMRoot  = "mmroot" //应用对外根目录，由节点注入

	ApiVarNodeAppCode = "nodeappcode"
)

func GetRequest() map[string]string {
	return nil
}

func GetWriter() io.Writer {
	return nil
}

type SessionStub struct {
	CreateSession  func() (sid string)
	SessionSet     func(sid, key string, value any) error
	SessionGet     func(sid, key string) (value any, err error)
	SessionDelete  func(sid, key string) error
	ReleaseSession func(sid string) error
}

var _ ISession = (*BackEndStub)(nil)

// CreateSession 创建一个session并返回id
// 注意：签名无 error 返回值，未接线时返回空 sid（调用方可判 sid==""）；
// 其余可返回 error 的方法一律显式报"未接线"（D1）。
func (s *BackEndStub) CreateSession() (sid string) {
	if s.SessionStub != nil && s.SessionStub.CreateSession != nil {
		return s.SessionStub.CreateSession()
	}
	return ""
}

// SessionSet 设置session中指定key的值
func (s *BackEndStub) SessionSet(sid, key string, value any) error {
	if s.SessionStub != nil && s.SessionStub.SessionSet != nil {
		return s.SessionStub.SessionSet(sid, key, value)
	}
	return errUnwired("SessionSet")
}

// SessionGet 获取session中指定key的值
func (s *BackEndStub) SessionGet(sid, key string) (value any, err error) {
	if s.SessionStub != nil && s.SessionStub.SessionGet != nil {
		return s.SessionStub.SessionGet(sid, key)
	}
	return nil, errUnwired("SessionGet")
}

// SessionDelete 删除session中指定key的值
func (s *BackEndStub) SessionDelete(sid, key string) error {
	if s.SessionStub != nil && s.SessionStub.SessionDelete != nil {
		return s.SessionStub.SessionDelete(sid, key)
	}
	return errUnwired("SessionDelete")
}

// ReleaseSession 释放指定的session
func (s *BackEndStub) ReleaseSession(sid string) error {
	if s.SessionStub != nil && s.SessionStub.ReleaseSession != nil {
		return s.SessionStub.ReleaseSession(sid)
	}
	return errUnwired("ReleaseSession")
}

var _ ILog = (*BackEndStub)(nil)

// BEOpenAppDataNode 打开节点的弥媒数据
func (s *BackEndStub) BEOpenAppDataNode(ver, mark string) (mmsid string, err error) {
	if s.BEAppDataStub != nil && s.BEAppDataStub.BEOpenAppDataNode != nil {
		return s.BEAppDataStub.BEOpenAppDataNode(ver, mark)
	}
	return "", errUnwired("BEOpenAppDataNode")
}

// BEOpenAppDataApp 打开应用的弥媒数据
func (s *BackEndStub) BEOpenAppDataApp(ver, mark string) (mmsid string, err error) {
	if s.BEAppDataStub != nil && s.BEAppDataStub.BEOpenAppDataApp != nil {
		return s.BEAppDataStub.BEOpenAppDataApp(ver, mark)
	}
	return "", errUnwired("BEOpenAppDataApp")
}

// BEMMSync 同步弥媒数据
func (s *BackEndStub) BEMMSync(strdhts string, mid string, param map[string]string) error {
	if s.BEAppDataStub != nil && s.BEAppDataStub.BEMMSync != nil {
		return s.BEAppDataStub.BEMMSync(strdhts, mid, param)
	}
	return errUnwired("BEMMSync")
}

// BELoginAsAuthor 以作者身份登录
func (s *BackEndStub) BELoginAsAuthor() (sid string, err error) {
	if s.BEAppDataStub != nil && s.BEAppDataStub.BELoginAsAuthor != nil {
		return s.BEAppDataStub.BELoginAsAuthor()
	}
	return "", errUnwired("BELoginAsAuthor")
}

// BELoginAsApp 以应用身份登录
func (s *BackEndStub) BELoginAsApp() (sid string, err error) {
	if s.BEAppDataStub != nil && s.BEAppDataStub.BELoginAsApp != nil {
		return s.BEAppDataStub.BELoginAsApp()
	}
	return "", errUnwired("BELoginAsApp")
}

// BESignPPT 签名PPT文档
func (s *BackEndStub) BESignPPT(info map[string]string, period int) (string, error) {
	if s.BEAppDataStub != nil && s.BEAppDataStub.BESignPPT != nil {
		return s.BEAppDataStub.BESignPPT(info, period)
	}
	return "", errUnwired("BESignPPT")
}

// BESign 签名操作
func (s *BackEndStub) BESign(info map[string]string) (string, error) {
	if s.BEAppDataStub != nil && s.BEAppDataStub.BESign != nil {
		return s.BEAppDataStub.BESign(info)
	}
	return "", errUnwired("BESign")
}

/*
BEAppDataStub

这里面放的是应用中后端数据操作的api接口。

一、BEOpenAppDataNode(ver, mark string) (mmsid string, err error)
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
	BELoginAsApp      func() (sid string, err error)
	BESignPPT         func(info map[string]string, period int) (string, error)
	BESign            func(info map[string]string) (string, error)
	BELoginAsAuthor   func() (sid string, err error)
}

/*
LogStub

这里面放的是应用中日志操作的api接口。
*/
type LogStub struct {
	Trace    func(format string, v ...interface{})
	Debug    func(format string, v ...interface{})
	Info     func(format string, v ...interface{})
	Warn     func(format string, v ...interface{})
	Error    func(format string, v ...interface{})
	Critical func(format string, v ...interface{})
}

// Trace 记录跟踪级别日志
func (s *BackEndStub) Trace(format string, v ...interface{}) {
	if s.LogStub != nil && s.LogStub.Trace != nil {
		s.LogStub.Trace(format, v...)
	}
}

// Debug 记录调试级别日志
func (s *BackEndStub) Debug(format string, v ...interface{}) {
	if s.LogStub != nil && s.LogStub.Debug != nil {
		s.LogStub.Debug(format, v...)
	}
}

// Info 记录信息级别日志
func (s *BackEndStub) Info(format string, v ...interface{}) {
	if s.LogStub != nil && s.LogStub.Info != nil {
		s.LogStub.Info(format, v...)
	}
}

// Warn 记录警告级别日志
func (s *BackEndStub) Warn(format string, v ...interface{}) {
	if s.LogStub != nil && s.LogStub.Warn != nil {
		s.LogStub.Warn(format, v...)
	}
}

// Error 记录错误级别日志
func (s *BackEndStub) Error(format string, v ...interface{}) {
	if s.LogStub != nil && s.LogStub.Error != nil {
		s.LogStub.Error(format, v...)
	}
}

// Critical 记录严重级别日志
func (s *BackEndStub) Critical(format string, v ...interface{}) {
	if s.LogStub != nil && s.LogStub.Critical != nil {
		s.LogStub.Critical(format, v...)
	}
}

// BEReadFile 读取应用 mmroot 文件并可选预处理（nil 安全：未接线返回错误）。
func (s *BackEndStub) BEReadFile(name string, ops ...string) ([]byte, error) {
	if s.BEWebStub != nil && s.BEWebStub.BEReadFile != nil {
		return s.BEWebStub.BEReadFile(name, ops...)
	}
	return nil, fmt.Errorf("BEReadFile: 未接线")
}
