// login_client — 最小可运行的 Leither 密钥认证集成示例。
//
// 依据《Leither 密钥认证接入指南》§3：通过 hprose WebSocket RPC 调
// LoginWithPPT(ppt) 登录，随后 GetVar(sid,"userid") 读取身份并打印，
// 最后 Logout(sid, "") 登出。任何失败以非零退出码退出。
//
// 注意：文档 §3 的 client.Invoke(name, args, &reply) 写法在
// hprose-golang v3.0.16 下不能编译（v3 的 Client.Invoke 只有
// (name, args) 两个参数，返回 []interface{}）。本程序改用 v3 的
// 类型化代理 UseService 模式，方法名与文档表格一一对应。
//
// 用法：
//
//	NODE_WS=ws://127.0.0.1:4800/ws/ go run . -ppt /path/to/login.ppt
//
// 环境变量：
//
//	NODE_WS  节点 WebSocket 地址（形如 ws://127.0.0.1:PORT/ws/，必填）
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hprose/hprose-golang/v3/rpc/core"
	"github.com/hprose/hprose-golang/v3/rpc/websocket"
)

// LoginReply 对应文档 §3 的登录返回结构。
type LoginReply struct {
	Uid     string
	Sid     string
	Isystem bool
}

// nodeStub 是节点 IAuth/GetVar RPC 的类型化代理（方法名与文档 §3/§4 一致）。
type nodeStub struct {
	LoginWithPPT func(ppt string) (*LoginReply, error)
	GetVar       func(sid, name string) (string, error)
	Logout       func(sid, info string) error
}

func fail(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
	os.Exit(1)
}

func main() {
	pptPath := flag.String("ppt", "", "登录 PPT 文件路径（必填）")
	flag.Parse()

	nodeWS := os.Getenv("NODE_WS")
	if nodeWS == "" {
		fail("环境变量 NODE_WS 未设置（形如 ws://127.0.0.1:PORT/ws/）")
	}
	if *pptPath == "" {
		fail("缺少 -ppt 参数（登录 PPT 文件路径）")
	}

	data, err := os.ReadFile(*pptPath)
	if err != nil {
		fail("读取 PPT 文件失败: %v", err)
	}
	ppt := strings.TrimSpace(string(data))
	if ppt == "" {
		fail("PPT 文件为空: %s", *pptPath)
	}

	websocket.RegisterTransport()
	client := core.NewClient(nodeWS)
	var stub *nodeStub
	client.UseService(&stub)

	// 1) 登录
	reply, err := stub.LoginWithPPT(ppt)
	if err != nil {
		fail("LoginWithPPT 失败: %v", err)
	}
	if reply == nil || reply.Sid == "" {
		fail("LoginWithPPT 返回为空 sid（reply=%+v）", reply)
	}
	fmt.Printf("LoginWithPPT OK: uid=%s sid=%s isystem=%v\n", reply.Uid, reply.Sid, reply.Isystem)

	// 2) 读取身份
	uid, err := stub.GetVar(reply.Sid, "userid")
	if err != nil {
		fail("GetVar(sid, \"userid\") 失败: %v", err)
	}
	fmt.Printf("GetVar userid = %s\n", uid)

	// 文档 §1.1：userid == 密钥 id；§3：登录返回的 Uid 亦为此值
	if reply.Uid != "" && uid != reply.Uid {
		fail("身份不一致: GetVar 返回 %q, 登录返回 %q", uid, reply.Uid)
	}

	// 3) 登出（文档 §1.3：异步释放，不做立即失效断言）
	if err := stub.Logout(reply.Sid, ""); err != nil {
		fail("Logout 失败: %v", err)
	}
	fmt.Println("Logout OK")
}
