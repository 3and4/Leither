package lapp

import (
	"fmt"
	"io"

	api "github.com/3and4/Leither/lapi"
)

func RunMApp(Entry string, Request map[string]string, args []any, wr io.Writer) (any, error) {
	lapi := api.GetLApi()
	if lapi == nil {
		return nil, fmt.Errorf("lapi is nil")
	}

	if wr != nil {
		fmt.Fprintln(wr, "write 2 wr")
	}

	ver, err := lapi.GetVar("", "ver")
	if err != nil {
		return nil, err
	}
	fmt.Println("Hello, Welcome to Leither ", ver)
	return ver, nil
}

// func main() {
// 	fmt.Println("Hello, World!")
// }
