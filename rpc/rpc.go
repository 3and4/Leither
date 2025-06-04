package rpc

import (
	"strings"
	"time"

	"github.com/hprose/hprose-golang/v3/rpc/core"
	. "github.com/hprose/hprose-golang/v3/rpc/websocket"
)

func init() {
	RegisterHandler()
	RegisterTransport()
}

// Package rpc provides RPC functionality for the Leither project
type LApiStub struct {
	VarActStub
}

type VarActStub struct {
	GetVar          func(sid, name string, args ...string) (any, error)
	GetVarByContext func(sid, name string, mapOpt map[string]string) (any, error)
	Act             func(sid, name string, args ...string) error
	GetGobVar       func(sid, name string, args ...string) ([]byte, error)
}

func InitLApiStubByUrl(url string) *LApiStub {
	var stub LApiStub

	hasWS := strings.HasPrefix(url, "ws")
	if !strings.HasPrefix(url, "http") && !hasWS {
		url = "http://" + url
	}

	//这里可以确定是http开的头
	t := strings.Replace(url, "http", "ws", 1) //现在支持
	if t[len(t)-1] != '/' {
		t = t + "/"
	}
	url = t + "ws3/" //待优化

	Client := core.NewClient(url)

	Client.UseService(&stub)
	Client.Timeout = time.Hour * 24
	//	Client.SetTimeout(time.Hour * 24)

	return &stub
}

//测试标签，应该是v0.1.9 上一级加go.mod 清除rpc中的go.mod 
