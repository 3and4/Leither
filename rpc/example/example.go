// example.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/rpc"
)

func main() {
	//节点的地址端口是127.0.0.1:4800
	stub := rpc.InitLApiStubByUrl("127.0.0.1:4800")
	ver, err := stub.GetVar("", "ver")
	if err != nil {
		panic(err)
	}

	fmt.Println("ver:", ver)
}
