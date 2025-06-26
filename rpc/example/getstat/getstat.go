// getstat.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/rpc"
)

func main() {
	stub := rpc.InitLApiStubByUrl("127.0.0.1:4800")

	var statCPU rpc.StatCPU
	err := stub.GetVarObj(&statCPU, "", rpc.ApiVarCPU)
	if err != nil {
		panic(err)
	}

	fmt.Println("statCPU", statCPU)

	var statUsage rpc.StatUsage
	err = stub.GetVarObj(&statUsage, "", rpc.ApiVarDisk)
	if err != nil {
		panic(err)
	}

	fmt.Println("StatUsage", statUsage)

	var statVirtualMemory rpc.StatVirtualMemory
	err = stub.GetVarObj(&statVirtualMemory, "", rpc.ApiVarVirtualMemory)
	if err != nil {
		panic(err)
	}

	fmt.Println("StatVirtualMemory", statVirtualMemory)
}
