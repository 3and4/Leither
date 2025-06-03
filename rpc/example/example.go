// example.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/rpc"
)

func main() {

	stub := rpc.InitLApiStubByUrl("127.0.0.1:4800")
	ver, err := stub.GetVar("", "ver")
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello World!", ver)
}
