// mdb.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/lapi"
)

// 这个示例演示了数据库的相关操作
// Leither提供了数据库功能，接口参考redis
// 数据库工作在一个弥媒容器中，使用之前要先打开弥媒，使用之后关闭

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

	err = testmdb(stub, sid)
	if err != nil {
		panic(err)
	}
}

func testmdb(lapi *lapi.LApiStub, sid string) error {
	fmt.Println("testmdb")

	mid, err := lapi.MMCreate(sid, "", "ext", "mark testmdbbase", 2, 0x07276705)
	if err != nil {
		return err
	}
	fmt.Println("mid=", mid)

	mmsid, err := lapi.MMOpen(sid, mid, "cur")
	if err != nil {
		return err
	}
	defer lapi.MMClose(mmsid)
	fmt.Println("mmsid=", mmsid)

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		err := lapi.Set(mmsid, key, value)
		if err != nil {
			return err
		}
		fmt.Println("Set", key, value)
	}

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value, err := lapi.Get(mmsid, key)
		if err != nil {
			return err
		}
		fmt.Println("Get", key, value)
	}

	// //	Zrevrangebyscore func(mmsid, key string, mins, maxs int64, offset int, count int) (ret []ScorePair, err error)
	// ret, err := lapi.Zrevrangebyscore(mmsid, "key", 0, 10, 0, -1)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("Zrevrangebyscore", ret)

	return nil
}
