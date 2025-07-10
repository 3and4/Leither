// mdb.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/lapi"
)

// 这个示例演示了数据库备份相关的功能
// Leither提供了弥媒备份功能
// 弥媒备份之后，当前弥媒就进行一次数据镜像和一个对应版本。
// 镜像的内容包括弥媒中的文件或数据库，以及弥媒引用的相关资源（文件或弥媒）

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

	err = testBackup(stub, sid)
	if err != nil {
		panic(err)
	}
}

func testBackup(lapi *lapi.LApiStub, sid string) error {
	fmt.Println("testBackup")
	mid, err := lapi.MMCreate(sid, "", "ext", "mark testmdbBackup", 2, 0x07276705)
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

	err = write(lapi, mmsid)
	if err != nil {
		return err
	}

	err = read(lapi, mmsid)
	if err != nil {
		return err
	}

	//备份
	ver, err := lapi.MMBackup(mmsid, mid, "MMBackup")
	if err != nil {
		return err
	}

	fmt.Println("ver =", ver)

	fmt.Println("read ver", ver)
	mmsid2, err := lapi.MMOpen(sid, mid, ver)
	if err != nil {
		return err
	}
	defer lapi.MMClose(mmsid2)
	fmt.Println("mmsid2=", mmsid2)

	err = read(lapi, mmsid2)
	if err != nil {
		return err
	}

	return nil
}

// 这里用list类型
func write(lapi *lapi.LApiStub, mmsid string) (err error) {
	fmt.Println("write")
	key := "key"
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("value%d", i)
		_, err = lapi.Lpush(mmsid, key, value)
		if err != nil {
			return err
		}
		fmt.Println("Lpush", key, value)
	}

	return nil
}

func read(lapi *lapi.LApiStub, mmsid string) (err error) {
	fmt.Println("read")
	key := "key"

	for i := int32(0); i < 10; i++ {
		value, err := lapi.Lindex(mmsid, key, i)
		if err != nil {
			return err
		}
		fmt.Println("Lindex", key, i, value)
	}

	return nil
}
