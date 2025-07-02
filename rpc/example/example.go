// example.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/rpc"
)

// 这个示例演示rpc句柄的生成和使用
func main() {
	//以下是不需要认证身份的示例
	//节点的地址端口是127.0.0.1:4800
	stub := rpc.InitLApiStubByUrl("127.0.0.1:4800")

	var ver string
	err := stub.GetVarObj(&ver, "", rpc.ApiVarVersion)
	if err != nil {
		panic(err)
	}

	fmt.Println("cur ver:", ver)

}
