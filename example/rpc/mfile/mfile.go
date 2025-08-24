// mdb.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/lapi"
)

// 这个示例演示了弥媒文件功能
// Leither提供了弥媒文件功能
// 弥媒工作在一个弥媒容器中，使用之前要先打开弥媒，使用之后关闭
func main() {
	//以下是不需要认证身份的示例
	//节点的地址端口是127.0.0.1:4800
	stub := lapi.InitLApiStubByUrl("127.0.0.1:4800")

	//获取一个通行证
	//这是通过127.0.0.1:4800,向本地节点申请的
	//有效期是24小时
	strPPT, err := lapi.GetLocalPassport(4800, 24)
	if err != nil {
		panic(err)
	}

	reply, err := stub.LoginWithPPT(strPPT)
	if err != nil {
		panic(err)
	}
	sid := reply.Sid
	defer stub.Logout(sid, "")

	err = testmfile(stub, sid)
	if err != nil {
		panic(err)
	}
}

func testmfile(stub *lapi.LApiStub, sid string) error {
	fmt.Println("testmfile")

	//	LeitherApp       = "Fc1BRTFafOGzq5P8KmkVJqwS2v2" //memo:leitherapp
	mmsid, err := stub.MMOpen(sid, lapi.LeitherApp, "last")
	if err != nil {
		return err
	}
	defer stub.MMClose(mmsid)
	fmt.Println("mmsid=", mmsid)

	ayfi, err := stub.MFReaddir(mmsid, -1)
	if err != nil {
		return err
	}
	for _, fi := range ayfi {
		fmt.Println("fi.Name", fi.Name())
	}

	// for i := 0; i < 10; i++ {
	// 	key := fmt.Sprintf("key%d", i)
	// 	value := fmt.Sprintf("value%d", i)
	// 	err := lapi.Set(mmsid, key, value)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println("Set", key, value)
	// }

	// for i := 0; i < 10; i++ {
	// 	key := fmt.Sprintf("key%d", i)
	// 	value, err := lapi.Get(mmsid, key)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	fmt.Println("Get", key, value)
	// }

	// // //	Zrevrangebyscore func(mmsid, key string, mins, maxs int64, offset int, count int) (ret []ScorePair, err error)
	// // ret, err := lapi.Zrevrangebyscore(mmsid, "key", 0, 10, 0, -1)
	// // if err != nil {
	// // 	return nil, err
	// // }
	// // fmt.Println("Zrevrangebyscore", ret)

	return nil
}
