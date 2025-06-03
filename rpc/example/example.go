// example.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/rpc"
	"github.com/hprose/hprose-golang/v3/rpc/core"
	. "github.com/hprose/hprose-golang/v3/rpc/websocket"
)

func init() {
	RegisterHandler()
	RegisterTransport()
}

func main() {
	var stub rpc.LApiStub
	// var stub struct {
	// 	GetVar func(sid, name string, args ...string) (any, error)
	// }

	client := core.NewClient("ws://127.0.0.1:4800/ws3/")
	// client.Use(log.Plugin)
	// var proxy struct {
	// 	Hello func(name string) (string, error)
	// }
	client.UseService(&stub)
	ver, err := stub.GetVar("", "ver")
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello World!", ver)
}
