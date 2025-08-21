// getstat.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/lapi"
)

func main() {
	stub := lapi.InitLApiStubByUrl("127.0.0.1:4800")

	var statCPU lapi.StatCPU
	err := stub.GetVarObj(&statCPU, "", lapi.ApiVarCPU)
	if err != nil {
		panic(err)
	}

	fmt.Println("statCPU", statCPU)

	var statUsage lapi.StatUsage
	err = stub.GetVarObj(&statUsage, "", lapi.ApiVarDisk)
	if err != nil {
		panic(err)
	}

	fmt.Println("StatUsage", statUsage)

	var statVirtualMemory lapi.StatVirtualMemory
	err = stub.GetVarObj(&statVirtualMemory, "", lapi.ApiVarVirtualMemory)
	if err != nil {
		panic(err)
	}

	fmt.Println("StatVirtualMemory", statVirtualMemory)
}
