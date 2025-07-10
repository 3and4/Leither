// mdb.go
package main

import (
	"fmt"

	"github.com/3and4/Leither/lapi"
)

// 这个示例演示了数据库事务相关的功能
// Leither提供了数据库事务功能
// 使用前打开事务， 批量操作之后，如果回滚，则数据恢复到改动前。
// 如果递交则变化一次性生效。

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

	err = testTx(stub, sid)
	if err != nil {
		panic(err)
	}
}

func testTx(lapi *lapi.LApiStub, sid string) error {
	fmt.Println("testx")
	mid, err := lapi.MMCreate(sid, "", "ext", "mark testmdbtx", 2, 0x07276705)
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

	err = write(lapi, mmsid, 0)
	if err != nil {
		return err
	}

	err = read(lapi, mmsid)
	if err != nil {
		return err
	}

	//加事务回滚
	err = testRollback(lapi, mmsid)
	if err != nil {
		return err
	}

	//加事务递交
	err = testCommit(lapi, mmsid)
	if err != nil {
		return err
	}

	return nil
}

func write(lapi *lapi.LApiStub, mmsid string, off int) (err error) {
	fmt.Println("write")
	key := "key"
	for i := 0; i < 10; i++ {
		field := fmt.Sprintf("field%d", i)
		value := fmt.Sprintf("value%d", i+off)
		_, err = lapi.Hset(mmsid, key, field, value)
		if err != nil {
			return err
		}
		fmt.Println("hset", key, field, value)
	}

	return nil
}

func read(lapi *lapi.LApiStub, mmsid string) (err error) {
	fmt.Println("read")
	key := "key"

	for i := 0; i < 10; i++ {
		field := fmt.Sprintf("field%d", i)
		value, err := lapi.Hget(mmsid, key, field)
		if err != nil {
			return err
		}
		fmt.Println("HGet", key, field, value)
	}

	return nil
}

func testRollback(lapi *lapi.LApiStub, mmsid string) (err error) {
	fmt.Println("testRollback")
	// Begin            func(mmsid string, timeout int) error
	// Commit           func(mmsid string) error
	// Rollback         func(mmsid string) error
	err = lapi.Begin(mmsid, 10)
	if err != nil {
		return err
	}

	err = write(lapi, mmsid, 100)
	if err != nil {
		return err
	}

	//这里显示还是100到109
	err = read(lapi, mmsid)
	if err != nil {
		return err
	}

	err = lapi.Rollback(mmsid)
	if err != nil {
		return err
	}

	//回滚之后，显0到9
	err = read(lapi, mmsid)

	return
}

func testCommit(lapi *lapi.LApiStub, mmsid string) (err error) {
	fmt.Println("testCommit")
	err = lapi.Begin(mmsid, 10)
	if err != nil {
		return err
	}

	err = write(lapi, mmsid, 100)
	if err != nil {
		return err
	}

	//这里显示还是100到109
	err = read(lapi, mmsid)
	if err != nil {
		return err
	}

	err = lapi.Commit(mmsid)
	if err != nil {
		return err
	}

	//回滚之后，显0到9
	err = read(lapi, mmsid)

	return
}
